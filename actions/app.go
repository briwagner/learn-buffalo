package actions

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"learnbuffalo/locales"
	"learnbuffalo/models"
	"learnbuffalo/public"
	"learnbuffalo/sendgrid_mailer"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/events"
	"github.com/gobuffalo/middleware/csrf"
	"github.com/gobuffalo/middleware/forcessl"
	"github.com/gobuffalo/middleware/i18n"
	"github.com/gobuffalo/middleware/paramlogger"
	"github.com/unrolled/secure"
)

var ENV = envy.Get("GO_ENV", "development")

var (
	app     *buffalo.App
	appOnce sync.Once
	T       *i18n.Translator
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	appOnce.Do(func() {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_learnbuffalo_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		// c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Check MongoDB connection
		// err := mongoconnector.Ping()
		// if err != nil {
		// 	log.Print(err)
		// }

		// Check Sendgrid_Mailer
		err := sendgrid_mailer.Register()
		if err != nil {
			log.Print(err)
		}

		// Setup and use translations:
		app.Use(translations())

		app.GET("/", HomeHandler)

		//AuthMiddlewares
		// app.Use(SetCurrentUser)
		// app.Use(Authorize)

		app.GET("/tags/{id}", TagsShow)
		app.GET("/blogs/new", Authorize(SetCurrentUser(BlogsCreate)))
		app.GET("/blogs/{id}", BlogsShow)
		app.GET("/blogs/", BlogsIndex)
		app.POST("/blogs", Authorize(SetCurrentUser(BlogsNew)))

		// Routes for User registration
		app.GET("/users/new", UsersNew)
		app.POST("/users", UsersCreate)

		// Authentication routes.
		app.GET("/auth", SetCurrentUser(AuthLanding))
		app.GET("/auth/login", AuthNew)
		app.POST("/auth", AuthCreate)
		app.DELETE("/auth", Authorize(SetCurrentUser(AuthDestroy)))

		// Alternative method to add middleware.
		// auth := app.Group("/auth")
		// auth.GET("/", SetCurrentUser(AuthLanding))
		// auth.GET("/login", AuthNew)
		// auth.POST("/", AuthCreate)
		// auth.DELETE("/", Authorize(SetCurrentUser(AuthDestroy)))

		app.ServeFiles("/", http.FS(public.FS())) // serve files from the public directory
	})

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(locales.FS(), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

func init() {
	_, err := events.Listen(func(e events.Event) {
		if e.Kind == "learnbuffalo:user:register" {
			slackURL := os.Getenv("SLACK_URL")
			if slackURL == "" {
				log.Print("slack url not set")
				return
			}

			username, err := e.Payload.Pluck("username")
			if err != nil {
				log.Print(err.Error())
				return
			}

			// Format and send message to Slack.
			msg := strings.NewReader(fmt.Sprintf(`{"text": "new user added %s"}`, username))
			req, err := http.NewRequest("POST", slackURL, msg)
			if err != nil {
				log.Printf("Could not build post request: %v", err.Error())
				return
			}
			req.Header.Set("Content-Type", "application/json")

			// Override default timeout.
			client := http.Client{
				Timeout: 5 * time.Second,
			}
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Could not make post request: %v", err.Error())
				return
			}

			// Failed to post to slack.
			if resp.StatusCode != 200 {
				defer resp.Body.Close()
				errMsg, _ := io.ReadAll(resp.Body)
				log.Printf("Slackbot error %s", errMsg)
			}
		}
	})

	if err != nil {
		log.Print(err.Error())
	}
}

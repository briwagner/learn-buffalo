package actions

import (
	"learnbuffalo/public"
	"learnbuffalo/templates"
	"time"

	"github.com/gobuffalo/buffalo/render"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// fs.FS containing templates
		TemplatesFS: templates.FS(),

		// fs.FS containing assets
		AssetsFS: public.FS(),

		// Add template helpers here:
		Helpers: render.Helpers{
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			// forms.FormKey:     forms.Form,
			// forms.FormForKey:  forms.FormFor,
			"formatTime": func(t time.Time, format string) string {
				switch format {
				case "short":
					return t.Format("Jan. 2, 2006")
				case "long":
					return t.Format("January 2, 2006")
				}
				return ""
			},
		},
	})
}

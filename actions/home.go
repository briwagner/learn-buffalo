package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

func AboutHandler(c buffalo.Context) error {
	c.Set("title", "About page")
	c.Set("content", "This is the about page.")
	return c.Render(http.StatusOK, r.HTML("about.html"))
}

type Payload struct {
	Data User `json:"data"`
}

type User struct {
	ID        int    `json":id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func UserHandler(c buffalo.Context) error {
	user_id := c.Param("user_id")
	user, err := getUserInfo(user_id)
	if err != nil {
		c.Flash().Add("warning", "Bad connection")
		c.Redirect(301, "/")
	}
	c.Set("user", user)
	return c.Render(http.StatusOK, r.HTML("user.html"))
}

// Make api request for user data.
func getUserInfo(user_id string) (User, error) {
	url := "https://reqres.in/api/users/" + user_id

	resp, err := http.Get(url)
	if err != nil {
		return User{}, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var data Payload
	err = json.Unmarshal(b, &data)
	if err != nil {
		return User{}, &json.MarshalerError{}
	}
	return data.Data, nil
}

package actions

import (
	"fmt"
	"learnbuffalo/models"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

// BlogsShow shows blog by ID.
func BlogsShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	blog := models.Blog{}
	blog_id := c.Param("id")

	err := tx.Eager().Find(&blog, blog_id)
	if err != nil {
		c.Flash().Add("warning", "Blog not found.")
		c.Redirect(404, "/")
	}

	if len(blog.BlogTags) > 0 {
		tag := blog.BlogTags[0]

		bt := models.BlogsTags{}
		err = tx.Eager("Blog").Where("tag_id = ?", tag.ID).Where("blog_id <> ?", blog.ID).Limit(3).All(&bt)
		if err != nil {
			log.Println(err)
		}

		c.Set("relatedBlogs", bt)
	}

	c.Set("blog", blog)
	return c.Render(http.StatusOK, r.HTML("blogs/show.html"))
}

// BlogsCreate shows the form to create a new blog.
func BlogsCreate(c buffalo.Context) error {
	b := models.Blog{}
	c.Set("blog", b)
	return c.Render(http.StatusOK, r.HTML("blogs/create.html"))
}

// BlogsNew responds to POST request to create a new blog.
func BlogsNew(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	b := &models.Blog{}
	err := c.Bind(b)
	if err != nil {
		c.Flash().Add("warning", "Form binding error")
		return c.Redirect(301, "/")
	}

	u := &models.User{}
	err = tx.Last(u)
	if err != nil {
		c.Flash().Add("warning", "Cannot find user")
		return c.Redirect(301, "/")
	}

	b.User = u

	verrs, err := tx.ValidateAndCreate(b)
	if err != nil {
		return c.Redirect(301, "/")
	}

	if verrs.HasAny() {
		c.Flash().Add("warning", "Form validation errors")
		return c.Redirect(301, "/")
	}

	c.Set("blog", b)
	c.Flash().Add("info", "Created blog")
	return c.Redirect(301, fmt.Sprintf("/blogs/%s", b.ID.String()))
}

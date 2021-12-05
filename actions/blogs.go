package actions

import (
	"database/sql"
	"fmt"
	"learnbuffalo/models"
	"log"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// BlogsShow shows blog by ID.
func BlogsIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	blogs := models.Blogs{}

	err := tx.All(&blogs)
	if err != nil {
		c.Flash().Add("warning", "No blogs found")
		c.Redirect(307, "/")
	}

	c.Set("blogs", blogs)
	return c.Render(http.StatusOK, r.HTML("blogs/index"))

}

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
	return c.Render(http.StatusOK, r.HTML("blogs/show"))
}

// BlogsCreate shows the form to create a new blog.
func BlogsCreate(c buffalo.Context) error {
	b := models.Blog{}
	c.Set("blog", b)
	return c.Render(http.StatusOK, r.HTML("blogs/create.html", "admin.html"))
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

	u := c.Value("current_user").(*models.User)

	if u.ID == uuid.Nil {
		c.Flash().Add("warning", "Cannot find user")
		return c.Redirect(301, "/")
	}

	b.User = u

	// Load tags to create blog-tags.
	err = c.Request().ParseForm()
	if err != nil {
		log.Println(err)
		c.Flash().Add("error", "Error parsing form.")
		return c.Redirect(301, "/")
	}

	// Process Tags field as comma-separated text.
	ts := c.Request().FormValue("Tags")
	newTags := strings.Split(ts, ",")
	for _, v := range newTags {
		v = strings.TrimSpace(v)

		t := &models.Tag{}
		err := tx.Where("name = ?", v).Last(t)
		if err != nil {
			// Tag not found in DB, so create one.
			if errors.Cause(err) == sql.ErrNoRows {
				t.Name = v
				err2 := tx.Create(t)
				if err2 != nil {
					log.Fatal(err2)
				}
			} else {
				log.Fatal(err)
				continue
			}
		}

		b.BlogTags = append(b.BlogTags, *t)
	}

	verrs, err := tx.Eager().ValidateAndCreate(b)
	if err != nil {
		return c.Redirect(301, "/")
	}

	if verrs.HasAny() {
		c.Flash().Add("warning", "Form validation errors")
		c.Set("blog", b)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("blogs/create.html", "admin.html"))
	}

	c.Flash().Add("info", "Created blog")
	return c.Redirect(301, fmt.Sprintf("/blogs/%s", b.ID.String()))
}

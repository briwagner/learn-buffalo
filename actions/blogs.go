package actions

import (
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

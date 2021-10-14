package actions

import (
	"learnbuffalo/models"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

// TagsShow returns tag by ID.
func TagsShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	tag := models.Tag{}
	tag_id := c.Param("id")

	err := tx.Eager().Find(&tag, tag_id)
	if err != nil {
		c.Flash().Add("warning", "Error loading tag.")
		c.Redirect(301, "/")
	}

	c.Set("tag", tag)
	return c.Render(http.StatusOK, r.HTML("tags/show.html"))
}

package grifts

import (
	"bufio"
	"encoding/json"
	"learnbuffalo/models"
	"os"

	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		savedTags := map[string]models.Tag{}

		// Add tags
		tags := []string{"Food", "Computers", "Travel"}
		for _, v := range tags {
			t := models.Tag{Name: v}
			err := models.DB.Create(&t)
			if err != nil {
				panic(err)
			}
			savedTags[t.Name] = t
		}

		// Add a user
		nikola := models.User{FirstName: "Nikola", LastName: "Tesla", Age: 35}
		err := models.DB.Create(&nikola)
		if err != nil {
			panic(err)
		}

		// Import blogs from file
		f, err := os.Open("blogs.json")
		if err != nil {
			panic(err)
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			var b models.Blog
			err = json.Unmarshal(s.Bytes(), &b)
			if err != nil {
				panic(err)
			}
			// Add author
			b.User = &nikola

			// Add tags
			t, ok := savedTags["Computers"]
			if ok {
				b.BlogTags = append(b.BlogTags, t)
			}
			err := models.DB.Create(&b)
			if err != nil {
				panic(err)
			}
		}

		return nil
	})

})

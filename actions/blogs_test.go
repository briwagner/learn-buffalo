package actions

import (
	"fmt"
	"learnbuffalo/models"
)

func (as *ActionSuite) Test_Blogs_Show() {
	as.LoadFixture("sample blogs")

	b := models.Blog{}
	err := as.DB.Last(&b)
	if err != nil {
		panic(err)
	}

	res := as.HTML(fmt.Sprintf("/blogs/%s", b.ID)).Get()
	body := res.Body.String()
	as.Contains(body, "Second Blog post", "Blog title appears on Show page.")
	as.Contains(body, "Nikola Tesla", "Author name appears on blog Show page.")
	as.Contains(body, "Computers", "Tag name appears on blog Show page.")
}

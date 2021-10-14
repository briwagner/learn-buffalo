package actions

import (
	"fmt"
	"learnbuffalo/models"
)

func (as *ActionSuite) Test_Tags_Show() {
	as.LoadFixture("sample tags")

	t := models.Tag{}
	err := as.DB.Last(&t)
	if err != nil {
		panic(err)
	}

	res := as.HTML(fmt.Sprintf("/tags/%s", t.ID)).Get()
	body := res.Body.String()
	as.Contains(body, "Travel", "Tag name appears on Show page.")
}

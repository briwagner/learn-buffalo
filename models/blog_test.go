package models

import "github.com/gofrs/uuid"

func (ms *ModelSuite) Test_Blog() {
	b1 := &Blog{
		Title: "First blog",
		Body:  "<p>Interesting content here</p>",
	}

	b2 := &Blog{
		Title: "Second blog",
		Body:  "<p>More interesting content here</p>",
	}

	u := &User{
		FirstName: "Joe",
		LastName:  "Smith",
		Age:       25,
		Blogs:     Blogs{*b1, *b2},
	}

	db := ms.DB
	verrs, err := db.Eager().ValidateAndCreate(u)
	if err != nil {
		panic(err)
	}

	ms.NotEqual(uuid.Nil, u.Blogs[0].ID, "Blog ID is generated when saved to DB.")
	ms.NotEqual(uuid.Nil, u.Blogs[1].ID, "Blog ID is generated when saved to DB.")
	ms.False(verrs.HasAny(), "Blog and user creation have no validation errors.")
}

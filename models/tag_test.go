package models

import (
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) Test_Tag() {

	ms.LoadFixture("sample users")
	t := &Tag{
		Name: "Miscellaneous",
	}

	db := ms.DB
	verrs, err := db.ValidateAndCreate(t)
	if err != nil {
		panic(err)
	}

	ms.NotEqual(uuid.Nil, t.ID, "Tag ID is generated when saved to DB.")
	ms.False(verrs.HasAny(), "Tag has no validation errors.")
}

func (ms *ModelSuite) Test_TagGetBlogs() {
	t1 := &Tag{
		Name: "Miscellaneous",
	}

	t2 := &Tag{
		Name: "Food",
	}

	db := ms.DB
	verrs, err := db.ValidateAndCreate(t1)
	ms.NoError(err)

	ms.NotEqual(uuid.Nil, t1.ID, "Tag ID is generated when saved to DB.")
	ms.False(verrs.HasAny(), "Tag has no validation errors.")

	verrs, err = db.ValidateAndCreate(t2)
	ms.NoError(err)

	ms.NotEqual(uuid.Nil, t2.ID, "Tag ID is generated when saved to DB.")
	ms.False(verrs.HasAny(), "Tag has no validation errors.")

	b1 := &Blog{
		Title:    "First blog",
		Body:     "<p>Interesting content here</p>",
		BlogTags: Tags{*t1, *t2},
	}

	b2 := &Blog{
		Title:    "Second blog",
		Body:     "<p>More interesting content here</p>",
		BlogTags: Tags{*t2},
	}

	u := User{
		FirstName:            "Joe",
		LastName:             "Smith",
		Age:                  25,
		Email:                "joe@smith.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Blogs:                Blogs{*b1, *b2},
	}

	err = db.Eager().Create(&u)
	ms.NoError(err)

	// Load tag.
	tag := &Tag{}
	err = db.Find(tag, t2.ID)
	ms.NoError(err)
	ms.Empty(tag.RelatedBlogs, "Blogs not loaded with tag.")

	err = tag.GetBlogs(db)
	ms.NoError(err)
	ms.NotEmpty(tag.RelatedBlogs, "GetBlogs function loads related blogs for tag.")
}

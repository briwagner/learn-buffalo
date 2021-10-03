package models

import "github.com/gofrs/uuid"

func (ms *ModelSuite) Test_Tag() {
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
	if err != nil {
		panic(err)
	}
	ms.NotEqual(uuid.Nil, t1.ID, "Tag ID is generated when saved to DB.")
	ms.False(verrs.HasAny(), "Tag has no validation errors.")

	verrs, err = db.ValidateAndCreate(t2)
	if err != nil {
		panic(err)
	}
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
		FirstName: "Joe",
		LastName:  "Smith",
		Age:       25,
		Blogs:     Blogs{*b1, *b2},
	}

	_, err = db.Eager().ValidateAndCreate(&u)
	if err != nil {
		panic(err)
	}

	// Load tag.
	tag := &Tag{}
	err = db.Find(tag, t2.ID)
	if err != nil {
		panic(err)
	}

	ms.Empty(tag.RelatedBlogs, "Blogs not loaded with tag.")
	err = tag.GetBlogs(db)
	if err != nil {
		panic(err)
	}
	ms.NotEmpty(tag.RelatedBlogs, "GetBlogs function loads related blogs for tag.")
}

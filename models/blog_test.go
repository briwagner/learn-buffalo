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
		FirstName:            "Joe",
		LastName:             "Smith",
		Age:                  25,
		Email:                "joe@smith.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}

	db := ms.DB
	_, err := u.Create(db)
	ms.NoError(err)

	b1.User = u
	b2.User = u
	verrs1, err := db.ValidateAndCreate(b1)
	ms.NoError(err)
	verrs2, err := db.ValidateAndCreate(b2)
	ms.NoError(err)

	ms.NotEqual(uuid.Nil, b1.ID, "Blog ID is generated when saved to DB.")
	ms.NotEqual(uuid.Nil, b2.ID, "Blog ID is generated when saved to DB.")
	ms.False(verrs1.HasAny(), "Blog1 creation has no validation errors.")
	ms.False(verrs2.HasAny(), "Blog2 creation has no validation errors.")
}

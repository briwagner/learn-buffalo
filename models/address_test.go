package models

import (
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) Test_Address() {
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
	if err != nil {
		panic(err)
	}

	a := &Address{
		Street: "1 Main Street",
		City:   "Everytown",
		State:  "IL",
		Zip:    "12345",
		UserID: u.ID,
	}
	verrs, err := db.ValidateAndCreate(a)
	if err != nil {
		panic(err)
	}

	ms.False(verrs.HasAny(), "Address creation has no validation errors.")
	ms.NotEqual(uuid.Nil, a.ID, "Address ID is generated when saved to DB.")

	a2 := &Address{}
	db.Eager().Find(a2, a.ID)

	ms.Equal(a.ID, a2.ID, "Find Address in database using ID.")
	ms.Equal("Joe Smith", a2.User.FullName(), "User is loaded with Address.")
}

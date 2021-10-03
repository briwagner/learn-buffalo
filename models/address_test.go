package models

import (
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) Test_Address() {
	a := &Address{
		Street: "1 Main Street",
		City:   "Everytown",
		State:  "IL",
		Zip:    "12345",
	}

	u := &User{
		FirstName:   "Joe",
		LastName:    "Smith",
		Age:         25,
		UserAddress: *a,
	}

	db := ms.DB
	verrs, err := db.Eager().ValidateAndCreate(u)
	if err != nil {
		panic(err)
	}

	ms.NotEqual(uuid.Nil, u.UserAddress.ID, "Address ID is generated when saved to DB.")
	ms.False(verrs.HasAny(), "Address and user creation have no validation errors.")

	a2 := &Address{}
	db.Eager().Find(a2, u.UserAddress.ID)

	ms.Equal(u.UserAddress.ID, a2.ID, "Find Address in database using ID.")
	ms.Equal("Joe Smith", a2.User.FullName(), "User is loaded with Address.")
}

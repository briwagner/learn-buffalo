package models

import "github.com/gofrs/uuid"

func (ms *ModelSuite) Test_User() {
	u := &User{
		FirstName: "Nikola",
		LastName:  "Tesla",
	}

	ms.Equal("Nikola Tesla", u.FullName(), "FullName returns user name.")

	db := ms.DB
	verrs, err := db.ValidateAndCreate(u)
	if err != nil {
		panic(err)
	}

	ms.NotNil(u.ID, "User ID is generated when saved to DB.")
	ms.True(verrs.HasAny(), "User cannot be created without age field.")
}

func (ms *ModelSuite) Test_UserAddress() {
	u := &User{
		FirstName: "Nikola",
		LastName:  "Tesla",
		Age:       25,
		UserAddress: Address{
			Street: "1 Main Street",
			City:   "Everytown",
			State:  "IL",
			Zip:    "12345",
		},
	}

	db := ms.DB
	_, err := db.Eager().ValidateAndCreate(u)
	if err != nil {
		panic(err)
	}

	ms.NotEqual(uuid.Nil, u.UserAddress.ID, "Address saved along with User.")
}

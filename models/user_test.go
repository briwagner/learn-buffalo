package models

func (ms *ModelSuite) Test_User() {
	u := &User{
		FirstName: "Nikola",
		LastName:  "Tesla",
		Age:       86,
	}

	ms.Equal("Nikola Tesla", u.FullName(), "FullName returns user name.")

	db := ms.DB
	verrs, err := db.ValidateAndCreate(u)
	ms.NoError(err)
	ms.NotNil(u.ID, "User ID is generated when saved to DB.")

	// Test with incomplete data.
	u2 := &User{
		FirstName: "Thomas",
		LastName:  "Edison",
	}
	verrs, err = db.ValidateAndCreate(u2)
	ms.NoError(err)
	ms.True(verrs.HasAny(), "User cannot be created without age field.")
}

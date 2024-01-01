package models

func (ms *ModelSuite) Test_User() {
	u := &User{
		FirstName:            "Nikola",
		LastName:             "Tesla",
		Age:                  86,
		Email:                "nikola@tesla.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}

	ms.Equal("Nikola Tesla", u.FullName(), "FullName returns user name.")

	db := ms.DB
	verrs, err := u.Create(db)
	ms.NoError(err)
	ms.NotNil(u.ID, "User ID is generated when saved to DB.")
	ms.False(verrs.HasAny(), "No validation errors")

	// Test with incomplete data.
	u2 := &User{
		FirstName: "Thomas",
		LastName:  "Edison",
	}
	verrs, err = db.ValidateAndCreate(u2)
	ms.NoError(err)
	ms.True(verrs.HasAny(), "User cannot be created without age field.")
}

func (ms *ModelSuite) Test_UserAddress() {
	u := &User{
		FirstName:            "Nikola",
		LastName:             "Tesla",
		Age:                  25,
		Email:                "nikola@tesla.com",
		Password:             "password",
		PasswordConfirmation: "password",
		UserAddress: Address{
			Street: "1 Main Street",
			City:   "Everytown",
			State:  "IL",
			Zip:    "12345",
		},
	}

	db := ms.DB
	err := db.Eager().Create(u)
	ms.NoError(err)

	ms.NotNil(u.UserAddress.ID, "Address saved along with User.")

	u2 := User{}
	err = db.Find(&u2, u.ID)
	ms.NoError(err)

	ms.Empty(u2.UserAddress, "User address not loaded by default")
	u2.GetAddress(db)
	ms.NotEmpty(u2.UserAddress, "GetAddress loads user address")
}

func (ms *ModelSuite) Test_UserBlogs() {
	u := &User{
		FirstName:            "Nikola",
		LastName:             "Tesla",
		Age:                  25,
		Email:                "nikola@tesla.com",
		Password:             "password",
		PasswordConfirmation: "password",
		Blogs: Blogs{Blog{
			Title: "First blog",
			Body:  "<p>Interesting content",
		}},
	}

	db := ms.DB
	err := db.Eager().Create(u)
	ms.NoError(err)

	u2 := User{}
	err = db.Find(&u2, u.ID)
	ms.NoError(err)

	ms.Empty(u2.Blogs, "Blogs not loaded with user by default.")
	ms.Equal("Nikola Tesla", u2.FullName(), "Confirm the correct user is loaded.")

	err = u2.GetBlogs(db)
	ms.NoError(err)

	ms.Len(u2.Blogs, 1, "GetBlogs loads the user's blogs.")
}

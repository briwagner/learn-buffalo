package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// BlogsTag is used by pop to map your blogs_tags database table to your go code.
type BlogsTag struct {
	ID        uuid.UUID `json:"id" db:"id"`
	BlogID    uuid.UUID `db:"blog_id"`
	Blog      *Blog     `belongs_to:"blogs"`
	TagID     uuid.UUID `db:"tag_id"`
	Tag       *Tag      `belongs_to:"tags"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (b BlogsTag) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// BlogsTags is not required by pop and may be deleted
type BlogsTags []BlogsTag

// String is not required by pop and may be deleted
func (b BlogsTags) String() string {
	jb, _ := json.Marshal(b)
	return string(jb)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (b *BlogsTag) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (b *BlogsTag) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (b *BlogsTag) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

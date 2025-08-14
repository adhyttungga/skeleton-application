package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string    `validate:"required" json:"name" bson:"name,omitempty"`
	Email     string    `validate:"required,email" json:"email" bson:"email,omitempty"`
	Password  string    `validate:"required" json:"password" bson:"password,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

func (u *User) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}

	u.UpdatedAt = time.Now()
	type my User
	return bson.Marshal((*my)(u))
}

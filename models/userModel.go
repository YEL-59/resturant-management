package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID `bson: "_id"`
	First_name    *string            `json:"quantity" validate:"required,min=2,max=100"`
	Last_name     *float64           `json:"quantity" validate:"required,min=2,max=100"`
	Password      *string            `json:"password " validate:"required,min=6"  `
	Email         *string            `json:"email"  validate:"email,required"`
	Avatar        *string            `json:"avatar"  `
	Phone         *string            `json:"phone"  `
	Token         *string            `json:"token"  `
	Refresh_Token *string            `json:"refresh_token"  `

	Created_at time.Time `json:"created_at"  `
	Update_at  time.Time `json:"update_at "  `
	User_id    string    `json:"user_id " `
}

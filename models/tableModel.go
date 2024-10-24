package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Table struct {
	Id               primitive.ObjectID `bson: "_id"`
	Number_of_guests *int               `json:"number_of_guests" validate:"required"`
	Table_number     *int               `json:"table_number" validate:"required"`

	Created_at time.Time `json:"created_at"  `
	Update_at  time.Time `json:"update_at "  `
	Table_id   string    `json:"table_id "  `
}

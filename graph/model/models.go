package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           string  `json:"id" bson:"_id"`
	CategoryName string  `json:"categoryName" bson:"categoryName"`
	Name         string  `json:"name" bson:"name"`
	Price        int     `json:"price" bson:"price"`
	Brand        *string `json:"brand" bson:"brand"`
	Description  *string `json:"description" bson:"description"`
	SellerUsername     string  `json:"sellerUsername" bson:"sellerUsername"`
}

func (p *Product) GetBSON() (interface{}, error) {
	if p.ID == "" {
		p.ID = primitive.NewObjectID().Hex()
	}
	type me *Product
	return me(p), nil
}

type Category struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type User struct {
	ID          string    `json:"id" bson:"_id"`
	Username    string    `json:"username" bson:"username"`
	Email       string    `json:"email" bson:"email"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	Password    string    `json:"password" bson:"password"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}

func (u *User) GetBSON() (interface{}, error) {
	if u.ID == "" {
		u.ID = primitive.NewObjectID().Hex()
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now().UTC()
	}
	type my *User
	return my(u), nil
}
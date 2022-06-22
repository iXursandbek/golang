package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first___name" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"last___name"  validate:"required,min=2,max=30"`
	Password        *string            `json:"password"     validate:"required,min=6"`
	Email           *string            `json:"email" 		validate:"email, required`
	Phone           *string            `json:"phone" 		validate:"required"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh___token"`
	Created_At      time.Time          `json:"created___at"`
	Updated_At      time.Time          `json:"updated___at"`
	User_ID         string             `json:"user___id"`
	UserCart        []ProductUser      `json:"u_ser_cart" bson:"u_ser_cart"`
	Address_Details []Addres           `json:"address___details" bson:"address___details"`
	Order_Status    []Order            `json:"order___status" bson:"order___status"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"product___id"`
	Product_Name *string            `json:"product___name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"product___id"`
	Product_Name *string            `json:"product___name"`
	Price        int                `json:"price"`
	Rating       *uint              `json:"rating"`
	Image        *string            `json:"image"`
}

type Addres struct {
	Addres_ID primitive.ObjectID `bson:"addres___id" json:"addres___id"`
	House     *string            `json:"house" bson:"house"`
	Street    *string            `json:"street" bson:"street"`
	City      *string            `json:"city" bson:"city"`
	Pincode   *string            `json:"pincode" bson:"pincode"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"order___id"`
	Order_Cart     []ProductUser      `json:"order___cart" bson:"order___cart"`
	Ordered_At     time.Time          `json:"ordered___at" bson:"ordered___at"`
	Price          int                `json:"price" bson:"price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment___method" bson:"payment___method"`
}

type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}

package models


type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
	Grade string `json:"grade"  bson:"grade"`
}

type Employee struct {
	Name    string  `json:"emp_name" bson:"emp_name"`
	Age     int     `json:"age" bson:"age"`
	Address Address `json:"address" bson:"address"`
}


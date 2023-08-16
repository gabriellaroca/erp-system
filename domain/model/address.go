package model

type Address struct {
	ZipCode    string `json:"zip_code" bson:"zip_code"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	Street     string `json:"street" bson:"street"`
	Number     string `json:"number" bson:"number"`
	District   string `json:"district" bson:"district"`
	Complement string `json:"complement" bson:"complement"`
}

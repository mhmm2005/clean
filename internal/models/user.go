package models

type User struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	Key      string `json:"key" bson:"key"`
	Days     int    `json:"days" bson:"days"`
	Comment  string `json:"comment" bson:"comment"`
	Status   bool   `json:"status" bson:"status"`
	Blocked  string `json:"blocked" bson:"blocked"`
	TempKey  string `json:"temp_key" bson:"temp_key"`
}

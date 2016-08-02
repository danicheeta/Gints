package models

import "time"

type User struct {
	Email     string `bson:"_id,inline"` //Email
	Password  string
	Activated bool
}

type Profile struct {
	Email      string `bson:"_id,inline"` //Email
	UserName   string
	FirstName  string
	LastName   string
	Avatar     string
	Gender     string
	BirthDate  time.Time
	Bio        string
	Theme      string
	Games      []string
	Gints      []string
	Geners     []string
	Friends    []string
	Achivemnts []string
	Google     string
	Twitter    string
	Github     string
	Score      float32
}

type Admin struct {
	GameName string `bson:"_id,inline"` //Email
	Admins   []string
}

type Master struct {
	Id string `bson:"_id,inline"` //Email
}

type Level struct {
	Id       string `bson:"_id,inline"`
	MinScore float32
	Name     string
}

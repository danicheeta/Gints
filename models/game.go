package models

type Game struct {
	Name        string   `bson:"_id"`
	Geners      []string //genre names
	Achivements []string
	Banner      string //url
	Thumbnail   string
	Description string
	Pinned      string //gint id
	Users       int
	K           float32
}

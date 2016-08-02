package models

import "github.com/maxwellhealth/bongo"

type Gint struct {
	bongo.DocumentBase `bson:",inline"`
	Email              string //user id
	Hint               string
	Game               string //game id
	Approves           int
	Declines           int
	Inappropriate      bool
	Cheat              bool
}

type Approve struct {
	bongo.DocumentBase `bson:",inline"`
	Email              string //user id
	Gint               string //gint id
}

type Decline struct {
	bongo.DocumentBase `bson:",inline"`
	Email              string //user id
	Gint               string //gint id
}

type Hashtags struct { // collection per hashtag!
	bongo.DocumentBase `bson:",inline"`
}

type GameGints struct { // collection per game
	bongo.DocumentBase `bson:",inline"`
}

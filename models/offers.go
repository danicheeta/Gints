package models

import "github.com/maxwellhealth/bongo"

type GameOffers struct {
	bongo.DocumentBase `bson:",inline"`
	Title              string
	Games              []string
}

type GintOffers struct {
	bongo.DocumentBase `bson:",inline"`
	Title              string
	Gints              []string
}

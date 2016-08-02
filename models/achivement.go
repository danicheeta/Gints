package models

import "github.com/maxwellhealth/bongo"

type Achivement struct {
	bongo.DocumentBase `bson:",inline"`
	Title              string
	Items              []string
}

type AchivementChecks struct {
	bongo.DocumentBase `bson:",inline"`
	Email              string
	Achivement         string
}

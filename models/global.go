package models

import "github.com/maxwellhealth/bongo"

type ForceSettings struct {
	bongo.DocumentBase `bson:",inline"`
	Theme              string
	Texture            string
}

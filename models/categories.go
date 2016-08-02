package models

type Categorie struct {
	Name string `bson:"_id"`
	Sub  []string
}

type SubCategorie struct {
	Name  string `bson:"_id"`
	Games []string
}

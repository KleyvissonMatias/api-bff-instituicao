package model

import "gopkg.in/mgo.v2/bson"

type InstituicaoEnsino struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Nome     string        `bson:"name" json:"nome"`
	Sigla    string        `bson:"sigla" json:"sigla"`
	Contato  string        `bson:"contato" json:"contato"`
	Endereco *Endereco     `bson:"endereco" json:"endereco"`
}

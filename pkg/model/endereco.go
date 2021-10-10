package model

type Endereco struct {
	Rua         string `bson:"rua" json:"rua"`
	Bairro      string `bson:"bairro" json:"bairro"`
	Cidade      string `bson:"cidade" json:"cidade"`
	Uf          string `bson:"uf" json:"uf"`
	Complemento string `bson:"complemento" json:"complemento"`
	Numero      int    `bson:"numero" json:"numero"`
}

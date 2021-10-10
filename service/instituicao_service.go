package service

import (
	"log"

	model "api-bff-instituicao/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ServiceDb struct {
	server   string
	database string
	user     string
	pword    string
}

var mongoDb *mgo.Database

const collection = "instituicao_ensino_db"

func (s *ServiceDb) Connect() {
	s.database = "instituicao"
	s.user = "user"
	s.pword = "root"
	s.server = "mongodb://127.0.0.1:27017"
	session, error := mgo.Dial(s.server)
	if error != nil {
		log.Fatal(error)
	}
	mongoDb = session.DB(s.database)
}

func (s *ServiceDb) GetAll() ([]model.InstituicaoEnsino, error) {
	var instituicaoEnsino []model.InstituicaoEnsino
	err := mongoDb.C(collection).Find(bson.M{}).All(&instituicaoEnsino)
	return instituicaoEnsino, err
}

func (s *ServiceDb) GetByID(id string) (model.InstituicaoEnsino, error) {
	var instituicaoEnsino model.InstituicaoEnsino
	err := mongoDb.C(collection).FindId(bson.ObjectIdHex(id)).One(&instituicaoEnsino)
	return instituicaoEnsino, err
}

func (s *ServiceDb) Create(instituicaoEnsino model.InstituicaoEnsino) error {
	err := mongoDb.C(collection).Insert(&instituicaoEnsino)
	return err
}

func (s *ServiceDb) Delete(id string) error {
	err := mongoDb.C(collection).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (s *ServiceDb) Update(id string, instituicaoEnsino model.InstituicaoEnsino) error {
	err := mongoDb.C(collection).UpdateId(bson.ObjectIdHex(id), &instituicaoEnsino)
	return err
}

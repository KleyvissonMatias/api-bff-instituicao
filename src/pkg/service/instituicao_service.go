package service

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ServiceDb struct {
	Server   string
	Database string
}

var mongoDb *mgo.Database

const collection = "instituicao_ensino"

func (s *ServiceDb) Connect() {
	session, error := mgo.Dial(s.Server)
	if error != nil {
		log.Fatal(error)
	}
	mongoDb = session.DB(s.Database)
}

func (s *ServiceDb) GetAll() ([]InstituicaoEnsino, error) {
	var instituicaoEnsino []InstituicaoEnsino
	err := mongoDb.C(collection).Find(bson.M{}).All(&instituicaoEnsino)
	return instituicaoEnsino, err
}

func (s *ServiceDb) GetByID(id string) (InstituicaoEnsino, error) {
	var instituicaoEnsino InstituicaoEnsino
	err := mongoDb.C(collection).FindId(bson.ObjectIdHex(id)).One(&instituicaoEnsino)
	return instituicaoEnsino, err
}

func (s *ServiceDb) Create(instituicaoEnsino InstituicaoEnsino) error {
	err := mongoDb.C(collection).Insert(&instituicaoEnsino)
	return err
}

func (s *ServiceDb) Delete(id string) error {
	err := mongoDb.C(collection).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (s *ServiceDb) Update(id string, instituicaoEnsino InstituicaoEnsino) error {
	err := mongoDb.C(collection).UpdateId(bson.ObjectIdHex(id), &instituicaoEnsino)
	return err
}

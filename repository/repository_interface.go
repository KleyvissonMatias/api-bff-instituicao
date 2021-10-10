package repository

type MongoRepository interface {
	GetAll() ([]interface{}, error)
	GetByID(id string) (interface{}, error)
	Create(user interface{}) error
	Delete(id string) error
	Update(id string, user interface{}) error
}

package storage

// Store is a storage store
type Store interface {
	Client() interface{}
	IsExist(id int64, name string, v interface{}) bool
	Get(id int64, name string, v interface{}) error
	Put(name string, v interface{}) error
	Update(id int64, name string, has interface{}, v interface{}) error
	Remove(id int64, name string, v interface{}) error
	List(size, current int, v interface{}) (int, error)
	GetAssociation(association string, v ...interface{}) error
}

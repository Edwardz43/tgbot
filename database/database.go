package db

// DB is the interface of database
type DB interface {
	Connect(connStr string)
	InsertOne(doc interface{}) (interface{}, error)
	InsertMany(docs []interface{}) error
	Update(filter, data interface{}) (interface{}, error)
	Upsert(filter, data interface{}) (interface{}, error)
	UpdateMany(filter, data interface{}) (interface{}, error)
	Contains(doc interface{}) (bool, error)
}

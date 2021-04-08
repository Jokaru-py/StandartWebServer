package storage

import (
	"database/sql"
	_ "github.com/lib/pq" //Чтобы отработала init
	"log"
)

//Instance
type Storage struct {
	config *Config //Как подключиться к БД
	db *sql.DB //через что делать запросы
}

//Storage Const
func New(config *Config) *Storage  {
	return &Storage{
		config: config,
	}
}

func (storage *Storage) Open() error  {
	db, err := sql.Open("postgres",storage.config.DatabaseURI)
	if err != nil{
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db

	log.Println("Database connection created successfully!")
	return nil
}

func (storage *Storage) Close()  {
	storage.db.Close()
}
package repositories

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sannonthachai/test/database"
	"github.com/sannonthachai/test/domains"
)

type Repository struct {
	Repo *database.Database
}

func NewRepo(db *database.Database) *Repository {
	var schema domains.TableSchema
	db.Db.AutoMigrate(&schema)
	return &Repository{
		Repo: db,
	}
}

func (repo *Repository) Create(code string, price int) {
	repo.Repo.Db.Create(&domains.TableSchema{
		Code:  code,
		Price: price,
	})
}

func (repo *Repository) FindAll() []domains.TableSchema {
	schemas := []domains.TableSchema{}
	repo.Repo.Db.Find(&schemas)
	return schemas
}

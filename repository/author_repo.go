package repository

import (
	"database/sql"
	"library-api/model"
	"log"
)

type AuthorRepository interface {
	GetAll() ([]model.Author, error)
	GetById(id int) (model.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

// GetAll implements AuthorRepository.
func (a *authorRepository) GetAll() ([]model.Author, error) {
	var authors []model.Author
	query := "SELECT * FROM authors"

	rows, err := a.db.Query(query)
	if err != nil {
		log.Println("authorRepository.Query", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var author model.Author

		if err := rows.Scan(&author.Id, &author.Name, &author.Bio); err!=nil{
		log.Println("authorRepository.Scan", err.Error())
			return nil, err
		}
		

		authors = append(authors, author)
	}

 	return authors, nil
}

// GetById implements AuthorRepository.
func (a *authorRepository) GetById(id int) (model.Author, error) {
	panic("unimplemented")
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{
		db: db,
	}
}

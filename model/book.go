package model

type Book struct {
	Id            int      `json:"id"`
	Title         string   `json:"title"`
	AuthorID      *int     `json:"-"` // foreign key
	Author        Author   `json:"author"`
	CategoryID    *int     `json:"-"` // foreign key
	Category      Category `json:"category"`
	PublishedYear int      `json:"published_year"`
	Price         float64  `json:"price"`
	Stock         int      `json:"stock"`
}

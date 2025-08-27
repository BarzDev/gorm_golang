package model

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
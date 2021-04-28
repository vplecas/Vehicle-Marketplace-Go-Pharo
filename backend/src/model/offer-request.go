package model

import "time"

type OfferRequest struct {
	// offer properties
	Id          string    `json:"id"`
	Price       int       `json:"price"`
	PublishDate time.Time `json:"publishDate"` // publish date = now
	Location    string    `json:"location"`
	Rates       []Rate    `json:"rates"`
	Comments    []Comment `json:"comments"`

	// vehicle properties
	Make           string    `json:"make"`
	ModelCar       string    `json:"modelCar"`
	ProductionDate time.Time `json:"productionDate"` // production date
	HP             int       `json:"hp"`
	Cubic          int       `json:"cubic"`
}

package entites

import "github.com/google/uuid"

type Item struct {
	Id       uuid.UUID
	Make     string
	Model    string
	Year     int
	Color    string
	Mileage  int
	ImageUrl string
	//Auction   Auction
	//AuctionId uuid.UUID
}

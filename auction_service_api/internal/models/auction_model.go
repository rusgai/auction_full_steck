package models

import (
	"time"

	"github.com/google/uuid"
)

type AuctionModel struct {
	Id             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	AuctionEnd     time.Time
	Seller         string
	Winner         string
	Make           string
	Model          string
	Year           int
	Color          string
	Mileage        int
	ImageUrl       string
	Status         string
	ReservePrice   int
	SoldAmount     int
	CurrentHighBid int
}
type CreateAuctionModel struct {
	Make         string
	Model        string
	Color        string
	Mileage      int
	Year         int
	ReservePrice int
	ImageUrl     string
	AuctionEnd   time.Time
}
type UpdateAuctionModel struct {
	Make    string
	Model   string
	Color   string
	Mileage int
	Year    int
}

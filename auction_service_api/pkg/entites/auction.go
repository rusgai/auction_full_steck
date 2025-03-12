package entites

import (
	"time"

	"github.com/google/uuid"
)

type Auction struct {
	ID             uuid.UUID
	ReservePrice   int
	Seller         string
	Winner         string
	SoldAmount     int
	CurrentHighBid int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	AuctionEnd     time.Time
	Status
	Item
}

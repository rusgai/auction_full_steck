package repositories

import (
	"auction_service_api/pkg/entites"
	"context"

	"github.com/google/uuid"
)

type AuctionRepository interface {
	GetAll(ctx context.Context) []entites.Auction
	GetItem(ctx context.Context, id uuid.UUID) (error *entites.Auction)
	Create(ctx context.Context, auction entites.Auction) error
	Update(ctx context.Context, auction entites.Auction) error
	Delete(ctx context.Context, id uuid.UUID) error
}

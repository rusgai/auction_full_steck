package impl

import (
	"auction_service_api/internal/repositories"
	"auction_service_api/pkg/entites"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type auctionRepositoryimpl struct {
	db sqlx.DB
}

func NewAuctionRepositoryImpl(db sqlx.DB) repositories.AuctionRepository {
	return &auctionRepositoryimpl{db}
}

func (auctionrepo *auctionRepositoryimpl) GetAll(ctx context.Context) []entites.Auction {
	query := `SELECT * FROM auction
			  INNER JOIN item 
			  ON auction.auction_id = item.item_id`
	var auctions []entites.Auction
	err := auctionrepo.db.SelectContext(ctx, auctions, query)
	if err != nil {
		return auctions
	}
	return nil
}
func (auctionrepo *auctionRepositoryimpl) GetItem(ctx context.Context, id uuid.UUID) (error *entites.Auction) {
	return nil
}
func (auctionrepo *auctionRepositoryimpl) Create(ctx context.Context, auction entites.Auction) error {
	return nil
}
func (auctionrepo *auctionRepositoryimpl) Update(ctx context.Context, auction entites.Auction) error {
	return nil
}
func (auctionrepo *auctionRepositoryimpl) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

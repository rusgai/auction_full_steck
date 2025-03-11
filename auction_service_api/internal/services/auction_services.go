package services

import (
	"auction_service_api/internal/models"
	"context"

	"github.com/google/uuid"
)

type AuctionService interface {
	GetAll(ctx context.Context) []models.AuctionModel
	GetItem(ctx context.Context, id uuid.UUID) (error *models.AuctionModel)
	Create(ctx context.Context, auction models.AuctionModel) error
	Update(ctx context.Context, auction models.AuctionModel) error
	Delete(ctx context.Context, id uuid.UUID) error
}

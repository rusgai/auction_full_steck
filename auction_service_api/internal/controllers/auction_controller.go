package controllers

import (
	"auction_service_api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type auctionController struct {
	s services.AuctionService
}

func NewAuctionController(s services.AuctionService) *auctionController {
	return &auctionController{s}
}
func (auction *auctionController) Router(rout *fiber.App) {
	rout.Get("/api/v1/auction", auction.getAll)
	rout.Get("/api/v1/auction/{id}", auction.getItem)
	rout.Post("/api/v1/auction", auction.create)
	rout.Put("/api/v1/auction", auction.update)
	rout.Delete("/api/v1/auction", auction.delete)
}
func (auction *auctionController) getAll(c *fiber.Ctx) error {
	ctx := c.Context()
	getAuctions := auction.s.GetAll(ctx)
	err := c.JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   getAuctions,
	})
	if err != nil {
		c.JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err,
		})
	}
	return nil
}
func (auction *auctionController) getItem(c *fiber.Ctx) error {
	return nil
}
func (auction *auctionController) create(c *fiber.Ctx) error {
	return nil
}
func (auction *auctionController) update(c *fiber.Ctx) error {
	return nil
}
func (auction *auctionController) delete(c *fiber.Ctx) error {
	return nil
}

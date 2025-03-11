package controllers

import "github.com/gofiber/fiber/v2"

type auctionController struct{}

func NewAuctionController() *auctionController {
	return &auctionController{}
}
func (auction *auctionController) Router(rout *fiber.App) {
	rout.Get("/api/v1/auction", auction.getAll)
}
func (auction *auctionController) getAll(c *fiber.Ctx) error {

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

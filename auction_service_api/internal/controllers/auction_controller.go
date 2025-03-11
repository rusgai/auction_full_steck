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
	err := c.JSON(fiber.Map{
		"status":  200,
		"massage": "Hemo massage",
	})
	if err != nil {
		return err
	}
	return nil
}

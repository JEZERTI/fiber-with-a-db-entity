package handlers

import (
	"fitness_club/models"
	"fitness_club/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateClient(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid json",
		})
	}
	client, err := services.UpdateClient(id, data)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(), // вытащили ошибку из userServices
		})
	}
	return c.JSON(client)

}

func GetClients(c *fiber.Ctx) error {
	clients := services.GetClients()
	limitstr := c.Query("limit")
	limit, err := strconv.Atoi(limitstr)

	if err != nil || limit <= 0 {
		limit = len(clients)
	}
	return c.Status(200).JSON(clients[:limit])
}

func GetClientByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	client, err := services.GetClientByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(client)
}

func CreateClient(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid json",
		})
	}
	createdClient := services.CreateClient(client)
	return c.Status(201).JSON(createdClient)
}

func DeleteClient(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	err = services.DeleteClient(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(404).JSON(fiber.Map{
		"error": "client not found",
	})
}

func UpdateClient2(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var updated models.Client
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid json",
		})
	}

	client, err := services.UpdateClient2(id, updated)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(client)
}

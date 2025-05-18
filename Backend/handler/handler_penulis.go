package handler

import (
	"Backend/controller"
	"Backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAllPenulis(c *fiber.Ctx) error {
	penulis, err := controller.GetAllPenulis(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data penulis",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil semua data penulis",
		"data":    penulis,
	})
}

func CreatePenulis(c *fiber.Ctx) error {
	var p model.Penulis

	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal parsing body",
			"error":   err.Error(),
		})
	}

	if err := controller.CreatePenulis(&p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat penulis",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Penulis berhasil dibuat",
		"data":    p,
	})
}

func GetPenulisByID(c *fiber.Ctx) error {
	id := c.Params("id")
	p, err := controller.GetPenulisByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Penulis ditemukan",
		"data":    p,
	})
}

func UpdatePenulisByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var p model.Penulis

	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parsing body",
		})
	}

	err := controller.UpdatePenulisByID(c.Context(), id, p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal update penulis: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Penulis berhasil diperbarui",
	})
}

func DeletePenulisByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := controller.DeletePenulisByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal hapus penulis: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Penulis dengan ID %s berhasil dihapus", id),
	})
}

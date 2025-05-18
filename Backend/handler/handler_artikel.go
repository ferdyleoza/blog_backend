package handler

import (
	"Backend/controller"
	"Backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

	func Homepage(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Home Page!")
	}

	func GetAllArtikels(c *fiber.Ctx) error {
		// Call the controller function to get all artikels
		artikels, err := controller.GetAllArtikels(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch artikels",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Fetched all artikels",
			"data":    artikels,
		})
	}
	
	func CreateArtikel(c *fiber.Ctx) error {
	var artikel model.Artikel

	if err := c.BodyParser(&artikel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal parsing request body",
			"error":   err.Error(),
		})
	}

	if err := controller.CreateArtikel(&artikel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan artikel",
			"error":   err.Error(),
		})
	}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Artikel berhasil dibuat",
			"data":    artikel,
		})
	}
	
	func GetArtikelByID(c *fiber.Ctx) error {
	id := c.Params("id")

	artikel, err := controller.GetArtikelByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Artikel ditemukan",
		"data":    artikel,
	})
}

// handler/artikel_handler.go
func UpdateArtikelByID(c *fiber.Ctx) error {
	id := c.Params("id")

	// ⬇️ Print raw body dari request
	fmt.Println("Raw body:", string(c.Body()))

	var updatedArtikel model.Artikel

	// ⬇️ Parse body ke struct
	if err := c.BodyParser(&updatedArtikel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parse body",
		})
	}

	// ⬇️ Print hasil parsing ke struct
	fmt.Printf("Parsed data: %+v\n", updatedArtikel)

	// Lanjut update
	err := controller.UpdateArtikelByID(c.Context(), id, updatedArtikel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal update artikel: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Artikel berhasil diperbarui",
	})
}

func DeleteArtikelByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Request DELETE masuk, ID:", id) // <- cek ini keluar nggak

	err := controller.DeleteArtikelByID(c.Context(), id)
	if err != nil {
		fmt.Println("Delete error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal hapus artikel: %v", err),
		})
	}

	fmt.Println("Delete berhasil") // <- ini juga
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Artikel dengan ID %s berhasil dihapus", id),
	})
}







	


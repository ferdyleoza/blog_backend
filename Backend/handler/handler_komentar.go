package handler

import (
	"Backend/controller"
	"Backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAllKomentars(c *fiber.Ctx) error {
	komentars, err := controller.GetAllKomentars(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil komentar",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil semua komentar",
		"data":    komentars,
	})
}

func CreateKomentar(c *fiber.Ctx) error {
	var komentar model.Komentar

	if err := c.BodyParser(&komentar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal parsing request body",
			"error":   err.Error(),
		})
	}

	if err := controller.CreateKomentar(&komentar); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan komentar",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Komentar berhasil dibuat",
		"data":    komentar,
	})
}

func GetKomentarByID(c *fiber.Ctx) error {
	id := c.Params("id")

	komentar, err := controller.GetKomentarByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Komentar ditemukan",
		"data":    komentar,
	})
}

func UpdateKomentarByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Raw body:", string(c.Body()))

	var updatedKomentar model.Komentar

	if err := c.BodyParser(&updatedKomentar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parse body",
		})
	}

	err := controller.UpdateKomentarByID(c.Context(), id, updatedKomentar)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal update komentar: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Komentar berhasil diperbarui",
	})
}

func DeleteKomentarByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Hapus komentar ID:", id)

	err := controller.DeleteKomentarByID(c.Context(), id)
	if err != nil {
		fmt.Println("Delete error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal hapus komentar: %v", err),
		})
	}

	fmt.Println("Komentar berhasil dihapus")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Komentar dengan ID %s berhasil dihapus", id),
	})
}

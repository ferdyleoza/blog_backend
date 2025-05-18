package handler

import (
	"Backend/controller"
	"Backend/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetAllKategoris(c *fiber.Ctx) error {
	kategoris, err := controller.GetAllKategoris(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data kategori",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil mengambil semua kategori",
		"data":    kategoris,
	})
}

func CreateKategori(c *fiber.Ctx) error {
	var kategori model.Kategori

	if err := c.BodyParser(&kategori); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gagal parsing body",
			"error":   err.Error(),
		})
	}

	if err := controller.CreateKategori(&kategori); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan kategori",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Kategori berhasil dibuat",
		"data":    kategori,
	})
}

func GetKategoriByID(c *fiber.Ctx) error {
	id := c.Params("id")

	kategori, err := controller.GetKategoriByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Kategori ditemukan",
		"data":    kategori,
	})
}

func UpdateKategoriByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Update ID:", id)

	var kategori model.Kategori

	if err := c.BodyParser(&kategori); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parse body",
		})
	}

	err := controller.UpdateKategoriByID(c.Context(), id, kategori)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal update kategori: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Kategori berhasil diperbarui",
	})
}

func DeleteKategoriByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("Hapus kategori ID:", id)

	err := controller.DeleteKategoriByID(c.Context(), id)
	if err != nil {
		fmt.Println("Delete kategori error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Gagal hapus kategori: %v", err),
		})
	}

	fmt.Println("Kategori berhasil dihapus")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Kategori dengan ID %s berhasil dihapus", id),
	})
}

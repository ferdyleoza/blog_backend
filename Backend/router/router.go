package router

import (
	"Backend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Routing untuk artikel
	api.Get("/", handler.Homepage)
	api.Get("/artikels", handler.GetAllArtikels)
	app.Post("/api/artikels", handler.CreateArtikel)
	api.Get("/artikels/:id", handler.GetArtikelByID)
	app.Put("/api/artikels/:id", handler.UpdateArtikelByID)
	app.Delete("/api/artikels/:id", handler.DeleteArtikelByID)

	// Routing untuk kategori
	app.Get("/api/kategoris", handler.GetAllKategoris)
	app.Post("/api/kategoris", handler.CreateKategori)
	app.Get("/api/kategoris/:id", handler.GetKategoriByID)
	app.Put("/api/kategoris/:id", handler.UpdateKategoriByID)
	app.Delete("/api/kategoris/:id", handler.DeleteKategoriByID)

		// Routing untuk komentar
	app.Get("/api/komentars", handler.GetAllKomentars)
	app.Post("/api/komentars", handler.CreateKomentar)
	app.Get("/api/komentars/:id", handler.GetKomentarByID)
	app.Put("/api/komentars/:id", handler.UpdateKomentarByID)
	app.Delete("/api/komentars/:id", handler.DeleteKomentarByID)

	// Routing untuk penulis
	app.Get("/api/penulis", handler.GetAllPenulis)
	app.Post("/api/penulis", handler.CreatePenulis)
	app.Get("/api/penulis/:id", handler.GetPenulisByID)
	app.Put("/api/penulis/:id", handler.UpdatePenulisByID)
	app.Delete("/api/penulis/:id", handler.DeletePenulisByID)
	

}
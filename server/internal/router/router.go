package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/latoulicious/SIPP/internal/config"
	"github.com/latoulicious/SIPP/internal/database"
	"github.com/latoulicious/SIPP/internal/handler"
	"github.com/latoulicious/SIPP/internal/middleware"
	"github.com/latoulicious/SIPP/internal/repository"
	"github.com/latoulicious/SIPP/internal/service"
)

func SetupRoutes(app *fiber.App, e *casbin.Enforcer) {
	app.Use(cors.New())

	api := app.Group("/api", logger.New())

	// Initialize services , repositories and handlers

	// Core Usage
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	mapelRepository := repository.NewMapelRepository(database.DB)
	mapelService := service.NewMapelService(mapelRepository)
	mapelHandler := handler.NewMapelHandler(mapelService)
	kelasRepository := repository.NewKelasRepository(database.DB)
	kelasService := service.NewKelasService(kelasRepository)
	kelasHandler := handler.NewKelasHandler(kelasService)
	tahunRepository := repository.NewTahunRepository(database.DB)
	tahunService := service.NewTahunService(tahunRepository)
	tahunHandler := handler.NewTahunHandler(tahunService)
	jurusanRepository := repository.NewJurusanRepository(database.DB)
	jurusanService := service.NewJurusanService(jurusanRepository)
	jurusanHandler := handler.NewJurusanHandler(jurusanService)

	// General Usage
	capaianRepository := repository.NewCapaianRepository(database.DB)
	capaianService := service.NewCapaianService(capaianRepository)
	capaianHandler := handler.NewCapaianHandler(capaianService)
	alurRepository := repository.NewAlurRepository(database.DB)
	alurService := service.NewAlurService(alurRepository)
	alurHandler := handler.NewAlurHandler(alurService)
	modulRepository := repository.NewModulRepository(database.DB)
	modulService := service.NewModulService(modulRepository)
	modulHandler := handler.NewModulHandler(modulService)
	// penilaianRepository := repository.NewPenilaianRepository(database.DB)
	// penilaianService := service.NewPenilaianService(penilaianRepository)
	// penilaianHandler := handler.NewPenilaianHandler(penilaianService)
	kognitifRepository := repository.NewKognitifRepository(database.DB)
	kognitifService := service.NewKognitifService(kognitifRepository)
	kognitifHandler := handler.NewKognitifHandler(kognitifService)
	formatifRepository := repository.NewFormatifRepository(database.DB)
	formatifService := service.NewFormatifService(formatifRepository)
	formatifHandler := handler.NewFormatifHandler(formatifService)
	sumatifRepository := repository.NewSumatifRepository(database.DB)
	sumatifService := service.NewSumatifService(sumatifRepository)
	sumatifHandler := handler.NewSumatifHandler(sumatifService)
	remedialRepository := repository.NewRemedialRepository(database.DB)
	remedialService := service.NewRemedialService(remedialRepository)
	remedialHandler := handler.NewRemedialHandler(remedialService)
	pengayaanRepository := repository.NewPengayaanRepository(database.DB)
	pengayaanService := service.NewPengayaanService(pengayaanRepository)
	pengayaanHandler := handler.NewPengayaanHandler(pengayaanService)
	soalRepository := repository.NewSoalRepository(database.DB)
	soalService := service.NewSoalService(soalRepository)
	soalHandler := handler.NewSoalHandler(soalService)
	itemRepository := repository.NewItemRepository(database.DB)
	itemService := service.NewItemService(itemRepository)
	itemHandler := handler.NewItemHandler(itemService)
	bankRepository := repository.NewBankRepository(database.DB)
	bankService := service.NewBankService(bankRepository)
	bankHandler := handler.NewBankHandler(bankService)

	// Use UserService for AuthHandler
	authHandler := handler.NewAuthHandler(userService)

	// Core Usage routes

	// User routes with authentication and RBAC middleware
	userRoutes := api.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware(e, userRepository))
	userRoutes.Use(middleware.RBACMiddleware(e, userRepository))
	userRoutes.Get("/", userHandler.GetUsers)
	userRoutes.Get("/:id", userHandler.GetUserByID)
	userRoutes.Post("/", userHandler.CreateUser)
	userRoutes.Put("/:id", userHandler.UpdateUser)
	userRoutes.Delete("/:id", userHandler.DeleteUser)

	// Add the new route for password change
	userRoutes.Post("/:id/change-password", userHandler.ChangePassword)

	// Authentication routes
	authRoutes := api.Group("/auth")
	authRoutes.Post("/login", authHandler.LoginHandler)

	// Mapel routes
	mapelRoutes := api.Group("/mapel")
	mapelRoutes.Get("/", mapelHandler.GetMapel)
	mapelRoutes.Get("/:id", mapelHandler.GetMapelByID)
	mapelRoutes.Post("/", mapelHandler.CreateMapel)
	mapelRoutes.Put("/:id", mapelHandler.UpdateMapel)
	mapelRoutes.Delete("/:id", mapelHandler.DeleteMapel)

	// Kelas routes
	kelasRoutes := api.Group("/kelas")
	kelasRoutes.Get("/", kelasHandler.GetKelas)
	kelasRoutes.Get("/:id", kelasHandler.GetKelasByID)
	kelasRoutes.Post("/", kelasHandler.CreateKelas)
	kelasRoutes.Put("/:id", kelasHandler.UpdateKelas)
	kelasRoutes.Delete("/:id", kelasHandler.DeleteKelas)

	// Tahun routes
	tahunRoutes := api.Group("/tahun")
	tahunRoutes.Get("/", tahunHandler.GetTahun)
	tahunRoutes.Get("/:id", tahunHandler.GetTahunByID)
	tahunRoutes.Post("/", tahunHandler.CreateTahun)
	tahunRoutes.Put("/:id", tahunHandler.UpdateTahun)
	tahunRoutes.Delete("/:id", tahunHandler.DeleteTahun)

	// Jurusan routes
	jurusanRoutes := api.Group("/jurusan")
	jurusanRoutes.Get("/", jurusanHandler.GetJurusan)
	jurusanRoutes.Get("/:id", jurusanHandler.GetJurusanByID)
	jurusanRoutes.Post("/", jurusanHandler.CreateJurusan)
	jurusanRoutes.Put("/:id", jurusanHandler.UpdateJurusan)
	jurusanRoutes.Delete("/:id", jurusanHandler.DeleteJurusan)

	// General usage routes

	// Capaian routes
	capaianRoutes := api.Group("/capaian")
	capaianRoutes.Get("/", capaianHandler.GetCapaian)
	capaianRoutes.Get("/:id", capaianHandler.GetCapaianByID)
	capaianRoutes.Post("/", capaianHandler.CreateCapaian)
	capaianRoutes.Put("/:id", capaianHandler.UpdateCapaian)
	capaianRoutes.Delete("/:id", capaianHandler.DeleteCapaian)

	// Alur routes
	alurRoutes := api.Group("/alur")
	alurRoutes.Get("/", alurHandler.GetAlur)
	alurRoutes.Get("/:id", alurHandler.GetAlurByID)
	alurRoutes.Post("/", alurHandler.CreateAlur)
	alurRoutes.Put("/:id", alurHandler.UpdateAlur)
	alurRoutes.Delete("/:id", alurHandler.DeleteAlur)

	// Modul routes
	modulRoutes := api.Group("/modul")
	modulRoutes.Get("/", modulHandler.GetModul)
	modulRoutes.Get("/:id", modulHandler.GetModulByID)
	modulRoutes.Post("/", modulHandler.CreateModul)
	modulRoutes.Put("/:id", modulHandler.UpdateModul)
	modulRoutes.Delete("/:id", modulHandler.DeleteModul)

	// Penilaian routes
	// penilaianRoutes := api.Group("/penilaian")
	// penilaianRoutes.Get("/", penilaianHandler.GetPenilaian)
	// penilaianRoutes.Get("/:id", penilaianHandler.GetPenilaianByID)
	// penilaianRoutes.Post("/", penilaianHandler.CreatePenilaian)
	// penilaianRoutes.Put("/:id", penilaianHandler.UpdatePenilaian)
	// penilaianRoutes.Delete("/:id", penilaianHandler.DeletePenilaian)

	// Kognitif routes
	kognitifRoutes := api.Group("/kognitif")
	kognitifRoutes.Get("/", kognitifHandler.GetKognitif)
	kognitifRoutes.Get("/:id", kognitifHandler.GetKognitifByID)
	kognitifRoutes.Post("/", kognitifHandler.CreateKognitif)
	kognitifRoutes.Put("/:id", kognitifHandler.UpdateKognitif)
	kognitifRoutes.Delete("/:id", kognitifHandler.DeleteKognitif)

	// Formatif routes
	formatifRoutes := api.Group("/formatif")
	formatifRoutes.Get("/", formatifHandler.GetFormatif)
	formatifRoutes.Get("/:id", formatifHandler.GetFormatifByID)
	formatifRoutes.Post("/", formatifHandler.CreateFormatif)
	formatifRoutes.Put("/:id", formatifHandler.UpdateFormatif)
	formatifRoutes.Delete("/:id", formatifHandler.DeleteFormatif)

	// Sumatif routes
	sumatifRoutes := api.Group("/sumatif")
	sumatifRoutes.Get("/", sumatifHandler.GetSumatif)
	sumatifRoutes.Get("/:id", sumatifHandler.GetSumatifByID)
	sumatifRoutes.Post("/", sumatifHandler.CreateSumatif)
	sumatifRoutes.Put("/:id", sumatifHandler.UpdateSumatif)
	sumatifRoutes.Delete("/:id", sumatifHandler.DeleteSumatif)

	// Remedial routes
	remedialRoutes := api.Group("/remedial")
	remedialRoutes.Get("/", remedialHandler.GetRemedial)
	remedialRoutes.Get("/:id", remedialHandler.GetRemedialByID)
	remedialRoutes.Post("/", remedialHandler.CreateRemedial)
	remedialRoutes.Put("/:id", remedialHandler.UpdateRemedial)
	remedialRoutes.Delete("/:id", remedialHandler.DeleteRemedial)

	// Pengayaan routes
	pengayaanRoutes := api.Group("/pengayaan")
	pengayaanRoutes.Get("/", pengayaanHandler.GetPengayaan)
	pengayaanRoutes.Get("/:id", pengayaanHandler.GetPengayaanByID)
	pengayaanRoutes.Post("/", pengayaanHandler.CreatePengayaan)
	pengayaanRoutes.Put("/:id", pengayaanHandler.UpdatePengayaan)
	pengayaanRoutes.Delete("/:id", pengayaanHandler.DeletePengayaan)

	// Soal routes
	soalRoutes := api.Group("/soal")
	soalRoutes.Get("/", soalHandler.GetSoal)
	soalRoutes.Get("/:id", soalHandler.GetSoalByID)
	soalRoutes.Post("/", soalHandler.CreateSoal)
	soalRoutes.Put("/:id", soalHandler.UpdateSoal)
	soalRoutes.Delete("/:id", soalHandler.DeleteSoal)

	// ItemSoal routes
	itemRoutes := api.Group("/item")
	itemRoutes.Get("/", itemHandler.GetItem)
	itemRoutes.Get("/:id", itemHandler.GetItemByID)
	itemRoutes.Post("/", itemHandler.CreateItem)
	itemRoutes.Put("/:id", itemHandler.UpdateItem)
	itemRoutes.Delete("/:id", itemHandler.DeleteItem)

	// BankSoal routes
	bankRoutes := api.Group("/bank")
	bankRoutes.Get("/", bankHandler.GetBank)
	bankRoutes.Get("/:id", bankHandler.GetBankByID)
	bankRoutes.Post("/", bankHandler.CreateBank)
	bankRoutes.Put("/:id", bankHandler.UpdateBank)
	bankRoutes.Delete("/:id", bankHandler.DeleteBank)

	// Public Route
	publicRoutes := api.Group("/public")
	publicRoutes.Get("/user", userHandler.GetAllUsersPublic)
	publicRoutes.Get("/tahun", tahunHandler.GetTahunPublic)
	publicRoutes.Get("/kelas", kelasHandler.GetKelasPublic)
	publicRoutes.Get("/mapel", mapelHandler.GetMapelPublic)
	publicRoutes.Get("/jurusan", jurusanHandler.GetJurusanPublic)

	// total count routes
	totalCountRoutes := api.Group("/total")
	totalCountRoutes.Get("/kognitif", kognitifHandler.CountRelatedQuestionsHandler)
	totalCountRoutes.Get("/formatif", formatifHandler.CountRelatedQuestionsHandler)
	totalCountRoutes.Get("/sumatif", sumatifHandler.CountRelatedQuestionsHandler)
	totalCountRoutes.Get("/pengayaan", pengayaanHandler.CountRelatedQuestionsHandler)
	totalCountRoutes.Get("/remedial", remedialHandler.CountRelatedQuestionsHandler)

	// misc routes
	app.Get("/api/get-jwt-secret", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"jwtSecret": config.JwtSecret})
	})

}

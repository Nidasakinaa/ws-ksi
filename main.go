package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/Nidasakinaa/ws-ksi/handler"
)

func main() {
    app := fiber.New()

    // Tambahkan middleware CORS
    app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://127.0.0.1:5500", // Hanya izinkan frontend tertentu
        AllowMethods:     "GET, POST, PUT, DELETE",
        AllowHeaders:     "Content-Type, Authorization",
        AllowCredentials: true, // 🔥 Wajib agar cookie bisa dikirim
    }))

    // Rute lainnya
    app.Post("/login", handler.Login)
    app.Post("/customer-login", handler.CustomerLogin)
    app.Post("/logout", handler.Logout)
    app.Post("/register", handler.Register)
    app.Get("/dashboard", handler.DashboardPage)

    app.Listen(":3000")
}
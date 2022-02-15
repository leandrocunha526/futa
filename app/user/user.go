package user

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/leandrocunha526/futa/app/middleware"
	"github.com/leandrocunha526/futa/app/user/events/login"
	"github.com/leandrocunha526/futa/app/user/events/register"
	"github.com/leandrocunha526/futa/app/user/events/ws"
	"github.com/leandrocunha526/futa/app/user/events/ws/handler"
)

func Init(app *fiber.App, db *sql.DB) {
	hub := ws.NewHub()
	go hub.Run()
	app.Get("/rooms/:roomId", func(c *fiber.Ctx) error {
		return handler.GetClientInRoom(c, hub)
	})

	app.Get("/ws/:roomId", handler.JoinRoom(hub))

	app.Post("/ws/register", func(c *fiber.Ctx) error {
		return register.Handler(c, db)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return login.Handler(c, db)
	})

	app.Post("/ws", middleware.JWTAuth, func(c *fiber.Ctx) error {
		return handler.CreateRoom(c, hub)
	})

	app.Get("/ws", func(c *fiber.Ctx) error {
		return handler.GetAvailableRoom(c, hub)
	})
}

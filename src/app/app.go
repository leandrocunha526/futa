package app

import (
    "github.com/gofiber/fiber"
    "github.com/gofiber/fiber/middleware/cors"
    "github.com/gofiber/fiber/middleware/logger"
    "github.com/gofiber/fiber/middleware/recovery"
    "github.com/leandrocunha526/moechat/app/user"
	"github.com/leandrocunha526/moechat/app/user/events/ws"
)

func Run() {
  app := fiber.New()
  db :=DbConn()
  hub := ws.NewHub()
  go hub.Run()
  
  app.Use(cors.New())
  app.Use(logger.New())
  app.Use(recovery.New())
  
  user.Init(app, db)
  
  app.Get("/", func(c *fiber.Ctx) error {
      return c.JSON("hello") 
})
  
  app.Listen(":3005")
}

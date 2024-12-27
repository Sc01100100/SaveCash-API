package middlewares

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"https://nekowawolf.xyz",
	"https://nekowawolf.github.io",
	"https://sc01100100.github.io",
}

var Internalhost string = os.Getenv("INTERNALHOST") + ":" + os.Getenv("PORT")

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowHeaders:     "Origin,Login,Content-Type",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
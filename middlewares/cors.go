package middlewares

import (
    "strings"

    "github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
    "https://nekowawolf.xyz",
    "https://nekowawolf.github.io",
    "https://sc01100100.github.io",
}

var Cors = cors.Config{
    AllowOrigins:     strings.Join(origins[:], ","),
    AllowHeaders:     "Origin, Content-Type, Accept",
    ExposeHeaders:    "Content-Length",
    AllowCredentials: true,
}
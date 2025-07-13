package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "email-service/internal/config"
    "email-service/internal/routes"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No se encontró el archivo .env, usando variables del entorno")
    }

    config.ConnectDB()

    router := gin.Default()
    routes.RegisterRoutes(router)

    port := os.Getenv("PORT")
    if port == "" {
        port = "4003"
    }

    log.Printf("🚀 Servidor corriendo en el puerto %s", port)
    router.Run(":" + port)
}

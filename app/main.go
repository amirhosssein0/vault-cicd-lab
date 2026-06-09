package main

import (
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

var version = "1.0.3"

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "up",
        })
    })

    r.GET("/version", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "version": version,
            "env":     os.Getenv("APP_ENV"),
        })
    })

    r.GET("/ready", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ready",
        })
    })

    r.GET("/secret", func(c *gin.Context) {
        data, err := os.ReadFile("/vault/secrets/db-password")
        if err != nil {
            c.JSON(http.StatusOK, gin.H{
                "secret": "secret not injected yet",
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "secret": string(data),
        })
    })

    r.Run(":8080")
}
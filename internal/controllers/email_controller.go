package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "email-service/internal/services"
)

type EmailRequest struct {
    To      string `json:"to"`
    Subject string `json:"subject"`
    Body    string `json:"body"`
}

func SendEmail(c *gin.Context) {
    var req EmailRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.SendEmail(req.To, req.Subject, req.Body); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Email enviado con éxito"})
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type PredictResponse struct {
	Input  []float64 `json:"input"`
	Output []float64 `json:"output"`
}

func main() {
	r := gin.Default()

	// Enable CORS for frontend (allow localhost:5175 and Tauri app)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5175", "http://127.0.0.1:5175"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Go backend is running!"})
	})

	// Call Python AI service
	r.GET("/ai/predict", func(c *gin.Context) {
		resp, err := http.Get("http://127.0.0.1:5001/predict")
		if err != nil {
			log.Println("Error calling Python service:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "AI service unavailable"})
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read AI response"})
			return
		}

		var predictResp PredictResponse
		if err := json.Unmarshal(body, &predictResp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid AI response"})
			return
		}

		c.JSON(http.StatusOK, predictResp)
	})

	r.Run(":8080")
}

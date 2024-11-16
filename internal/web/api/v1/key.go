package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

var verificationKeys = make(map[string]string)

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

func generateRandomKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateKey(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	key := generateRandomKey()
	verificationKeys[username] = key

	c.JSON(http.StatusOK, gin.H{"key": key})
}

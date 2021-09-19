package main

import (
	handler "github.com/MahmudulTushar/Adda-Backend-Go-Msg/http"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	allowOrigin := os.Getenv("ALLOW_ORIGIN")

	if port == "" {
		port = defaultPort
	}
	if allowOrigin == "" {
		allowOrigin = "http://localhost:3000"
	}

	server := gin.Default()

	server.GET("/", handler.PlaygroundHandler())
	server.POST("/query", handler.GraphqlHandler())

	socketServer := handler.CreateSocketServer()
	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer socketServer.Close()

	server.Use(GinMiddleware(allowOrigin))
	server.GET("/socket.io/", gin.WrapH(socketServer))
	server.POST("/socket.io/", gin.WrapH(socketServer))
	server.StaticFS("/public", http.Dir("./asset"))

	if err := server.Run(":" + port); err != nil {
		log.Fatal("failed run app: ", err)
	}
}

package server

import (
	"github.com/GGmaz/wallet/user-service/config"
	"github.com/GGmaz/wallet/user-service/internal/ctx"
	"github.com/GGmaz/wallet/user-service/internal/db"
	v1 "github.com/GGmaz/wallet/user-service/internal/server/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	config *config.Config
}

func dbMiddleware(gormDB *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(ctx.Transaction, gormDB)
		c.Next()
	}
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	r := gin.Default()
	r.Use(gin.Logger())

	gormDB, err := db.Init(server.config.Db)

	if err != nil {
		log.Fatal("Could not connect to the database" + err.Error())
		return
	}

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(dbMiddleware(gormDB))
	r.Use(CORSMiddleware())

	v1.RegisterVersion(r)

	err = r.Run(":" + server.config.Port)

	if err != nil {
		log.Fatal("Could not start the server" + err.Error())
		return
	}

	println("Starting server on port: " + server.config.Port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

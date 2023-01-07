package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guidoenr/fulbo/model/psdb"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

const (
	SecretKey = "secret"
)

// StartRouter turns on the gin-gonic server and initialize the entire REST-API
// routes and endpoints
func StartRouter(db *psdb.PostgreDB) {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//router.LoadHTMLGlob("api/templates/*")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET, POST"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// initializing the db to use the connection pool
	err := db.InitDB()
	if err != nil {
		log.Error().Msgf("initializing db: %v", err)
	}

	defer db.CloseDB()

	// players routes
	players := router.Group("/players")
	{
		players.GET("/", showPlayers(db))
		players.GET("/:id", showPlayerByID(db))
		players.POST("/:id/create", createPlayer(db))
	}

	router.Run()
}

// -------------------------- CONTROLLERS

// showPlayers is the main page for the players
func showPlayers(db *psdb.PostgreDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		players, err := GetPlayers(db)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"players": players})
	}
}

// showPlayerByID find the player given the id(PK)
func showPlayerByID(db *psdb.PostgreDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		player, err := GetPlayerByID(db, c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"player": player})
	}
}

// showPlayerByNickname find the player given the nickname
func showPlayerByNickname(db *psdb.PostgreDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		player, err := GetPlayerByNickname(db, c.Param("name"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"player": player})
	}
}

// createPlayer creates a player
func createPlayer(db *psdb.PostgreDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO, this will be one of the latest thing to code
		player, err := CreatePlayer(db)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"ok": player})
	}
}

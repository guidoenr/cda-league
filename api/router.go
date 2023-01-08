package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guidoenr/fulbo/model"
	"github.com/guidoenr/fulbo/model/psdb"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

// StartRouter turns on the gin-gonic server and initialize the entire REST-API
// routes and endpoints
func StartRouter(postgreDB *psdb.PostgreDB) {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// setting the cors config
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET, POST"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// initializing the postgreDB to use the connection pool
	err := postgreDB.InitDB()
	if err != nil {
		log.Error().Msgf("initializing postgreDB: %v", err)
	}
	defer postgreDB.CloseDB()

	// players routes
	players := router.Group("/players")
	{
		players.GET("/", showPlayers(postgreDB))
		players.GET("/:id", showPlayerByID(postgreDB))
		players.GET("/nickname/:nickname", showPlayerByNickname(postgreDB))
		players.POST("/:id/create", createPlayer(postgreDB))
		players.POST("/:id/update", updatePlayer(postgreDB))
	}

	// match routes
	match := router.Group("/match")
	{
		match.GET("/", generateMatch(postgreDB))
	}

	router.Run()
}

// -------------------------- CONTROLLERS
// -------------------------- PLAYERS

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
		player, err := GetPlayerByNickname(db, c.Param("nickname"))
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

// createPlayer creates a player
func updatePlayer(db *psdb.PostgreDB) gin.HandlerFunc {
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

// -------------------------- CONTROLLERS
// -------------------------- MATCH
// createPlayer creates a player
func generateMatch(db *psdb.PostgreDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		// ok, here we have to obtain the list of players
		// that the user put in the website, then with that
		// list of players we gonna generate a match
		// suppose this list of players is the one that the user send us
		players, _ := GetPlayers(db)

		var match model.Match
		match.Init(players)
		match.GenerateTeams()

		var err error

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"team1": match.Team1, "team2": match.Team2})
	}
}

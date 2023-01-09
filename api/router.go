package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/guidoenr/fulbo/model"
	"github.com/guidoenr/fulbo/model/psdb"
	"net/http"
	"time"
)

type Router struct {
	db         *psdb.PostgreDB
	controller *PlayerControler
	ginRouter  *gin.Engine
}

func (r *Router) Init(postgreDB *psdb.PostgreDB) {
	var controller PlayerControler
	controller.Init(postgreDB)

	// gin.SetMode(gin.ReleaseMode) TODO:later
	r.ginRouter = gin.Default()
	r.db = postgreDB
	r.controller = &controller
}

// StartRouter turns on the gin-gonic server and initialize the entire REST-API
// routes and endpoints
func (r *Router) StartRouter() {

	// setting the cors config
	r.ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET, POST"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// players routes
	players := r.ginRouter.Group("/players")
	{
		players.GET("/", r.showPlayers())
		players.GET("/rank", r.showPlayersRanked())
		players.GET("/:id", r.showPlayerByID())
		players.GET("/nickname/:nickname", r.showPlayerByNickname())
		players.POST("/:id/create", r.createPlayer())
		players.POST("/:id/update", r.updatePlayer())
	}

	// match routes
	match := r.ginRouter.Group("/match")
	{
		match.GET("/", r.generateMatch())
	}

	r.ginRouter.Run()
}

// -------------------------- CONTROLLERS
// -------------------------- PLAYERS

// showPlayers is the main page for the players
func (r *Router) showPlayers() gin.HandlerFunc {
	return func(c *gin.Context) {
		players, err := r.controller.GetPlayers()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"players": players})
	}
}

// showPlayersRanked returns the list of players ordered by their elo-gamesWon-goalsPerMatch
func (r *Router) showPlayersRanked() gin.HandlerFunc {
	return func(c *gin.Context) {
		players, err := r.controller.GetPlayersRankedByElo()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"players": players})
	}
}

// showPlayerByID find the player given the id(PK)
func (r *Router) showPlayerByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		player, err := r.controller.GetPlayerByID(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"player": player})
	}
}

// showPlayerByNickname find the player given the nickname
func (r *Router) showPlayerByNickname() gin.HandlerFunc {
	return func(c *gin.Context) {
		player, err := r.controller.GetPlayerByNickname(c.Param("nickname"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"player": player})
	}
}

// createPlayer creates a player
func (r *Router) createPlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO, this will be one of the latest thing to code
		// you will need a param inside the CreatePlayer() func for sure
		player, err := r.controller.CreatePlayer()
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"ok": player})
	}
}

// createPlayer creates a player
func (r *Router) updatePlayer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO, this will be one of the latest thing to code
		player, err := r.controller.CreatePlayer()
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
func (r *Router) generateMatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		// ok, here we have to obtain the list of players
		// that the user put in the website, then with that
		// list of players we gonna generate a match
		// suppose this list of players is the one that the user send us
		players, _ := r.controller.GetPlayers()

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

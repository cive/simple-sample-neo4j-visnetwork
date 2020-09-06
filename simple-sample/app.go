package main

import (
	"log"
	"net/http"
	"time"

	"github.com/cive/simple-sample-neo4j-visnetwork/core"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Neo struct {
	conf *core.Neo4jConfig
}

func NewNeo() *Neo {
	conf := core.NewDB()
	if conf != nil {
		return &Neo{conf}
	}
	return nil
}

func (neo Neo) GetActedInReq(c *gin.Context) {
	data := neo.conf.GetActedIn()
	c.JSON(http.StatusOK, data)
}

func main() {
	time.Sleep(time.Second * 10)
	neo := NewNeo()
	if neo == nil {
		log.Fatal("neo4j config is invalid.")
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/acted_in", neo.GetActedInReq)
	r.Run(":8080")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run()
}

// setupRouter creates the routing of the Books API
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/:operation/:num1/:num2", getOperation)
	return router

}

func getOperation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": nil})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

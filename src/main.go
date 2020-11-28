package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run()
}

// setupRouter creates the routing of the Books API
func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/:operation/:num1/:num2", getOperation)
	return router

}

func getOperation(c *gin.Context) {
	// Get parameters
	operation := c.Param("operation")
	num1, _ := strconv.Atoi(c.Param("num1"))
	num2, _ := strconv.Atoi(c.Param("num2"))

	switch operation {
	case "sum":
		c.HTML(http.StatusOK, "result.html",
			gin.H{"operation": operation,
				"num1":   num1,
				"num2":   num2,
				"result": add(num1, num2),
			},
		)
	case "multiply":
		c.HTML(http.StatusOK, "result.html",
			gin.H{"operation": operation,
				"num1":   num1,
				"num2":   num2,
				"result": add(num1, num2),
			},
		)
	default:
		c.HTML(http.StatusOK,
			"invalid-route.html",
			gin.H{"result": "Invalid operation"},
		)
	}

}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

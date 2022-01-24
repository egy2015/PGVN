package main
import (
	"github.com/gin-gonic/gin"
 	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", home)

	router.Run(":8888")

}

func home(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"title": "Home",
		"subtitle": "This is homepage",
	})

}
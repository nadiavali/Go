package main

import (
	"net/http"
	"github.com/gin-goinc/gin"
)

type pk struct {
	ID string `json:"id"`
	Item string `json: "item"`
}

func GetPk(c *gin.Cotext) {
	n = 10
	res, err := http.Get(url)[:n+1]

}

url := "https://pokeapi.co/api/v2/item/"
func main() {
	router := gin.Default()

	router.GET("url/n", Getpk)

	router.run("localhost:80:80")
}
package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"

	"github.com/gin-gonic/gin"
)


func getFbPhoto() string {
	url := "http://graph.facebook.com/899295334/picture?type=large&redirect=false"
	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil { log.Fatal("NewRequest: ", reqErr); return "" }
	client := &http.Client{}
	res, doErr := client.Do(req)
	if doErr != nil { log.Fatal("Do: ", doErr); return "" }
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil { log.Fatal(readErr); return "" }
	type Record struct {
		Data struct {
			IsSilhouette bool `json:""is_silhouette`
			URL string `json:"url"`
		} `json:"data"`
	}
	var rec Record
	jsonErr := json.Unmarshal(body, &rec)
	if jsonErr != nil { log.Fatal(jsonErr); return "" }
	return rec.Data.URL
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"photoLink": getFbPhoto(),
		})
	})

	router.Run(":" + port)
}

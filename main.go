package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"database/sql"
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
			IsSilhouette bool `json:"is_silhouette"`
			URL string `json:"url"`
		} `json:"data"`
	}
	var rec Record
	jsonErr := json.Unmarshal(body, &rec)
	if jsonErr != nil { log.Fatal(jsonErr); return "" }
	return rec.Data.URL
}

func main() {
	db, err := sql.Open("postgres", "user=duyhainguyen dbname=gotesque sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT username FROM users")
	fmt.Println(rows)
	port := os.Getenv("PORT")
	port = "5000"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		description := "Hello my name is Duy Hai NGUYEN"
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"photoLink": getFbPhoto(),
			"description": description,
		})
	})

	router.Run(":" + port)
}

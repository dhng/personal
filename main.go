package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"fmt"
	"github.com/yanatan16/golang-instagram/instagram"
	"net/url"
)

func DoSomeInstagramApiStuff(accessToken string) {
	api := New("", accessToken)

	if ok, err := api.VerifyCredentials(); !ok {
		panic(err)
	}

	var myId string

	// Get yourself!
	if resp, err := api.GetSelf(); err != nil {
		panic(err)
	} else {
	    // A response has two fields: Meta which you shouldn't really care about
	    // And whatever your getting, in this case, a User
		me := resp.User
		fmt.Printf("My userid is %s and I have %d followers\n", me.Id, me.Counts.FollowedBy)
	}

	params := url.Values{}
	params.Set("count", "1")
	if resp, err := api.GetUserRecentMedia("self" /* this works :) */, params); err != nil {
		panic(err)
	} else {
	    if len(resp.Medias) == 0 { // [sic]
	    	panic("I should have some sort of media posted on instagram!")
	    }
	    media := resp.Medias[0]
	    fmt.Println("My last media was a %s with %d comments and %d likes. (url: %s)\n", media.Type, media.Comments.Count, media.Like.Count, media.Link)
	}
}

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

package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AutoComplete(c *gin.Context) {
	title, ok := c.GetQuery("title")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing query parameter"})
		return
	}

	url := "https://online-movie-database.p.rapidapi.com/auto-complete?q=" + title

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "a553b227a6msh537fce1b58b7c67p15d455jsn53ea9024e2c0")
	req.Header.Add("X-RapidAPI-Host", "online-movie-database.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)


	/*
		get all response
		@ var v interface{}
	*/

	// get spesific properti object
	type Result struct{
		D []interface{} 
		Q string
		V int
	}
	var r Result

	json.NewDecoder(res.Body).Decode(&r) //decode response

	if r.D == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Movie not found"})
		return 
	}

	c.JSON(200, r)
}


func DetailMovie(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": ok})
		return
	}

	url := "https://online-movie-database.p.rapidapi.com/title/get-details?tconst=" + id

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "a553b227a6msh537fce1b58b7c67p15d455jsn53ea9024e2c0")
	req.Header.Add("X-RapidAPI-Host", "online-movie-database.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// type ResultId struct {
	// 	Image []interface{}
	// }
	// // var v interface{}
	// var resultId ResultId
	// json.NewDecoder(res.Body).Decode(&resultId) //decode response

	// if resultId.Image == nil{
	// 	c.IndentedJSON(201, gin.H{"message": resultId})
	// 	return
	// }

	c.JSON(200, body)
	fmt.Println(body)
}
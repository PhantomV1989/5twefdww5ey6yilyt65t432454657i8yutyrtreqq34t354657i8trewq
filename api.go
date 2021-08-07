package main

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

type PostReq struct {
	FileName    string `json:"file_name"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/image", uploadImage)
	router.GET("/images", getImages)
	router.GET("/image/:id", getImage)
	router.PATCH("/image", patchImage) //only once
	return router
}

func uploadImage(c *gin.Context) {
	var postReq PostReq
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	} else {
		err := json.Unmarshal(body, &postReq)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
		} else {
			c.JSON(200, gin.H{
				"id":          "",
				"description": postReq.Description,
				"available":   "true",
				"code":        "200",
			})
		}
	}
}

func getImages(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": []interface{}{},
	})
}

func getImage(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":          "",
		"description": "",
		"available":   "true",
		"code":        "200",
	})
}

func patchImage(c *gin.Context) {
	var postReq PostReq
	c.Bind(&postReq)
	c.JSON(200, gin.H{
		"id":          "",
		"description": postReq.Description,
		"available":   "true",
		"code":        "200",
	})
}

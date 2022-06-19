package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	fmt.Println("GetArticles 111111")
}

func GetArticle(c *gin.Context) {
	fmt.Println("GetArticle 222222222")
}

func AddArticle(c *gin.Context) {
	fmt.Println("AddArticle 3333333")
}

func UpdateArticle(c *gin.Context) {
	fmt.Println("UpdateArticle 44444444")
}

func DeleteArticle(c *gin.Context) {
	fmt.Println("DeleteArticle 5555555")
}

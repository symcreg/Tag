package controller

import (
	"github.com/gin-gonic/gin"
)

type Tag struct {
	Id          int      `json:"id"`
	UserId      int      `json:"userId"`
	ImgId       int      `json:"imgId"`
	tags        []string `json:"tags"`
	Description string   `json:"description"`
}

func AddTagHandler(c *gin.Context) {

}
func PutTagHandler(c *gin.Context) {

}
func DeleteTagHandler(c *gin.Context) {

}

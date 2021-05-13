package oez

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": Config.Common.Title,
	})
}

//func Favicon(c *gin.Context)  {
//	c.File("static/favicon.ico")
//}

func CreateUrl(c *gin.Context) {
	var url Url
	if err := c.ShouldBind(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "错误的请求."})
		return
	}
	if match, _ := regexp.MatchString("(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]", url.Url); !match {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "错误的URL."})
		return
	}
	url.ClientIP = c.ClientIP()
	id, err := url.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "系统错误."})
		return
	}
	if c.GetHeader("Content-Type") == "application/x-www-form-urlencoded" {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": Config.Common.Title,
			"id":    Encode10To62(id),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": Encode10To62(id)})
}

func GetUrl(c *gin.Context) {
	id := c.Param("id")
	if id == "favicon.ico" {
		c.File("static/favicon.ico")
		return
	}
	url ,err := Get(id)
	if err == nil {
		c.Redirect(http.StatusFound, url.Url)
		url.Look()
	} else {
		c.String(http.StatusNotFound, err.Error())
	}
}

func GetUrlJSON(c *gin.Context) {
	id := c.Param("id")
	url ,err := Get(id)
	if err == nil {
		c.JSON(http.StatusOK, url)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
	}
}

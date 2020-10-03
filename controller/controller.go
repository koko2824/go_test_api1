package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/koko2824/first_postgresql/db"
	"net/http"
	"strconv"
)

type Controller struct {}

func (Controller) GetAll(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"status":"OK",
		"users": db.GetAll(),
	})
}

func (Controller) GetOne(c *gin.Context)  {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":"OK",
		"user": db.GetOne(id),
	})
}

func (Controller) Insert(c *gin.Context)  {
	name := c.PostForm("name")
	role := c.PostForm("role")
	db.Insert(name,role)

	c.JSON(http.StatusOK, gin.H{
		"status":"OK",
		"created user": map[string]string{
			"name":name,
			"role":role,
		},
	})
}

func (Controller) Update(c *gin.Context)  {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	name := c.PostForm("name")
	role := c.PostForm("role")
	db.Update(id, name, role)

	c.JSON(http.StatusOK, gin.H{
		"status":"ok",
		"updated user": map[string]string{
			"id":n,
			"name":name,
			"role":role,
		},
	})
}

func (Controller) Delete(c *gin.Context)  {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	deletedUser := db.GetOne(id)
	db.Delete(id)

	c.JSON(http.StatusOK, gin.H{
		"status":"ok",
		"deleted user": deletedUser,
	})
}
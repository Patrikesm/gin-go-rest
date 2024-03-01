package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patrike-miranda/gin-go-rest/database"
	"github.com/patrike-miranda/gin-go-rest/models"
)

func GetAll(c *gin.Context) {
	var students []models.Student

	//Realiza um select all com base na definição anterior
	database.DB.Find(&students)

	c.JSON(200, students)
}

func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API says": "Olá " + name + ", tudo bem?",
	})
}

func New(c *gin.Context) {
	var student models.Student

	err := c.ShouldBindJSON(&student)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if err := models.ValidateData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)

	c.JSON(200, student)
}

func GetOne(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func Update(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.First(&student, id)

	fmt.Println(student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Student not found",
		})
		return
	}

	err := c.ShouldBindJSON(&student)

	fmt.Println(student)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if err := models.ValidateData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	database.DB.Save(&student)

	fmt.Println(student)

	c.JSON(http.StatusOK, student)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	var student models.Student

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Student deleted",
	})
}

func GetByDocument(c *gin.Context) {
	var student []models.Student
	document := c.Param("document")

	database.DB.Where(&models.Student{CPF: document}).Find(&student)

	if len(student) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func RenderizeIndexPage(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func NotFoundRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", nil)
}

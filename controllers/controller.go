package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rjribeiro/goBackend/database"
	"github.com/rjribeiro/goBackend/models"
	"log"
	"net/http"
)

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200,
		alunos,
	)
}

func GetAlunoById(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{"Not found": "Aluno nao encontrado"},
		)
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func GetAlunoByCPF(c *gin.Context) {
	cpf := c.Param("cpf")
	var aluno models.Aluno
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{"Not found": "Aluno nao encontrado"},
		)
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func CreateAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result := database.DB.Create(&aluno)
	log.Println(result.Error, result.RowsAffected)
	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "aluno deletado"})
}

func UpdateAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

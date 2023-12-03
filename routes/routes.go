package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rjribeiro/goBackend/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/alunos/:id", controllers.GetAlunoById)
	r.GET("alunos/cpf/:cpf", controllers.GetAlunoByCPF)
	r.POST("/alunos", controllers.CreateAluno)
	r.DELETE("alunos/:id", controllers.DeleteAluno)
	r.PATCH("alunos/:id", controllers.UpdateAluno)
	r.Run()
}

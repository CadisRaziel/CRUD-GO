package controller

import (
	"github.com/CadisRaziel/CRUD-GO/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	//retornando o erro que criamos
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")

	c.JSON(err.Code, err)

	//Nosso objeto de erro no postman (sem esse tratamento de erro no postman nao aparece nada.)
	// {
	// 	"message": "Você chamou a rota de forma errada",
	// 	"error": "bad request",
	// 	"code": 400
	// }
}

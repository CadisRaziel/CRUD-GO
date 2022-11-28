package controller

import (
	"fmt"

	"github.com/CadisRaziel/CRUD-GO/src/configuration/rest_err"
	"github.com/CadisRaziel/CRUD-GO/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)

	//retornando o erro que criamos
	//err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")

	//c.JSON(err.Code, err)

	//Nosso objeto de erro no postman (sem esse tratamento de erro no postman nao aparece nada.)
	// {
	// 	"message": "Você chamou a rota de forma errada",
	// 	"error": "bad request",
	// 	"code": 400
	// }
}

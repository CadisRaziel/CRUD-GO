package routes

import (
	"github.com/CadisRaziel/CRUD-GO/src/controller"
	"github.com/gin-gonic/gin"
)

//... -> lembrando tres pontinhos eu posso passar 1 ou N handlers (func (*gin.RouterGroup).GET(relativePath string, handlers ...gin.HandlerFunc)
//Middleware -> é um handler
//função que vai iniciarlizar as rotas
func InitRoutes(r *gin.RouterGroup) {
	//:userId" -> para receber o id pelo parametro
	//r.GET("/getUserById/:userId", func(c *gin.Context) {}, func(c *gin.Context) {}) -> essas 2 func posso fazer um middleware(handler) nelas ou seja posso por quantas eu quise, ali tem um exemplo caso eu queria validar o token jwt ou ver se o usuario tem alguma permissão especial

	//atrelando nossas rotas as funções separadas com o 'controller'
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}

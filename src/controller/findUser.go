package controller

import "github.com/gin-gonic/gin"

//gin.Context -> quando chegar uma requisição ele vai me dar um 'gin.Context'
//dentro dele vamos ter todas nossas informações sobre a request exemplo: request uri, body etc..
func FindUserById(c *gin.Context)    {}
func FindUserByEmail(c *gin.Context) {}

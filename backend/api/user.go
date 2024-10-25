package api

import (	
	"net/http"
	//"fmt"
	"walfare/services"
	"walfare/models"
	"github.com/gin-gonic/gin"
)

//登入ing
func DoLoginHandler(c *gin.Context){
	from := models.User{}

	if err :=c.ShouldBind(&from); err != nil {
		c.IndentedJSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}

	//調用函式去進行驗證，正確會返回一個token
	token, err := services.Login(from.Account, from.Password)
	
	if err != nil {
		c.IndentedJSON(http.StatusOK, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"tokenName": "token", "tokenValue": token})
}

//註冊
func RegisterHandler(c *gin.Context){
	form := models.User{}

	if err := c.ShouldBind(&form); err != nil {
		c.IndentedJSON(http.StatusOK, err.Error())
		return
	}

	if err := services.Register(form); err != nil {
		c.IndentedJSON(http.StatusOK, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, "200")
}
package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	//"fmt"
	"walfare/models"
	"walfare/services"
	"walfare/utils"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

// 登入ing
func DoLoginHandler(c *gin.Context) {
	from := models.User{}

	if err := c.ShouldBindBodyWithJSON(&from); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//調用函式去進行驗證，正確會返回一個token
	token, err := services.Login(from.Account, from.Password)

	// token無效
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"tokenName": "token", "tokenValue": token})
}

// 註冊
func RegisterHandler(c *gin.Context) {
	form := models.User{}

	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := services.Register(&form); err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"msg": form.Email})
}

func GetVerifyEmailHandler(c *gin.Context) {
	email := c.Query("email")
	fmt.Print(email)
	if err := services.GetVerify(email); err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"msg": email})
}

func VerifyEmailHandler(c *gin.Context) {
	form := struct {
		Code  string `json:"code"`
		Email string `json:"email"`
	}{}

	if err := c.ShouldBindBodyWithJSON(&form); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.SetVerify(form.Email, form.Code); err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := user.GetUserByEmail(form.Email); err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": "查詢失敗"})
		return
	}

	const TokenExpireDuration = time.Hour * 24 * 30 * 12 * 30
	newToken, _ := utils.GenerateToken(user.ID, "", TokenExpireDuration)

	c.IndentedJSON(http.StatusOK, gin.H{"msg": newToken})
}

// UpdateuserHandler 处理用户更新请求
func UpdateuserHandler(c *gin.Context) {
	var user models.User

	// 从请求中绑定数据到 map 中
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新用户信息
	if err := user.UpdateUser(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// 返回成功响应
	c.IndentedJSON(http.StatusOK, user)
}

func GetUserByUserIDHandler(c *gin.Context) {
	// 从上下文中获取用户 ID
	userID := c.GetUint("UserID")

	var user models.User
	if err := user.GetUserByID(userID); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func GetBindQrCode(c *gin.Context) {
	// 生成代碼
	code := utils.GenerateCode()

	// 將代碼放入 QR Code 資料
	data := code

	// 生成 QR code 的 PNG 格式
	var png []byte
	var err error
	png, err = qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回 JSON 響應，包含 QR code 以及 code
	c.JSON(http.StatusOK, gin.H{
		"code":  code,                                                                            // 返回生成的代碼
		"image": fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(png)), // 返回 Base64 編碼的圖片
	})
}

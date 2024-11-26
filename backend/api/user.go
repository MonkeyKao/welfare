package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	//"fmt"
	"walfare/models"
	"walfare/services"
	"walfare/utils"

	"github.com/gin-gonic/gin"
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
	userID := c.GetUint("UserID")

	// 从请求中绑定数据到 map 中
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = userID

	// 更新用户信息
	if err := user.UpdateUser(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	if user.Account == "" {
		user.GetUserByID(user.ID)
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

// 上傳頭像處理
func UploadAvatar(c *gin.Context) {
	userID := c.GetUint("UserID")
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件無法處理"})
		return
	}

	// 保存文件到伺服器
	uploadDir := fmt.Sprintf("../uploads/%v", userID)
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}
	// 获取文件扩展名
	ext := filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadDir, fmt.Sprintf("%s.%s", "avatar", ext))
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失敗"})
		return
	}

	var user models.User
	user.ID = userID
	user.Avatar = filePath
	user.UpdateUser()

	c.JSON(http.StatusOK, gin.H{
		"message": "頭像上傳成功",
		"avatar":  fmt.Sprintf("/uploads/%s", file.Filename),
	})
}

func GetAvatar(c *gin.Context) {
	userID := c.GetUint("UserID")

	var user models.User
	user.GetUserByID(userID)

	// 如果用户没有设置头像，使用默认头像
	if user.Avatar == "" {
		user.Avatar = "../uploads/logo.png" // 设置默认头像的路径
	}

	// 读取头像文件
	fileContent, err := os.ReadFile(user.Avatar)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "读取头像文件失败"})
		return
	}

	// 将文件内容转换为 Base64 编码
	base64Content := base64.StdEncoding.EncodeToString(fileContent)

	// 返回 Base64 编码的头像
	c.JSON(http.StatusOK, gin.H{
		"avatar_base64": base64Content,
	})
}

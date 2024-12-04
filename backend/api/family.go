package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"walfare/models"
	"walfare/services"

	"github.com/gin-gonic/gin"
)

func GetFmailyCodeHandler(c *gin.Context) {
	familyID, _ := strconv.ParseUint(c.Param("familyId"), 10, 64)
	code, png := services.GenrateFmailyQRCode(uint(familyID))
	c.JSON(http.StatusOK, gin.H{
		"code":  code,                                                                            // 返回生成的代碼
		"image": fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(png)), // 返回 Base64 編碼的圖片
	})
}

func CreateFamilyMemberHandler(c *gin.Context) {
	userID := c.GetUint("UserID")
	familyID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	familyMember := models.FamilyMember{
		UserID:   userID,
		FamilyID: uint(familyID),
		Role:     "成員",
	}

	if err := familyMember.CreateFamilyMember(); err != nil {
		c.IndentedJSON(http.StatusBadGateway, err)
		return
	}

	c.IndentedJSON(http.StatusOK, "")
}

func CreateFamilyHandler(c *gin.Context) {
	userID := c.GetUint("UserID")
	familyName := c.Param("name")

	family := models.Family{
		FamilyName: familyName,
	}

	if err := family.CreateFamily(); err != nil {
		c.IndentedJSON(http.StatusBadGateway, err)
		return
	}

	familyMember := models.FamilyMember{
		UserID:   userID,
		FamilyID: family.ID,
		Role:     "擁有著",
	}

	familyMember.CreateFamilyMember()

	c.IndentedJSON(http.StatusOK, family)
}

func GetAllFamilyByUserIDHandler(c *gin.Context) {
	userID := c.GetUint("UserID")
	data, err := models.GetAllFamilyByUserID(userID)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, string(data))
}

func JoinFmailyHandler(c *gin.Context) {
	userID := c.GetUint("UserID")
	code := c.Param("code")

	if err := services.JoinFmaily(userID, code); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{})
}

func DeleteFamilyHandler(c *gin.Context) {
	familyID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var family models.Family

	family.ID = uint(familyID)

	if err := family.DeleteFamily(); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{})
}

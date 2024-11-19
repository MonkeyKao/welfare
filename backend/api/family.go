package api

import (
	"net/http"
	"strconv"
	"walfare/models"

	"github.com/gin-gonic/gin"
)

func CreateFamilyMemberHandler(c *gin.Context) {
	userID := c.GetUint("UserID")
	familyID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	familyMember := models.FamilyMember{
		UserID:   userID,
		FamilyID: uint(familyID),
		Role:     "測試",
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
		c.IndentedJSON(http.StatusBadGateway, err)
		return
	}

	c.IndentedJSON(http.StatusOK, string(data))
}

package services

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
	"walfare/models"
	"walfare/utils"
)

var tokenCache = []string{}

func Login(account string, password string) (string, error) {
	var user models.User

	const TokenExpireDuration = time.Hour * 48

	//檢查帳號是否存在
	if err := models.GetUserByAccount(&user, account); err != nil {
		return "", err
	}

	decryptText, err := utils.Decryption(user.Password, user.Salt)
	if err != nil {
		return "", errors.New("解密錯誤")
	}

	if password == decryptText {
		var token string
		if token, err = utils.GenerateToken(user.ID, TokenExpireDuration); err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", errors.New("密碼錯誤")
	}
}

func Register(user models.User) error {
	// 判斷賬號是否已經存在
	if err := models.GetUserByAccount(&models.User{}, user.Account); err == nil {
		return errors.New("賬號已經存在")
	}

	// 加密密碼
	encryptedText, salt, err := utils.Encryption(user.Password)
	if err != nil {
	}
	user.Salt = salt
	user.Password = encryptedText

	const charset = "0123456789" // 字符集
	code := ""                   // 创建一个空字符串用于存储验证码

	for i := 0; i < 6; i++ {
		code += string(charset[rand.Intn(len(charset))]) // 随机从字符集中抽取字符并追加到字符串
	}
	const TokenExpireDuration = time.Minute * 5
	num, err := strconv.ParseUint(code, 10, 32) // ParseUint 将字符串转换为 uint64

	uintNum := uint(num) // 将 uint64 转换为 uint
	token, _ := utils.GenerateToken(uintNum, TokenExpireDuration)
	if err := utils.SendEmail(user.Email, code); err != nil {
		return err
	}

	tokenCache = append(tokenCache, token)
	// if err := models.CreateUser(&user); err != nil {
	// 	return err
	// }
	return nil
}

func verifyEmail(code string) {
	// // 使用 for 循环遍历切片
	// for _, token := range tokenCache {

	// }
}

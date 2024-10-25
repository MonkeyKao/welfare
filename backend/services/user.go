package services

import (
	"errors"
	"walfare/models"
	"walfare/utils"
)

func Login(account string,password string)(string,error){
	var user models.User

	//檢查帳號是否存在
	if err := models.GetUserByAccount(&user, account); err !=nil{
		return "",err
	}

	decryptText,err := utils.Decryption(user.Password,user.Salt)
	if err != nil {
		return "",errors.New("解密錯誤")
	}

	if password == decryptText {
		var token string
		if token, err = utils.GenerateToken(user.Id); err != nil {
			return "",err
		}
		return token, nil
	}else {
		return "",errors.New("密碼錯誤")
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

	if err := models.CreateUser(&user); err != nil {
		return err
	}

	return nil
}
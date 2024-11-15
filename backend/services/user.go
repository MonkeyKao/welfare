package services

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"walfare/models"
	"walfare/utils"
)

var verifyCode = make(map[string]string)
var verifyAccount = []string{}
var mu sync.Mutex

// 取得驗證
func GetVerify(email string) error {
	mu.Lock()

	code := utils.GenerateCode()
	verifyCode[email] = code

	mu.Unlock()

	// if err := utils.SendEmail(email, code); err != nil {
	// 	fmt.Print(err.Error())
	// 	return errors.New("發送email失敗")
	// }

	go func(email, oldCode string) {
		time.Sleep(5 * time.Minute)

		mu.Lock() // 這裡加鎖以確保執行緒安全
		defer mu.Unlock()

		if code, exist := verifyCode[email]; exist && code == oldCode {
			delete(verifyCode, email)
		}
	}(email, code)

	return nil
}

// 設定驗證
func SetVerify(email, receiveCode string) error {
	mu.Lock()
	defer mu.Unlock()
	if oldCode, exist := verifyCode[email]; exist {
		if oldCode == receiveCode {
			fmt.Printf("%s已經新增二級驗證清單", email)
			verifyAccount = append(verifyAccount, email)
			delete(verifyCode, email)
			return nil
		} else {
			return errors.New("驗證碼錯誤")
		}
	} else {
		return errors.New("驗證碼未發送")
	}
}

func IsVerify(email string) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, item := range verifyAccount {
		if item == email {
			return true
		}
	}
	return false
}

func Register(user *models.User) error {
	// 判斷賬號是否已經存在
	if err := user.GetUserByAccount(user.Account); err == nil {
		return errors.New("賬號已經存在")
	}

	// 加密密碼
	encryptedText, salt, err := utils.Encryption(user.Password)
	if err != nil {
	}
	user.Salt = salt
	user.Password = encryptedText

	// 創建賬號
	if err := user.CreateUser(); err != nil {
		return err
	}

	go func(email string, uId uint) {
		time.Sleep(5 * time.Minute)
		mu.Lock()
		defer mu.Unlock()
		flag := true
		for _, v := range verifyAccount {
			if v == email {
				flag = false
				break
			}
		}

		if flag {
			user.DeleteUser()
		}
	}(user.Email, user.ID)

	// 請求二級認證
	if err := GetVerify(user.Email); err != nil {
		return err
	}

	return nil
}

func Login(account string, password string) (string, error) {
	var user models.User
	const TokenExpireDuration = time.Hour * 24 * 30 * 12 * 30

	//檢查帳號是否存在
	if err := user.GetUserByAccount(account); err != nil {
		return "", errors.New("account is not exist")
	}

	// 解密密碼
	decryptText, err := utils.Decryption(user.Password, user.Salt)
	if err != nil {
		return "", err
	}

	// 比對密碼
	if password == decryptText {
		var token string
		if token, err = utils.GenerateToken(user.ID, "000000", TokenExpireDuration); err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", errors.New("密碼錯誤")
	}
}

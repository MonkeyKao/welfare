package services

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"walfare/models"
	"walfare/utils"
)

var tokenCache = make(map[uint]string)

// 用於存儲任務的 map 和一個互斥鎖
var tasks = make(map[uint]chan struct{})
var mu sync.Mutex

func Login(account string, password string) (string, error) {
	var user models.User
	const TokenExpireDuration = time.Hour * 24 * 30 * 12 * 30

	//檢查帳號是否存在
	if err := models.GetUserByAccount(&user, account); err != nil {
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

func Register(user *models.User) (string, error) {
	// 判斷賬號是否已經存在
	if err := models.GetUserByAccount(&models.User{}, user.Account); err == nil {
		return "", errors.New("賬號已經存在")
	}

	// 加密密碼
	encryptedText, salt, err := utils.Encryption(user.Password)
	if err != nil {
	}
	user.Salt = salt
	user.Password = encryptedText

	// 創建賬號
	if err := models.CreateUser(user); err != nil {
		return "", err
	}

	// 生成驗證碼
	const charset = "123456789" // 字符集
	code := ""                  // 创建一个空字符串用于存储验证码

	for i := 0; i < 6; i++ {
		code += string(charset[rand.Intn(len(charset))]) // 随机从字符集中抽取字符并追加到字符串
	}

	fmt.Printf("用戶%d的驗證碼為%s", user.ID, code)
	const TokenExpireDuration = time.Hour * 24 * 30 * 12 * 30
	// 生成token
	token, _ := utils.GenerateToken(user.ID, code, TokenExpireDuration)
	// 發送郵箱
	if err := utils.SendEmail(user.Email, code); err != nil {
		return "", err
	}
	AddTask(user.ID)
	tokenCache[user.ID] = token

	return token, nil
}

func VerifyEmail(code string, token string) (uint, error) {
	userClaims, err := utils.ParseToken(token)
	if err != nil {
		return 0, errors.New("Token解析錯誤")
	}

	fmt.Println(userClaims)

	if userClaims.Code != code {
		return 0, errors.New("驗證碼錯誤")
	}

	CancelTask(userClaims.UserID)
	return userClaims.UserID, nil
}

// 添加一個異步延遲任務，分配一個控制碼
func AddTask(controlCode uint) {
	fmt.Printf("任務 %d 開始執行並等待五分鐘: %v\n", controlCode, time.Now())

	mu.Lock()
	tasks[controlCode] = make(chan struct{})
	mu.Unlock()

	go func(controlCode uint) {
		select {
		case <-time.After(5 * time.Minute): // 等待五分鐘
			fmt.Printf("%d已經刪除", controlCode)
			models.DeleteUser(controlCode) // 如果未被取消，則執行任務
		case <-tasks[controlCode]: // 接收到取消信號
			fmt.Printf("任務 %d 已被取消\n", controlCode)
		}

		// 從 map 中刪除任務
		mu.Lock()
		delete(tasks, controlCode)
		mu.Unlock()
	}(controlCode)
}

// 執行 B 方法來取消指定控制碼的任務
func CancelTask(controlCode uint) {
	mu.Lock()
	defer mu.Unlock()

	if task, exists := tasks[controlCode]; exists {
		close(task)                // 關閉取消通道以發送取消信號
		delete(tasks, controlCode) // 從 map 中刪除該任務
		fmt.Printf("取消任務 %d 成功\n", controlCode)
	} else {
		fmt.Printf("控制碼 %d 的任務不存在\n", controlCode)
	}
}

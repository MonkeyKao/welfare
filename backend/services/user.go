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

var tokenCache = []string{}

// 定義一個結構來保存任務的取消通道
type Task struct {
	cancelChan chan struct{}
}

// 用於存儲任務的 map 和一個互斥鎖
var tasks = make(map[uint]*Task)
var mu sync.Mutex

func Login(account string, password string) (string, error) {
	var user models.User
	const TokenExpireDuration = time.Hour * 24 * 30 * 12 * 30

	//檢查帳號是否存在
	if err := models.GetUserByAccount(&user, account); err != nil {
		return "", errors.New("account is not exist")
	}
	fmt.Print(user.Password)
	decryptText, err := utils.Decryption(user.Password, user.Salt)
	if err != nil {
		return "", err
	}

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

func Register(user *models.User) error {
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

	if err := models.CreateUser(user); err != nil {
		return err
	}
	const charset = "123456789" // 字符集
	code := ""                  // 创建一个空字符串用于存储验证码

	for i := 0; i < 6; i++ {
		code += string(charset[rand.Intn(len(charset))]) // 随机从字符集中抽取字符并追加到字符串
	}
	const TokenExpireDuration = time.Minute * 5

	token, _ := utils.GenerateToken(user.ID, code, TokenExpireDuration)
	if err := utils.SendEmail(user.Email, code); err != nil {
		return err
	}
	AddTask(user.ID)
	tokenCache = append(tokenCache, token)

	return nil
}

func VerifyEmail(code string, userId uint) (uint, error) {
	for _, token := range tokenCache {
		UserClaims, err := utils.ParseToken(token)
		if err != nil {
			removeElement(tokenCache, token)
			continue
		}

		if UserClaims.Code == code && UserClaims.UserID == userId {
			CancelTask(userId)
			return UserClaims.UserID, nil
		}
	}

	return 0, errors.New("例外處理")
}

// 添加一個異步延遲任務，分配一個控制碼
func AddTask(controlCode uint) {
	fmt.Printf("任務 %d 開始執行並等待五分鐘: %v\n", controlCode, time.Now())

	// 創建取消通道並添加到 map 中
	task := &Task{cancelChan: make(chan struct{})}
	mu.Lock()
	tasks[controlCode] = task
	mu.Unlock()

	go func(controlCode uint) {
		select {
		case <-time.After(5 * time.Minute): // 等待五分鐘
			models.DeleteUser(controlCode) // 如果未被取消，則執行任務
		case <-task.cancelChan: // 接收到取消信號
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
		close(task.cancelChan)     // 關閉取消通道以發送取消信號
		delete(tasks, controlCode) // 從 map 中刪除該任務
		fmt.Printf("取消任務 %d 成功\n", controlCode)
	} else {
		fmt.Printf("控制碼 %d 的任務不存在\n", controlCode)
	}
}

func removeElement(slice []string, element string) []string {
	// 创建一个新的切片来存储移除指定元素后的结果
	newSlice := []string{}

	// 遍历原切片，添加不等于指定元素的元素到新切片
	for _, item := range slice {
		if item != element {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice // 返回新的切片
}

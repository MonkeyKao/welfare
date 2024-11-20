package services

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"walfare/models"
	"walfare/utils"

	"github.com/skip2/go-qrcode"
)

var codeToFamilyID = make(map[string]uint)
var mutex sync.Mutex

func GenrateFmailyQRCode(familyId uint) (string, []byte) {

	var code string
	for {
		code = utils.GenerateCode()

		if _, exists := codeToFamilyID[code]; !exists {
			mutex.Lock()
			codeToFamilyID[code] = familyId
			mutex.Unlock()
			fmt.Printf("已經加入%s驗證碼", code)
			break
		}
	}

	png, _ := qrcode.Encode(code, qrcode.Medium, 256)

	go func(code string) {
		time.Sleep(time.Minute * 5)
		mutex.Lock()
		delete(codeToFamilyID, code)
		mutex.Unlock()
		fmt.Printf("已經移除%s驗證碼", code)
	}(code)

	return code, png
}

func JoinFmaily(UserId uint, code string) error {
	if FamilyId, exits := codeToFamilyID[code]; exits {
		familyMember := models.FamilyMember{
			UserID:   UserId,
			FamilyID: FamilyId,
			Role:     "測試",
		}

		if err := familyMember.CreateFamilyMember(); err != nil {
			return errors.New("加入失敗")
		}

	} else {
		return errors.New("code不存在")
	}

	return nil
}

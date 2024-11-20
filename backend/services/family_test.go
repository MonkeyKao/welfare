package services

import (
	"sync"
	"testing"
)

func TestGenrateFmailyQRCode(t *testing.T) {
	const numTasks = 100
	var wg sync.WaitGroup

	// 模擬 100 個任務，ID 從 1 到 100
	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go func(familyId uint) {
			defer wg.Done()

			got, got1 := GenrateFmailyQRCode(familyId)

			// 可以根據需要驗證生成的結果
			t.Logf("FamilyID: %d, Got: %s, Got1: %v", familyId, got, got1)
		}(uint(i))
	}

	// 等待所有併發任務完成
	wg.Wait()
}

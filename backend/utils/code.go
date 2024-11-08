package utils

import "math/rand"

func GenerateCode() string {
	// 生成驗證碼
	const charset = "123456789" // 字符集
	code := ""                  // 创建一个空字符串用于存储验证码

	for i := 0; i < 6; i++ {
		code += string(charset[rand.Intn(len(charset))]) // 随机从字符集中抽取字符并追加到字符串
	}

	return code
}

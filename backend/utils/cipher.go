package utils

import(
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//生成指定長度隨機字符串
func GenerateRandomStringWithCharset(length int)(string, error) {
	bytes := make([]byte, length) //建立一個length長度的的位元組切片

	for i := range bytes{
		randomByte,err := rand.Int(rand.Reader,big.NewInt(int64(len(charset))))
		if err != nil{
			return "",err
		}
		bytes[i] = charset[randomByte.Int64()]
	}

	return string(bytes),nil
}

//用AES進行加密
//plainText ->要加密的明文 (使用者打的密碼)
func Encryption(plainText string)(string,string, error){
	//生成隨機salt ( 密鑰 )
	salt, err :=GenerateRandomStringWithCharset(16)//丟進去上面函式裡隨機生成
	if err != nil{
		return "","",err
	}

	// 將 salt 作為密鑰
	key := []byte(salt)[:16] // AES-128需要16 字節密鑰

	block, err := aes.NewCipher(key) //生成AES加密區塊
	if err != nil {
		return "", "", err
	}
	
	//cipherText ->加密的結果
	cipherText := make([]byte,aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _,err := io.ReadFull(rand.Reader, iv); err != nil {
		return "","",err
	}

	//加密明文
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:],[]byte(plainText))

	//加密結果轉換為base64字串
	encryptedText := base64.StdEncoding.EncodeToString(cipherText)

	return encryptedText,salt, nil
}

// Decryption 使用 AES 進行解密
func Decryption(encryptedText string, salt string) (string, error) {
	// 將 salt 作為密鑰
	key := []byte(salt)[:16] // AES-128，16 字節密鑰

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

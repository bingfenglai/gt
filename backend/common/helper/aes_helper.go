package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	
	"io"
)

// aes cfb 加密
func AesEncryptCFB(origData []byte, key []byte) (string,error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        //panic(err)
		return "",err
    }
    encrypted := make([]byte, aes.BlockSize+len(origData))
    iv := encrypted[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        //panic(err)
		return "",err

    }
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	// 进行hex编码后再返回
    return hex.EncodeToString(encrypted),nil
}


// aes cfb 解密
func AesDecryptCFB(cypher string, key []byte) (string,error) {
	// 先hex解码
	encrypted,_:= hex.DecodeString(cypher)
    block, _ := aes.NewCipher(key)
    if len(encrypted) < aes.BlockSize {
        return "",errors.New("密文不合法")
    }
    iv := encrypted[:aes.BlockSize]
    encrypted = encrypted[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(encrypted, encrypted)
    return string(encrypted),nil
}
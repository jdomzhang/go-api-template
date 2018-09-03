package enc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
)

func VerifyMD5WithSecret(secret string, xmlData string) bool {
	obj, _ := Values{}.ParseXML(xmlData)

	sign := obj.Get("sign")
	obj.Del("sign")

	rawString := obj.String() + "&key=" + secret
	newSign := fmt.Sprintf("%x", md5.Sum([]byte(rawString)))

	return strings.ToUpper(sign) == strings.ToUpper(newSign)
}

func SignMD5WithSecret(secret string, obj Values) string {
	obj.Del("sign")

	rawString := obj.String() + "&key=" + secret
	newSign := fmt.Sprintf("%x", md5.Sum([]byte(rawString)))

	// to upper, important for wechat
	newSign = strings.ToUpper(newSign)

	obj.Add("sign", newSign)

	return newSign
}

func SignSHA1(obj Values) string {
	rawData := obj.Bytes()
	sign := fmt.Sprintf("%x", sha1.Sum(rawData))

	return sign
}

func DecryptWXData(key, iv, ciphertext string) (plaintext string, err error) {
	keyB, _ := base64.StdEncoding.DecodeString(key)
	ivB, _ := base64.StdEncoding.DecodeString(iv)
	encB, _ := base64.StdEncoding.DecodeString(ciphertext)

	result, err := decryptWXCBC(keyB, ivB, encB)

	if err != nil {
		return "", err
	}

	// strings.Trim(string(result), "\0x0e")

	return string(result), nil
}

func decryptWXCBC(key, iv, ciphertext []byte) (plaintext []byte, err error) {
	var block cipher.Block

	if block, err = aes.NewCipher(key); err != nil {
		return
	}

	if len(ciphertext) < aes.BlockSize {
		fmt.Printf("ciphertext too short")
		return
	}

	// iv = iv[:aes.BlockSize]
	// ciphertext = ciphertext[aes.BlockSize:]

	// fmt.Println("debugging:", block.BlockSize(), len(iv))

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(ciphertext, ciphertext)

	plaintext = ciphertext

	// trim PKCS#7填充
	// bs := []byte{0x0e}
	// if n := bytes.Index(plaintext, bs); n > 0 {
	// 	plaintext = plaintext[:n]
	// }
	plaintext = PKCS7UnPadding(plaintext)

	return
}

// PKCS7UnPadding will unpad the PKCS7 padded data
func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

package encrypting

import (
	str "strings"
)

type Encrypter interface {
	Encrypt(stringForEncrypt string) string
	Decrypt(stringForDecrypt string) string
}

type VowelLattersEncrypter struct {
	vowelLattersEncryptingMap map[string]string
}

func GetEncrypter() Encrypter {
	vowelLattersEncryptingMap := map[string]string{
		"a": "1",
		"e": "2",
		"i": "3",
		"o": "4",
		"u": "5",
	}
	return VowelLattersEncrypter{vowelLattersEncryptingMap}
}

func (encrypter VowelLattersEncrypter) Encrypt(stringForEncrypt string) string {
	var encryptedString string
	for _, latter := range str.ToLower(stringForEncrypt) {
		stringLatter := string(latter)
		if latterEncriptChar, ok := encrypter.vowelLattersEncryptingMap[stringLatter]; ok {
			stringLatter = latterEncriptChar
		}
		encryptedString += stringLatter
	}
	return encryptedString
}

func (encrypter VowelLattersEncrypter) Decrypt(stringForDecrypt string) string {
	var decryptedString string
	for _, latter := range str.ToLower(stringForDecrypt) {
		stringLatter := string(latter)
		for vowelLatter, latterEncriptChar := range encrypter.vowelLattersEncryptingMap {
			if stringLatter == latterEncriptChar {
				stringLatter = vowelLatter
			}
		}
		decryptedString += stringLatter
	}
	return decryptedString
}

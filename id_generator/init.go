package id_generator

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateId(s string) string {
	hash := sha256.New()
	//输入数据
	hash.Write([]byte(s))
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}

package authv4

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"time"
)

/*
AddAuthParams 添加鉴权相关参数 -
appKey : 应用ID
salt : 随机值
curtime : 当前时间戳(秒)
signType : 签名版本
sign : 请求签名
@param appKey    您的应用ID
@param appSecret 您的应用密钥
@param paramsMap 请求参数表
*/
func AddAuthParams(appKey string, appSecret string, params map[string][]string) {
	salt := getUuid()
	curtime := strconv.FormatInt(time.Now().Unix(), 10)
	sign := CalculateSign(appKey, appSecret, salt, curtime)
	params["appKey"] = []string{appKey}
	params["salt"] = []string{salt}
	params["curtime"] = []string{curtime}
	params["signType"] = []string{"v4"}
	params["sign"] = []string{sign}
}

/*
CalculateSign 计算v4鉴权签名 -
计算方式 : sign = sha256(appKey + input(q) + salt + curtime + appSecret)

@param appKey    您的应用ID
@param appSecret 您的应用密钥
@param q         请求内容
@param salt      随机值
@param curtime   当前时间戳(秒)
@return 鉴权签名sign
*/
func CalculateSign(appKey string, appSecret string, salt string, curtime string) string {
	strSrc := appKey + salt + curtime + appSecret
	return encrypt(strSrc)
}

func encrypt(strSrc string) string {
	bt := []byte(strSrc)
	bts := sha256.Sum256(bt)
	return hex.EncodeToString(bts[:])
}

func getUuid() string {
	b := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return ""
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

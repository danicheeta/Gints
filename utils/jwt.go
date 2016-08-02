package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func init() {
	fmt.Println("11111111")
	asd := NewJWTGenerator("dani")
	fmt.Println(asd.GenerateJWT("danicheeta@yahoo.com", "master"))
	fmt.Println()
	fmt.Println("22222222")
}

type JWTPayload struct {
	Admin  string
	Email  string
	Expire int64
}

type JWTGenerator struct {
	Secret []byte
}

func NewJWTGenerator(secret string) *JWTGenerator {
	return &JWTGenerator{Secret: []byte(secret)}
}

func (generator *JWTGenerator) EncodeBase64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func (generator *JWTGenerator) DecodeBase64(src string) []byte {
	dst, _ := base64.StdEncoding.DecodeString(src)
	return dst
}

func (generator *JWTGenerator) EncodeJWT(header string, payload string) []byte {
	return generator.EncodeHMAC(header + "." + payload)
}

func (generator *JWTGenerator) EncodeHMAC(src string) []byte {
	mac := hmac.New(sha256.New, generator.Secret)
	mac.Write([]byte(src))
	return mac.Sum(nil)
}

func (generator *JWTGenerator) ValidateJWT(src string) bool {
	jwt := strings.Split(src, ".")
	if len(jwt) != 3 {
		return false
	}
	return hmac.Equal(generator.DecodeBase64(jwt[2]),
		generator.EncodeJWT(jwt[0], jwt[1]))
}

func (generator *JWTGenerator) GenerateJWTGameAdmins(email string, game string) (jwt string) {
	jwt = generator.GenerateJWT(email, game)
	return
}

func (generator *JWTGenerator) GenerateJWTMasters(email string) (jwt string) {
	jwt = generator.GenerateJWT(email, "master")
	return
}

func (generator *JWTGenerator) GenerateJWTGameUsers(email string) (jwt string) {
	jwt = generator.GenerateJWT(email, "user")
	return
}

func (generator *JWTGenerator) GenerateJWT(email string, admin string) (jwt string) {
	header := make(map[string]string)
	payload := make(map[string]string)
	header["alg"] = "HS256"
	header["typ"] = "JWT"
	payload["mail"] = email
	payload["sub"] = "auth"
	now := time.Now().Unix()
	payload["exp"] = strconv.FormatInt(now+5*3600, 10)
	payload["admin"] = admin
	jsonHeader, _ := json.Marshal(header)
	jsonPayload, _ := json.Marshal(payload)
	bs64Header := generator.EncodeBase64(jsonHeader)
	bs64payload := generator.EncodeBase64(jsonPayload)
	bs64signature := generator.EncodeBase64(generator.EncodeJWT(bs64Header, bs64payload))
	jwt = bs64Header + "." + bs64payload + "." + bs64signature[:len(bs64signature)-1]
	return
}

func (generator *JWTGenerator) GenerateActivation(email string) (jwt string) {
	header := make(map[string]string)
	payload := make(map[string]string)
	header["alg"] = "HS256"
	header["typ"] = "JWT"
	payload["mail"] = email
	payload["sub"] = "activate"
	now := time.Now().Unix()
	payload["exp"] = strconv.FormatInt(now+24*3600, 10)
	jsonHeader, _ := json.Marshal(header)
	jsonPayload, _ := json.Marshal(payload)
	bs64Header := generator.EncodeBase64(jsonHeader)
	bs64payload := generator.EncodeBase64(jsonPayload)
	bs64signature := generator.EncodeBase64(generator.EncodeJWT(bs64Header, bs64payload))
	jwt = bs64Header + "." + bs64payload + "." + bs64signature
	return
}

func (generator *JWTGenerator) RenewJWT(jwt string) (new string) {
	payload := generator.Decode(jwt)
	new = generator.GenerateJWT(payload.Email, payload.Admin)
	return
}

func (generator *JWTGenerator) CheckExpire(exp int64) (expire bool) {
	now := time.Now().Unix()
	expire = now > exp
	return
}

func (generator *JWTGenerator) CheckReLogin(exp int64) (relogin bool) {
	now := time.Now().Unix()
	relogin = now > exp+3600*24*14
	return
}

func (generator *JWTGenerator) Decode(jwt string) (payload JWTPayload) {
	bs64Payload := generator.DecodeBase64(strings.Split(jwt, ".")[1])
	mapPayload := make(map[string]string)
	json.Unmarshal(bs64Payload, &mapPayload)
	payload.Email = mapPayload["mail"]
	payload.Admin = mapPayload["admin"]
	payload.Expire, _ = strconv.ParseInt(mapPayload["exp"], 10, 64)
	return
}

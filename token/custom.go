package token

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GenerateCustomToken generate a custom token
func GenerateCustomToken(str string, exp int) string {
	if exp <= 0 {
		return base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:forever", str)))
	}
	return base64.URLEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%d", str, time.Now().Add(time.Second*time.Duration(exp)).Unix())))
}

// ParseCustomToken validate custom token
func ParseCustomToken(token string) (string, error) {
	uDec, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	u := strings.Split(string(uDec), ":")
	expire, err := strconv.ParseInt(u[1], 10, 64)
	if err != nil {
		return "", err
	}
	if expire <= time.Now().Unix() {
		return "", errors.New("token is invalid")
	}
	return u[0], nil
}

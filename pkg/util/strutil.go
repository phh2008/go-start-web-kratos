package util

import (
	"github.com/google/uuid"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var snakeReg = regexp.MustCompile("[A-Z][a-z]")
var ColumnReg = regexp.MustCompile("^[A-Za-z0-9_]+$") //字母数字下划线
var DirectReg = regexp.MustCompile("^asc|desc|ASC|DESC$")

const underline = "_"

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
var innerRand = rand.New(rand.NewSource(time.Now().UnixMilli()))

// SnakeCase 驼峰转下划线
func SnakeCase(src string) string {
	str := snakeReg.ReplaceAllStringFunc(src, func(s string) string {
		return underline + s
	})
	return strings.ToLower(strings.TrimLeft(str, underline))
}

func UUID() string {
	u, _ := uuid.NewUUID()
	return strings.ReplaceAll(u.String(), "-", "")
}

func Random(length int) string {
	size := len(chars)
	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		idx := innerRand.Intn(size)
		result[i] = chars[idx]
	}
	return string(result)
}

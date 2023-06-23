package xjwt

import (
	"fmt"
	"github.com/cristalhq/jwt/v5"
	"helloword/internal/conf"
	"testing"
	"time"
)

// 生成 token
func TestCreateToken(t *testing.T) {
	var c = &conf.Bootstrap{}
	c.Jwt.Key = "xxx"
	jwtHelper := NewJwtHelper(c.Jwt)
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        "1000",
			Subject:   "tom",
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * 1)},
		},
	}
	token, err := jwtHelper.CreateToken(claims)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Token     %s\n", token.String())
	fmt.Printf("Algorithm %s\n", token.Header().Algorithm)
	fmt.Printf("Type      %s\n", token.Header().Type)
	fmt.Printf("Claims    %s\n", token.Claims())
	fmt.Printf("HeaderPart    %s\n", token.HeaderPart())
	fmt.Printf("ClaimsPart    %s\n", token.ClaimsPart())
	fmt.Printf("PayloadPart   %s\n", token.PayloadPart())
	fmt.Printf("SignaturePart %s\n", token.SignaturePart())
}

// 校验token
func TestVerifyToken(t *testing.T) {
	var c = &conf.Bootstrap{}
	c.Jwt.Key = "xxx"
	var tk = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxMDAwIiwic3ViIjoidG9tIiwiZXhwIjoxNjg0NTczMTMyLCJSb2xlIjoiIn0.a49j_EhGie4TgGmSFsH1Ffbx0JBf5G_y2dsk8zIiKqw"
	jwtHelper := NewJwtHelper(c.Jwt)
	token, err := jwtHelper.VerifyToken(tk)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Token     %s\n", token.String())
	fmt.Printf("Algorithm %s\n", token.Header().Algorithm)
	fmt.Printf("Type      %s\n", token.Header().Type)
	fmt.Printf("Claims    %s\n", token.Claims())
	fmt.Printf("HeaderPart    %s\n", token.HeaderPart())
	fmt.Printf("ClaimsPart    %s\n", token.ClaimsPart())
	fmt.Printf("PayloadPart   %s\n", token.PayloadPart())
	fmt.Printf("SignaturePart %s\n", token.SignaturePart())
	user, err := jwtHelper.ParseToken(token)
	fmt.Println(user)
	if !user.IsValidExpiresAt(time.Now()) {
		fmt.Println("expire .......")
	}
}

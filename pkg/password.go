package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 加密密码，生成 bcrypt 哈希
func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed)
}

// CheckPassword 校验明文密码和哈希是否匹配
func CheckPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

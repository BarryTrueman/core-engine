package coreengine

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	defaultDebugFormat = "%v | %s | %s"
	defaultDebugColor  = "\x1b[32m%s\x1b[0m"
)

func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(rune(rand.Intn(len(letters))))
	}
	return sb.String()
}

func CreateDirectoryIfNotExists(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func CheckIfFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func GetFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func GetCurrentDirectory() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return cwd
}

func GetDirectoryDepth(filePath string) int {
	depth := 0
	for _, part := range strings.Split(filepath.ToSlash(filePath), "/") {
		if part == "." || part == ".." {
			depth--
		} else {
			depth++
		}
	}
	return depth
}

func GetFileNameWithoutExtension(filePath string) string {
	return filepath.Base(filepath.ToSlash(filePath))
}

func GetFileExtension(filePath string) string {
	name := filepath.Base(filepath.ToSlash(filePath))
	ext := filepath.Ext(name)
	return ext
}

func GenerateUUID() string {
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		rand.Int63()<<32|rand.Int63()&0xffffffff,
		rand.Int63()<<16|rand.Int63()&0xffff,
		rand.Int63()<<8|rand.Int63()&0xff,
		rand.Int63()<<4|rand.Int63()&0xf,
		rand.Int63()&0x0fff,
	)
	return uuid
}

func ValidateEmail(email string) bool {
	if len(email) < 7 {
		return false
	}
	if strings.Count(email, "@") != 1 {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts[0]) == 0 {
		return false
	}
	if len(parts[1]) == 0 {
		return false
	}
	return true
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}
	if !strings.ContainsAny(password, "0123456789") {
		return false
	}
	if !strings.ContainsAny(password, "!@#$%^&*()_+-={}:<>?,./") {
		return false
	}
	return true
}

func IsDebugMode() bool {
	_, isSet := os.LookupEnv("DEBUG")
	return isSet
}

func GetDebugFormat() string {
	debugFormat := os.Getenv("DEBUG_FORMAT")
	if len(debugFormat) > 0 {
		return debugFormat
	}
	return defaultDebugFormat
}

func GetDebugColor() string {
	debugColor := os.Getenv("DEBUG_COLOR")
	if len(debugColor) > 0 {
		return debugColor
	}
	return defaultDebugColor
}

func IsStringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func GetRandomElement[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}

func IsNil[T any](v *T) bool {
	if v == nil {
		return true
	}
	return *v == nil
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfNil[T any](v *T) {
	if v == nil {
		panic(errors.New("nil value"))
	}
}
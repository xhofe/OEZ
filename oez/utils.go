package oez

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var CHARS = "InsV3Sf0obzp2i4gj1yYGqQv6wUtmBxlMAP7KHd8uTXFk9aRJWNC5EOhZDcLer"

const (
	// 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
	SCALE = 62
	REGEX = "^[0-9a-zA-Z]+$"
	NUM   = 6
)

func RandomStr(str string) string {
	chars := []rune(str)
	for i := len(chars) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		chars[i], chars[num] = chars[num], chars[i]
	}
	return string(chars)
}

func Encode10To62(val uint) string {
	if val < 0 {
		panic("val cannot be negative.")
	}
	str := ""
	var remainder int
	for math.Abs(float64(val)) > SCALE-1 {
		remainder = int(val % SCALE)
		str = string(CHARS[remainder]) + str
		val = val / SCALE
	}
	str = string(CHARS[val]) + str
	//for i := len(str); i < NUM; i++ {
	//	str = string(CHARS[0]) + str
	//}
	return str
}

func Decode62To10(val string) (uint, error) {
	if match, _ := regexp.MatchString(REGEX, val); !match {
		return 0, fmt.Errorf("illegal string: %s", val)
	}
	var result uint = 0
	index, length := 0, len(val)
	for i := 0; i < length; i++ {
		index = strings.Index(CHARS, string(val[i]))
		result += uint(index * int(math.Pow(float64(SCALE), float64(length-i-1))))
	}
	return result, nil
}

// RelativePath 获取相对可执行文件的路径
func RelativePath(name string) string {
	if filepath.IsAbs(name) {
		return name
	}
	e, _ := os.Executable()
	return filepath.Join(filepath.Dir(e), name)
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// CreatNestedFile 给定path创建文件，如果目录不存在就递归创建
func CreatNestedFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !Exists(basePath) {
		err := os.MkdirAll(basePath, 0700)
		if err != nil {
			log.Printf("无法创建目录，%s", err)
			return nil, err
		}
	}
	return os.Create(path)
}

// IsEmpty 返回给定目录是否为空目录
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}

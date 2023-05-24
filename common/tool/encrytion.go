package tool

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5ByString(str string) string {
	md5 := md5.New()
	_, err := io.WriteString(md5, str)
	if err != nil {
		panic(err)
	}
	sum := md5.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

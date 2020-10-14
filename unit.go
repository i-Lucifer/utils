package util

import (
	"fmt"
	"os"
)

// P 输出
func P(v ...interface{}) {
	fmt.Println(v...)
}

// DD 断点调试
func DD(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(0)
}

// Typeof 类型判断
func Typeof(v interface{}) {
	varType := fmt.Sprintf("%T", v)
	P(varType)
}

// FileSize 计算文件尺寸
func FileSize() {

}

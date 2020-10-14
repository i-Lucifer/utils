package hash

import (
	"testing"
)

// 单元测试
func TestMd5(t *testing.T) {
	str := "lucifer"
	md5 := Md5(str)
	// Luci := Md5("Lucifer")
	t.Log(md5)
	t.Log(Md5Check(str, md5))
	// fmt.Println(md5)
	// fmt.Println(Md5Check(str, md5))
	// fmt.Println(Md5Check(str, Luci))
}

// 以下两个基准测试性能是差不多的
// BenchmarkMd5 基准测试
func BenchmarkMd5(b *testing.B) {
	b.ResetTimer()
	str := "lucifer"
	for i := 0; i < b.N; i++ {
		_ = Md5(str)
	}
}

// BenchmarkByteMd5 基准测试
func BenchmarkByteMd5(b *testing.B) {
	b.ResetTimer()
	str := "lucifer"
	for i := 0; i < b.N; i++ {
		s := []byte(str)
		_ = ByteMd5(s)
	}
}

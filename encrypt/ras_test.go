package encrypt

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenRsaKey(t *testing.T) {
	bits := 2048
	path := "./"
	if err := GenRsaKey(bits, path); err != nil {
		p("秘钥生成失败")
	}
	p("秘钥生成成功")
}

func TestRsa(t *testing.T) {
	// 读取公钥私钥
	var err error
	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		os.Exit(-1)
	}
	privateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		os.Exit(-1)
	}

	// 开始加密
	str := "一篇诗，一壶酒"

	// 得到 []byte 数字
	Enc, err := RsaEncrypt([]byte(str))
	p(Enc)

	Dec, err := RsaDecrypt([]byte(Enc))
	p(Dec)

	// 得到字符串
	data, err := RsaEncryptString(str)
	p(data)

	oldStr, _ := RsaDecryptString(data)
	p(oldStr)
}

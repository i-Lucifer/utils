package accesstoken

const url = "https://api.weixin.qq.com/cgi-bin/token"

type param struct {
	appid     string
	secret    string
	grantType string
}

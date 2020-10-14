package subscribemsg

type params struct {
	templateID       string
	page             string
	miniprogramState string // trial developer formal
	lang             string // zh_CN
	touser           string // openid
	data             Data
}

// Data 订阅消息 模板 内容
type Data struct {
	thing5 string
	thing6 string
}

// Push 推送消息
func Push() {

}

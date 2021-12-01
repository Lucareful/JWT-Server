package types

// Authorization code 认证模式请求体.
//| 参数          | 类型   | 说明                                                         |
//| ------------- | ------ | -----------------------------------------------------------|
//| client_id     | string | 在oauth2 server 注册的client_id                             |
//| response_type | string | 固定值`code`                                                |
//| scope         | string | 权限范围,`str1,str2,str3`, 如果没有特殊说明,填`all`            |
//| state         | string | 验证请求的标志字段                                            |
//| redirect_uri  | string | 发放`code`用的回调uri,回调时会在uri后面跟上`?code=**&state=xxx` |
type Authorization struct {
	ClientID     string `json:"client_id"      form:"client_id"      binding:"required"`
	ResponseType string `json:"response_type"  form:"response_type"  binding:"required,eq=code"`
	Scope        string `json:"scope"          form:"scope"          binding:"required"`
	State        string `json:"state"          form:"state"          binding:"omitempty,oneof=1 2"`
	RedirectURI  string `json:"redirect_uri"   form:"redirect_uri"   binding:"required"`
}

// AccessToken 访问令牌请求体.
type AccessToken struct {
	AuthCode string `json:"auth_code"      form:"auth_code"      binding:"required"`
}

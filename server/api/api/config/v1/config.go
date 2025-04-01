package v1

import "github.com/gogf/gf/v2/frame/g"

type GetSmtpSuffixReq struct {
	g.Meta `path:"/config/getSmtpSuffix" method:"get" tags:"配置" summary:"获取SMTP后缀"`
}

type GetSmtpSuffixRes struct {
	List []string `json:"list"`
}

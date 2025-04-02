package apin

type EnglishGrammarTranslationReq struct {
	Text string `json:"text" v:"required#文本不能为空" dc:"文本"`
}

type EnglishGrammarTranslationRes struct {
	Result string `json:"result"`
}

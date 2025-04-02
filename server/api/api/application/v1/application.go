package v1

import (
	"hotgo/internal/model/input/apin"

	"github.com/gogf/gf/v2/frame/g"
)

// 英语语法翻译
type EnglishGrammarTranslationReq struct {
	g.Meta `path:"/application/englishGrammarTranslation" tags:"英语语法翻译" method:"post" summary:"英语语法翻译"`
	apin.EnglishGrammarTranslationReq
}

// 结果可能是文本或者sse
type EnglishGrammarTranslationRes struct {
	apin.EnglishGrammarTranslationRes
}

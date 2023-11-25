package fnapiio

type Lang string

const (
	LangNone   Lang = ""
	LangEN     Lang = "en"
	LangAR     Lang = "ar"
	LangDE     Lang = "de"
	LangES     Lang = "es"
	LangES419  Lang = "es-419"
	LangFR     Lang = "fr"
	LangIR     Lang = "it"
	LangJA     Lang = "ja"
	LangKO     Lang = "ko"
	LangPL     Lang = "pl"
	LangPtBR   Lang = "pt-BR"
	LangRU     Lang = "ru"
	LangTR     Lang = "tr"
	LangZhCN   Lang = "zh-CN"
	LangZhHant Lang = "zh-Hant"
)

// SetLanguage updates the client's default language.
//
// When no language is specified in the request's lang field option,
// this will be used for multilingual endpoints.
func (c *Client) SetLanguage(lang Lang) { c.lang = lang }

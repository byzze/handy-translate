package translate

import (
	"lyzee-translate/translate/caiyun"
	"lyzee-translate/translate/youdao"
)

type Transalte interface {
	PostQuery(value string) []string
}

func GetTransalteWay(name string) Transalte {
	switch name {
	case "youdao":
		return new(Youdao)
	case "caiyun":
		return new(Caiyun)
	default:
		return nil
	}
}

type Youdao struct{}

type Caiyun struct{}

func (y *Youdao) PostQuery(value string) []string {
	return youdao.PostQuery(value)
}

func (c *Caiyun) PostQuery(value string) []string {
	return caiyun.PostQuery(value)
}

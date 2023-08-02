package translate

import (
	"handy-translate/config"
	"handy-translate/translate/caiyun"
	"handy-translate/translate/youdao"
)

type Transalte interface {
	PostQuery(value string) []string
}

func GetTransalteWay(name string) Transalte {
	switch name {
	case youdao.Way:
		return &youdao.Youdao{
			Translate: config.Translate{
				Key:    config.Data.Translate[name].Key,
				Secret: config.Data.Translate[name].Secret,
			},
		}
	case caiyun.Way:
		return &caiyun.Caiyun{
			Translate: config.Translate{
				Token: config.Data.Translate[name].Token,
			},
		}
	default:
		return nil
	}
}

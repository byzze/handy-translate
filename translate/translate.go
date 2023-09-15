package translate

import (
	"handy-translate/config"
	"handy-translate/translate/baidu"
	"handy-translate/translate/caiyun"
)

type Transalte interface {
	PostQuery(value string) []string
}

func GetTransalteWay(name string) Transalte {
	switch name {
	case caiyun.Way:
		return &caiyun.Caiyun{
			Translate: config.Translate{
				Token: config.Data.Translate[name].Token,
			},
		}
	default:
		return &baidu.Baidu{
			Translate: config.Translate{
				Key: config.Data.Translate[name].Key,
			},
		}
	}
}

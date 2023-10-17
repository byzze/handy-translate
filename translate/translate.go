package translate

import (
	"handy-translate/config"
	"handy-translate/translate/baidu"
	"handy-translate/translate/caiyun"
	"handy-translate/translate/youdao"
)

type Translate interface {
	GetName() string
	PostQuery(value string) ([]string, error)
}

func GetTransalteWay(way string) Translate {
	var t Translate
	switch way {
	case youdao.Way:
		t = &youdao.Youdao{
			Translate: config.Translate{
				Name:  config.Data.Translate[way].Name,
				AppID: config.Data.Translate[way].AppID,
				Key:   config.Data.Translate[way].Key,
			},
		}
	case caiyun.Way:
		t = &caiyun.Caiyun{
			Translate: config.Translate{
				Name:  config.Data.Translate[way].Name,
				AppID: config.Data.Translate[way].AppID,
				Key:   config.Data.Translate[way].Key,
			},
		}
	case baidu.Way:
		t = &baidu.Baidu{
			Translate: config.Translate{
				Name:  config.Data.Translate[way].Name,
				AppID: config.Data.Translate[way].AppID,
				Key:   config.Data.Translate[way].Key,
			},
		}
	}

	return t
}

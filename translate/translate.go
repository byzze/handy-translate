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

func GetTransalteWay() []Translate {
	var trList []Translate
	for _, v := range config.Data.Translate {
		var t Translate
		switch v.Name {
		case youdao.Way:
			t = &youdao.Youdao{
				Translate: config.Translate{
					Name:   v.Name,
					Secret: v.Secret,
					AppID:  v.AppID,
				},
			}
		case caiyun.Way:
			t = &caiyun.Caiyun{
				Translate: config.Translate{
					Name:   v.Name,
					Key:    v.Key,
					Secret: v.Secret,
					Token:  v.Token,
				},
			}
		case baidu.Way:
			t = &baidu.Baidu{
				Translate: config.Translate{
					Name:   v.Name,
					Key:    v.Key,
					Secret: v.Secret,
					Token:  v.Token,
					AppID:  v.AppID,
				},
			}
		}
		trList = append(trList, t)
	}
	return trList
}

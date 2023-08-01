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
			Key:    config.Data.Translate[name].Key,
			Secret: config.Data.Translate[name].Secret,
		}
	case caiyun.Way:
		return &caiyun.Caiyun{
			Key:    config.Data.Translate[name].Key,
			Secret: config.Data.Translate[name].Secret,
		}
	default:
		return nil
	}
}

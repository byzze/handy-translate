package main

import (
	"encoding/json"
	"handy-translate/config"
	"handy-translate/translate_service"
	"handy-translate/utils"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v3/pkg/application"
)

var fromLang, toLang = "auto", "zh"

type Person struct {
	name string
}

// App is a service that greets people
type App struct {
}

// Greet greets a person
func (*App) Greet(name string) string {
	return "Hello " + name
}

// GreetPerson greets a person
func (*App) GreetPerson(person Person) string {
	return "Hello " + person.name
}

// MyFetch URl
func (a *App) MyFetch(URL string, content map[string]interface{}) interface{} {
	return utils.MyFetch(URL, content)
}

func (a *App) Transalte(queryText, fromLang, toLang string) {
	logrus.WithFields(logrus.Fields{
		"queryText": queryText,
		"fromLang":  fromLang,
		"toLang":    toLang,
	}).Info("Transalte")

	// 翻译loading
	app.Events.Emit(&application.WailsEvent{Name: "loading", Data: "true"})
	defer app.Events.Emit(&application.WailsEvent{Name: "loading", Data: "false"})

	transalteWay := translate_service.GetTransalteWay(config.Data.TranslateWay)

	// curName := transalteWay.GetName()
	// 使用 strings.Replace 替换 \r 和 \n 为空格

	result, err := transalteWay.PostQuery(queryText, fromLang, toLang)
	if err != nil {
		logrus.WithError(err).Error("PostQuery")
	}

	logrus.WithFields(logrus.Fields{
		"result":       result,
		"transalteWay": transalteWay.GetName(),
	}).Info("Transalte")

	transalteRes := strings.Join(result, "\n")
	sendDataToJS(queryText, transalteRes, "")
}

func sendDataToJS(query, result, explian string) {
	sendQueryText(query)
	sendResult(result, explian)
}

func sendQueryText(queryText string) {
	app.Events.Emit(&application.WailsEvent{Name: "query", Data: queryText})
}

func sendResult(result, explian string) {
	app.Events.Emit(&application.WailsEvent{Name: "result", Data: result})
	app.Events.Emit(&application.WailsEvent{Name: "explian", Data: explian})
}

func (a *App) GetTransalteMap() string {
	var translateList = config.Data.Translate
	bTranslate, err := json.Marshal(translateList)
	if err != nil {
		logrus.WithError(err).Error("Marshal")
	}
	return string(bTranslate)
}

func (a *App) SetTransalteWay(translateWay string) {
	config.Data.TranslateWay = translateWay
	translate_service.SetQueryText("")
	config.Save()
	logrus.WithField("Translate", config.Data.Translate).Info("SetTransalteList")
}

func (a *App) GetTransalteWay() string {
	return config.Data.TranslateWay
}

// Show 通过名字控制窗口事件
func (a *App) Show(windowName string) {
	if _, ok := windowMap[windowName]; ok {
		windowMap[windowName].Center()
		windowMap[windowName].Show()
	}
}

// Hide 通过名字控制窗口事件
func (a *App) Hide(windowName string) {
	if _, ok := windowMap[windowName]; ok {
		windowMap[windowName].Hide()
	}
}

# 概述
基于wails框架，结合Go+React，开发支持多平台(Windows, Linux, Mac)的翻译工具。当鼠标选中文本时，点击鼠标中键，弹出窗口渲染翻译结果。目前Windows平台效果较好，其他平台理论上也可以编译运行，但体验没Windows好。通过wails构建生成的包容量相较于Electron小，仅有10M左右。 本次版本开发使用的是wails[v3](https://v3alpha.wails.io/)版本，该版本处于alpha版本，不太稳定，若需要稳定开发可以参考[v2](https://wails.io/)版本

# 功能说明
- [X] 鼠标选中文字进行翻译
- [X] 通过配置文件自定义快捷键
- [X] 支持有道，百度，彩云翻译源
- [X] 支持截图OCR翻译
- [X] 系统托盘
  
# 效果展示
按压**鼠标中键**弹出窗口
按压**CTRL+SHIFT+F**弹出窗口

![示例视频](https://raw.githubusercontent.com/byzze/oss/main/handly-translate/effect.gif)

# 安装编译环境
安装wails(重要), 此软件基于v3版本开发，但v3处于alpha测试版本
**v3**
`https://v3alpha.wails.io/`

**v2**
`https://wails.io/docs/gettingstarted/installation/`


# 编译构建

## 方式一
直接编译可执行文件
`go build -tags production -ldflags="-w -s -H windowsgui" -o handy-translate.exe` 

## 方式二
下载对应的wails版本进行构建，并替换`go.mod`文件内容`replace github.com/wailsapp/wails/v3 => D:\go_project\wails\v3`(注意该路径对应的系统路径)

# 配置翻译源
填写对应的翻译秘钥
**修改配置名**
`config.toml.bak -> config.toml`

**填写对应的api信息**
```toml
appname = "handy-translate"
translate_way = "baidu"

[keyboards] 
toolBar = ["center", "", ""] # 小窗口翻译快捷键， 表示鼠标中键
screenshot = ["ctrl", "shift",  "f"] # 截图快捷键

[translate]
[translate.baidu] # https://fanyi-api.baidu.com/api/trans/product/apidoc
name = "百度翻译"
appID = "20230823001790949"
key = "hTlcbpu7xxxxxxxxx"

[translate.youdao] # https://ai.youdao.com/DOCSIRMA/html/trans/api/wbfy/index.html
name = "有道翻译"
appID = "appKey"
key = "appSecret"
```
# 执行
双击生成文件`./handry-translate.exe`

# OCR models
实现截图ocr解析文件模型，该模型有点大，大约75M, 文件夹：models

# 参考用到的工具组件链接
- [robotgo](https://github.com/go-vgo/robotgo) 鼠标，键盘监听
- [wails v2](https://wails.io)
- [wails v3](https://v3alpha.wails.io/)
- [NEXTUI](https://nextui.org/) 前端UI组件
- [pot-desktop](https://github.com/pot-app/pot-desktop) rust开发的跨平台翻译工具
- [go-qoq](https://github.com/duolabmeng6/go-qoq) wails3开发的翻译工具
# 概述
一款便捷翻译的工具，使用wails框架开发，wails支持Go+Vue, React等多种前端框架结合使用，同时也支持多平台Windows, Linux, Mac开发，该工具支持自定义快捷键，当鼠标选中的文本时，按压鼠标中键，弹出窗口渲染翻译结果，由于目前没有多余的开发设备，所以仅验证了Windows，wails生成的包容量相较于Electron的是相当的小了, 仅有10M左右

# 功能说明
- [X] 选择文字翻译为中文
- [X] 自定义快捷键
- [X] 多种翻译源
  
# 效果展示
按压鼠标中键
![示例视频](https://raw.githubusercontent.com/byzze/oss/main/handly-translate/exp.gif)

# 安装编译环境
安装wails(重要)
`https://wails.io/docs/gettingstarted/installation/`

## 检查wails所需环境是否安装成功
`wails doctor`

# OCR models
实现截图ocr解析文件模型，该模型有点大，大约75M

# 配置翻译源
填写对应的翻译秘钥即可
百度翻译源: https://docs.caiyunapp.com/blog/2021/12/30/hello-world
有道翻译：https://ai.youdao.com/DOCSIRMA/html/trans/api/wbfy/index.html
彩云翻译：https://docs.caiyunapp.com/blog/2021/12/30/hello-world

**修改配置名**
`config.toml.bak -> config.toml`

**填写对应的api信息**
```toml
appname = 'handy-translate'
keyboard = ['center', '', ''] # 快捷键配置,默认鼠标中键 可指定固定顺序["ctrl","shift","c"] 通过配置文件或界面操作配置快捷键
translate_way = 'baidu'

[translate]
[translate.baidu]
name = '百度翻译'
appID = ''
key = ''

[translate.youdao]
name = '有道翻译'
appID = ''
key = ''
```

# 编译构建

**方式一**
编译可执行文件
`wails build` 
复制`config.toml->./build/bin/`

双击windows生成文件`./build/bin/handry-translate.exe`

**方式二**
编译安装包，**建议**使用该方式，可以打包配置文件
`wails build -nsis`
双击生成文件`./build/bin/handy-translate-amd64-installer.exe`安装, 并执行安装成功后的`handry-translate.exe`文件

# 参考用到的工具组件链接
https://wails.io (使用了Go+VUE)

https://www.naiveui.com/zh-CN/os-theme/docs/installation (前端组件)


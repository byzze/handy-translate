# 概述
一款便捷翻译的工具，支持多平台Windows, Linux, Mac，支持自定义快捷键，鼠标选中的文本，按压鼠标中键，弹出窗口渲染翻译结果，但目前开发没有多余的设备，所以仅验证适配了Windows

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


# 配置翻译源
填写对应的翻译秘钥即可
百度翻译源: https://docs.caiyunapp.com/blog/2021/12/30/hello-world
有道翻译：https://ai.youdao.com/DOCSIRMA/html/trans/api/wbfy/index.html
彩云翻译：https://docs.caiyunapp.com/blog/2021/12/30/hello-world

**修改配置名**
`config.toml.bak -> config.toml`

**填写对应的api信息**
```toml
keyboard = ["","",""] # 快捷键配置,默认鼠标中键 可指定固定顺序["ctrl","shift","c"] 通过配置文件或界面操作配置快捷键
[[translate]]
name = '百度翻译'
appID = ''
key = ''

[[translate]]
name = '有道翻译'
appID = ''
key = ''

[[translate]]
name = '彩云翻译'
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

# 参考用到的工具组件
https://wails.io (后端go组件)

https://www.naiveui.com/zh-CN/os-theme/docs/installation (前端组件)


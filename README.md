# 概述
一款便捷翻译的工具，支持多平台Windows, Linux, Mac，支持自定义快捷键，鼠标选中的文本，按压鼠标中键，弹出窗口渲染翻译结果

# 功能说明
- [x] 选择文字翻译为中文
- [x] 自定义快捷键
- [ ] 多种翻译源(已完成百度，彩云)

# 安装
安装wails(重要)
`https://wails.io/docs/gettingstarted/installation/`

修改配置名
`config.toml.bak -> config.toml`

配置百度，彩云翻译秘钥(后续会支持多种翻译)
```txt
    [translate.default]
    name = "xx翻译"
    key = "xxxx"
    secret = ""
    token = ""
```
# 编译
`wails build`
# 运行
双击windows生成文件`./build/bin/handry-translate.exe`
# 使用效果
按压鼠标中键，弹出窗口，支持自定义快捷键
![Alt text](image.png)
# 参考用到的工具组件
https://wails.io (后端go组件)

https://www.naiveui.com/zh-CN/os-theme/docs/installation (前端组件)


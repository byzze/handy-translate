# 概述
一款便捷翻译的工具，支持多平台Windows, Linux, Mac，支持自定义快捷键，鼠标选中的文本，按压鼠标中键，弹出窗口渲染翻译结果

# 功能说明
- [x] 选择文字翻译为中文
- [x] 自定义快捷键
- [ ] 多种翻译源

# 支持的翻译源
填写对应的翻译秘钥即可
百度翻译源: https://docs.caiyunapp.com/blog/2021/12/30/hello-world
有道翻译：https://ai.youdao.com/DOCSIRMA/html/trans/api/wbfy/index.html
彩云翻译：https://docs.caiyunapp.com/blog/2021/12/30/hello-world

```toml
keyboard = ["","",""] # 快捷键配置,默认鼠标中键 可指定固定顺序["ctrl","shift","c"] 通过配置文件或界面操作配置快捷键
translate = [
    {name = "百度翻译", key = "", secret = "", token = ""},
    {name = "有道翻译", key = "", secret = "", token = ""},
    {name = "彩云翻译", key = "", secret = "", token = ""},
]
```
修改配置名
`config.toml.bak -> config.toml`

# 安装
安装wails(重要)
`https://wails.io/docs/gettingstarted/installation/`
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


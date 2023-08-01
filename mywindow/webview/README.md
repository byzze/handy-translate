# 记录使用webview

## 项目地址
https://github.com/webview/webview

## 处理环境问题
https://github.com/webview/webview/issues/776

### 脚本运行，运行demo
```bash
mkdir my-project && cd my-project
mkdir build libs "libs/webview"
curl -sSLo "libs/webview/webview.h" "https://raw.githubusercontent.com/webview/webview/master/webview.h"
curl -sSLo "libs/webview/webview.cc" "https://raw.githubusercontent.com/webview/webview/master/webview.cc"

mkdir libs\webview2
curl -sSL "https://www.nuget.org/api/v2/package/Microsoft.Web.WebView2" | tar -xf - -C libs\webview2
copy /Y libs\webview2\build\native\x64\WebView2Loader.dll build

go mod init example.com/m
go get github.com/webview/webview

set CGO_CPPFLAGS="-I%cd%\libs\webview2\build\native\include"
set CGO_LDFLAGS="-L%cd%\libs\webview2\build\native\x64"

curl -sSLo basic.go "https://raw.githubusercontent.com/webview/webview/master/examples/basic.go"

go build -ldflags="-H windowsgui" -o build/basic.exe basic.go && "build/basic.exe"
```

#### 执行脚本遇到的问题
1. 下载链接不成功
C:\Users\loyd\Desktop\byzze\my-project>curl -sSLo "libs/webview/webview.cc" "https://raw.githubusercontent.com/webview/webview/master/webview.cc"
curl: (56) Recv failure: Connection was reset

C:\Users\loyd\Desktop\byzze\my-project>curl -sSLo basic.go "https://raw.githubusercontent.com/webview/webview/master/examples/basic.go"
curl: (56) Recv failure: Connection was reset

C:\Users\loyd\Desktop\byzze\my-project>curl -sSLo basic.go "https://raw.githubusercontent.com/webview/webview/master/examples/basic.go"
curl: (28) Failed to connect to raw.githubusercontent.com port 443 after 87193 ms: Couldn't connect to server

**重试几次**就能下载，或者手动访问网址拷贝粘贴
2. 下载失败
curl -sSL "https://www.nuget.org/api/v2/package/Microsoft.Web.WebView2" | tar -xf - -C libs\webview2
**通过切换手机热点成功下载**

## 安装gcc

## 在powershell执行报错
PS C:\Users\loyd\Desktop\byzze\handy-translate> go run main.go
# github.com/webview/webview
In file included from webview.cc:1:
webview.h:1105:10: fatal error: WebView2.h: No such file or directory
 1105 | #include "WebView2.h"
      |          ^~~~~~~~~~~~
compilation terminated.
**切换原生cmd执行成功**
# Sync Transfor

局域网文件传输

## 需要修改远程仓库代码

- `go`中，默认是不允许修改第三方库的代码的，可通过下面的方法来修改

	- 克隆远程仓库到本地

	- 添加`replace 要修改的仓库地址(如：github.com/zserge/lorca) => 克隆下来的本地仓库路径`到`go mod`内，该语句和`require`平级
	- 执行`go mod tidy`

修改后`go.mod`如下：

```Go
module lorcaRelated

go 1.20

replace github.com/zserge/lorca => 克隆下来的lorca项目绝对路径

require (
	github.com/zserge/lorca v0.1.10 // indirect
	golang.org/x/net v0.9.0 // indirect
)
```



## lorca

- `lorca`打开`chrome`窗口后，窗口默认上面会有一个提示条`Chrome is being controlled by automated test softwares`，要关闭这个提示，需要打开`lorca/ui.go`然后把`--enable-automation`这个选项注释掉就行

### 报错修复

`新版Chrome`更新后，运行项目，报错如下

```shell
websocket.Dial ws://127.0.0.1:51684/devtools/browser/1d178186-4808-46d1-affd-5201a4b1d59f: bad status
```

解决方法:

调用`lorca.New`时添加参数`--remote-allow-origins=*`，解决方法参考[这里](https://github.com/zserge/lorca/issues/184)
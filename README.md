# study_go

# Go安装



下载 ：从 [Go 官网](https://golang.google.cn/)下载对应系统的Go

安装,选择不太复杂的路劲保存Go

安装完成后通过命令行执行

```powershell
			go version
```

验证查看Go版本信息

配置GOPATH环境变量:

*GOPATH* 是一个环境变量，指向本机存放Go项目的目录（工作目录）

*GOPATH* 路劲最好只设置一个，所有的项目代码都放到*GOPATH*的src目录下

Linux环境编辑

```powershell
sudo vi ~./bash_profile

​				export GOROOT='your go home'

​				export GOPATH='your dir'

​				export PATH=$PATH$:$GOROOT$/bin
```

保存后source ~./bash_profile使其生效

在GOPATH目录新建*src , pkg , bin*目录分别用来存放*源代码，包，执行文件*

执行

```powershell
go env
```

将会看到如下系统环境信息

![image-20191113104837269](/Users/k66/Desktop/go/src/github.com/wjse/study_go/README.assets/image-20191113104837269.png)


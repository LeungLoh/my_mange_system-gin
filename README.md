# my_mange_system-gin

<img src="https://img.shields.io/badge/gin-1.7.2-brightgreen.svg" alt="gin">
<img src="https://img.shields.io/badge/gorm-1.21.10-brightgreen.svg" alt="gorm">
<img src="https://img.shields.io/badge/mysql-1.0.0-brightgreen.svg" alt="mysql">


## 前言

配合[前端](https://github.com/LeungLoh/my_mange_system)使用gin+gorm编写的一套简单的后台管理程序
## 功能

-   [x] 登录/注销
-   [x] Dashboard获取内存信息
-   [x] 表格增删改查
-   [x] jwtauth 验证
-   [x] cors 跨域
-   [x] response 封装统一返回格式

##  安装步骤

~~~shell
git clone https://github.com/LeungLoh/my_mange_system-gin.git      // 把模板下载到本地
cd my-mange-system-gin                                             // 进入模板目录
go mod tidy                                                        // 安装项目依赖
go run main.go                                                     // 开启服务器，浏览器访问 http://localhost:3000
go build -0 main main.go                                           // 执行构建命令
~~~

## 部署
* 安装nginx 配置反向代理
* 开放对外的防火墙端口
* 将前端项目编译好的dist文件cpoy,并配置nginx
* go 编译好的项目使用nohup 或者 docker部署均可

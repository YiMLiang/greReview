# gre-review
给小艺鸭写的背单词软件
<br>单词写在excel中，暂不支持自动导入，手动输入你想背的单词，格式

|  En   | Zh  |
|  ----  | ----  |
| hehe  | 呵呵 |
| haha  | 哈哈 |
| heihei | 嘿嘿 |


## common
`filePath.go` 中 FileDir修改文件存放的位置

## model
两个结构体Word 和 File，后续可扩展

## util
一些常用工具类，后续可扩展

## 用法
1. mac 系统直接拷贝src下gre二进制文件到随机目录下（除了单词本文件存放位置）
2. 如果是linux或windows系统执行 `go env` 查看 GOOS属性，
设置成相应的即可
```
linux环境执行：
CGO_ENABLED=0 GOOS=linux GOARCH=amd64
windows环境执行：
CGO_ENABLED=0 GOOS=windows GOARCH=amd64
```
3. `go build /你gre.go的目录/gre.go 编译.go文件`，就会生成当前目录名的可执行文件并放置于当前目录下
<br>

|附加参数|备  注|
|----| ----|
|-o|指定编译文件的名字|
|-p n|开启并发编译，默认情况下该值为 CPU 逻辑核数|
|-a|强制重新构建|
|-n|打印编译时会用到的所有命令，但不真正执行|
|-x|打印编译时会用到的所有命令|
|-race|开启竞态检测|
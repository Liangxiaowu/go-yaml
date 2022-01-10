# go-yaml
Yaml parsing Toolkit

#### 介绍
[gopkg.in/yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3#section-readme) 已经是个非常好用的包，但是在实际开发中总有类型转换带来的麻烦，`go-yaml`只是在它的基础上，简单的一层封装，作为某些类型的定义转换。

#### go module
```
go get github.com/Liangxiaowu/go-yaml
```


#### yaml文件
默认情况下会读取当前项目目录下的`./configs/app.yaml`文件

#### 初始化
```go
## 默认读取./configs/app.yaml文件
yaml := New()

## 自定义yaml文件路径
yaml := New(FilePath("./conf/app.yaml"))

## 自定义yaml文件,读取./configs/xxx.yaml文件
yaml := New(Name("xxx.yaml"))

## 自定义yaml文件地址,读取./conf/app.yaml文件
yaml := New(Dir("./conf"))
```
#### 结构体列子
app.yaml:
```yaml
user:
  name: wunder
  age: 18

```
main.go
```go
type User struct {
    Name string
    Age  int `json:"age"`
}

func getUser()  {
    var u User           # 映射结构体
    err := New().Get(&u) # 按照顶级查询出一个结构体
    fmt.Println(err)
    fmt.Println(u)
}
```

#### 获取指定参数值
```go
....
```
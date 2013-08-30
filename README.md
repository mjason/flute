flute
=====

一个基于mux的restful框架, 选择最好的第三方包进行整合,尽量不重复制造轮子

## 项目依赖的第三方包
- github.com/gorilla/mux
- github.com/bitly/go-simplejson
- github.com/robfig/config

## 例子说明
```go
package main

import (
  "flute"
)

type Users struct {
  flute.Controller
}

func (u *Users) Index() {
  u.RenderText(200, "mj")
}

func (u *Users) Show() {
  u.RenderText(200, "mj"+u.GetVars("id"))
}

func main() {
  flute.HandleFunc("/", "GET", func(c *flute.Context) {
    c.RenderText(200, "hello word")
  })
  var users Users
  flute.Resources("users", &users)
  flute.Start(":1234")
}
```

十分像rails的一些东西, 内部实现很多闭包,支持restful路由, 基于mux的小封装,更大的自由

## 接下来的东西
- 封装好mgo驱动,提供更加方便的方法
- 增加插件功能,比如before和after,以便做框架验证
- 增加文档

## 适用范围
- 不适合与web开发, 这个框架没有想过集成什么template或者我去实现
- 适合写web api, 所有功能都将围绕web api来开发

我怎么会想到笛子这个单词了~~

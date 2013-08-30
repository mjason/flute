// flute - 一个简单的golang web api 开发框架
//
// Overview
// 开发这套框架主要是因为在开发api中,我找不到一个比较好的框架,ruby的框架主要还是太多东西了,我希望有一个简单的框架,同时性能更好一些的框架, 所以开发这套框架
// 这套框架的路由的想法来自rails的方法, 目前实现了Resources方法. 这套框架
// 这套框架可以干什么呢? 仅仅能开发web api而已,没有其他的功能.
// 目前Render方法还没做好, 如果你有耐心请等一等, 如果没有可以用http你们的包,代码不多可以直接看代码, 封装都很薄的.
//
// 下面建立一个简单的例子, 一个简单的hello word.
//    flute.HandleFunc("/", "GET", func(c *flute.Context) {
//      c.RenderText(200, "hello word")
//    })
//    flute.Start(":1234")
// github的wiki后续会加入更多的例子
package flute

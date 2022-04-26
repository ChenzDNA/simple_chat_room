## 几把聊天

实现了基本的实时通信。

使用 `Transfer-Encoding: chunked` 实现服务端到客户端的长连接，客户端每次进行发送消息的请求都会由这个长连接推送到所有客户端。

### backend:
- [GoFrame](https://github.com/gogf/gf)

### frontend:
- [Vue3](https://github.com/vuejs/core)
- [TypeScript](https://github.com/microsoft/TypeScript)
- [vite](https://github.com/vitejs/vite)
- [naive-ui](https://github.com/TuSimple/naive-ui)
- [axios](https://github.com/axios/axios)

前端有些依赖没有用上（ [pinia](https://github.com/vuejs/pinia) 和 [vue-router4](https://github.com/vuejs/router) ），先放着不管了，反正都配好了，以后哪天心血来潮加一堆功能也说不定。

[踩坑记录](https://chenzdna.github.io/post/goframe-shi-yong-chunked-shi-xian-ji-shi-tong-xin/)

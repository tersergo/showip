# showip: 一个查看客户端ip的web服务

## 服务启动参数

> $ showip -port=80 -path=/showip

- port：服务响应端口(默认端口80)
- path：服务响应的路径(默认访问路径/showip)

## 服务访问参数(api)

> http://{your-web-services}{:port}/{path}?format={text|json|array|xml|html}

- format: web输出格式参数，可选参数text|json|array|xml|html（默认参数text）

1. format=text 或者 format参数为空时，返回结果为纯文本
> 请求 http://localhost/showip 或者 http://localhost/showip?format=text
```text
IP: 127.0.0.1
```
2. format=array 返回结果为ip数组
> http://192.168.2.1/showip?format=array
```js
["192.168.2.13"]
```
3. format=json 返回结果为json格式
```json
{
   "IP": "192.168.2.13",
   "RemoteAddress": "192.168.2.13"
}
```
4. format=xml  返回结果为xml格式
```xml
<? version="1.0" encoding="UTF-8" ?>
<showip>
    <RemoteAddress>192.168.2.13</RemoteAddress>
    <IP>192.168.2.13</IP>
</showip>
```
5. format=html 返回结果为html无序列表格式
```html
<ul class="showip">
    <li>192.168.2.13</li>
</ul>
```


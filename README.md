# showip: 一个通用获取客户端IP和服务器IP的web服务

## 功能特点

- 支持客户端IP展示和api获取
- 支持服务器IP展示和api获取
- 支持ip的V4和v6获取
- 支持获取代理和转发协议的用户的真实IP
- 支持获取自定义协议或扩展协议头中IP

## 服务启动参数(shell)

> $ showip -port=80 -path=/showip 

- port：服务响应端口(默认端口80)
- path：服务响应的路径(默认路径/showip)
- header：设置客户端IP相关的header头参数（优先获取指定header名称多个参数,分隔）

### 请求参数名称重命名和关闭参数响应

- format：请求输出格式参数名称重命名和关闭设置（默认名称format），format=0时关闭格式参数响应操作
- mode：服务器模式参数重命名和关闭设置，（默认名称mode），mode=0时关闭服务器IP相关的响应操作
- via：响应头X-Via名称重命名和关闭设置（默认名称X-Via），via=0时会关闭X-Via服务器IP头输出
```text
-- Response Header 默认输出下面内容（该信息可通过启动参数via来重命名或关闭输出）
  X-Via: 192.168.2.1
```

## 服务访问参数(web|api)

> http://{service-name}[:port]/{path}?[format={text|json|array|xml|html}]&[obj={showip}]&[mode=host]

- format: 可选参数，输出格式，默认参数text可选参数text|json|array|xml|html
- obj: 可选参数，html和xml输出对象名称，html时为ul对象id名称，xml时为根节点名称
- mode: 可选参数，默认不传返回客户端IP信息，当mode=host时返回服务器IP信息

1. format参数为空或format=text时，返回客户端第一个有效IP的纯文本格式，*当参数包含mode=host时返回结果为服务器第一个有效IP的纯文本格式*
> 请求 http://localhost/showip 或者 http://localhost/showip?format=text
```text
IP: 127.0.0.1
-- ip如果是V6时，返回以下结果
IP: ::1
```
2. format=array 返回结果为获取到的所有IP的数组形式
> http://192.168.2.1/showip?format=array
```js
["192.168.2.13"]
```
3. format=json 返回结果为客户端IP的json格式
```json
{
   "IP": "192.168.2.13",
   "RemoteAddress": "192.168.2.13"
  /* 如果请求包含以下头部协议，返回可能会包含下列内容
  "X-Forwarded-For": "xx.xx.xx.xx"
  "X-Real-IP": "xx.xx.xx.xx"
  "Proxy-Client-IP": "xx.xx.xx.xx"
  "WL-Proxy-Client-IP": "xx.xx.xx.xx"
  */
}
/* format=json&mode=host 返回结果为服务器IP的json格式 
{
  "IP": "192.168.2.1",
  "IPV4": "192.168.2.1"
  // 如果是ipV6时会有 "IPV6": "ff::aa:bb:cc:dd"
}
*/
```
4. format=xml  返回结果为客户端IP的xml格式
> xml的根节点名称可以通过obj参数设置
```xml
<? version="1.0" encoding="UTF-8" ?>
<showip>
    <IP>192.168.2.13</IP>
    <RemoteAddress>192.168.2.13</RemoteAddress>
</showip>
<!-- format=xml&mode=host  返回结果为服务器IP的xml格式
<? version="1.0" encoding="UTF-8" ?>
<showip>
    <IP>192.168.2.1</IP>
    <IPV4>192.168.2.1</IPV4>
    // 如果是ipV6时会有 <IPV6>ff::aa:bb:cc:dd</IPV6>
</showip>
-->
```
5. format=html 返回结果为html无序列表格式
> html中ul的id名称可以通过请求obj参数来设置，样式固定为showip不会随指定obj名称改变
```html
<ul id="showip" class="showip">
    <li>192.168.2.13</li>
</ul>
```

## 客户端IP获取的顺序

默认情况下会依次获取用户请求头部header

- X-Forwarded-For
- X-Real-IP
- Proxy-Client-IP
- WL-Proxy-Client-IP
- RemoteAddr 

如果要获取的IP信息不在以上头部设置里，比如转发代理实现了HTTP_CLIENT_IP头设置，可以通过启动参数header追加加上（多个参数用半角,分隔）
```shell
$ showip -header=HTTP_CLIENT_IP
```
如果默认的获取IP头部顺序不符合要求，比如需要优先获取X-Real-IP头设置，可以通过启动参数header调整优先顺序
```shell
$ showip -header=X-Real-IP,X-Forwarded-For
```


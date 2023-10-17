# showip: 通用获取客户端IP和服务器IP的web服务




## 功能特点

- 支持客户端ip展示和api获取
- 支持服务器ip展示和api获取
- 支持本地ip的v4和v6输出
- 支持代理协议、转发协议中用户的真实ip
- 支持自定义转发协议、扩展协议头中ip（可在服务启动时指定header）

## 服务启动参数(shell)

> $ showip -port=80 -path=/showip 

- port：服务响应端口(默认端口80)
- path：服务响应的路径(默认路径/showip)
- header：自定义转发协议header名称参数，IP相关header名称（指定或优先获取header名称多个参数,分隔）

### 请求参数名称重命名和关闭参数响应

- format：请求格式参数名称设置（默认名称format），format=0时关闭请求格式参数
- mode：服务器模式参数名称设置，（默认名称mode），mode=0时关闭，即不再返回服务器IP信息
- via：响应头X-Via名称设置（默认名称X-Via），via=0时会关闭X-Via服务器IP的输出头（如下示例）
```text
- Response Header
  X-Via: 192.168.2.1
```

## 服务访问参数(web && api)

> http://{service-name}{:port}/{path}?format={text|json|array|xml|html}&obj={showip}&mode={host}

- format: 可选参数，输出格式，默认参数text可选参数text|json|array|xml|html
- obj: 可选参数，html和xml输出对象名称，html时为ul对象id名称，xml时为根节点名称
- mode: 可选参数，默认不传返回客户端IP信息，当mode=host时返回**服务器IP信息**

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
> xml的根节点名称可以通过obj参数设置
```xml
<? version="1.0" encoding="UTF-8" ?>
<showip>
    <RemoteAddress>192.168.2.13</RemoteAddress>
    <IP>192.168.2.13</IP>
</showip>
```
5. format=html 返回结果为html无序列表格式
> html中ul的id名称可以通过obj参数设置，样式固定为showip不会随指定名称改变
```html
<ul id="showip" class="showip">
    <li>192.168.2.13</li>
</ul>
```

## 关于客户端真实IP获取的逻辑顺序

程序默认情况会以下顺序，从用户请求协议头部header中，来获取用户的真实IP信息

1. X-Forwarded-For
2. X-Real-IP
3. Proxy-Client-IP
4. WL-Proxy-Client-IP
5. RemoteAddr 

如果要获取的IP信息不在以上头部设置里，比如转发代理HTTP_CLIENT_IP协议，可以通过启动参数header追加上（多个参数用半角,分隔），响应时会优先获取HTTP_CLIENT_IP协议中的IP
```shell
$ showip -header=HTTP_CLIENT_IP
```
如果默认的获取IP头部顺序不符合要求，比如需要优先获取X-Real-IP头设置，可以通过启动参数header调整优先顺序，之后会按调整后的顺序来响应
```shell
$ showip -header=X-Real-IP,X-Forwarded-For
```


# showip: 通用获取客户端IP和服务器IP的web服务

## 目录

* [功能特点](#功能特点)
* [程序安装](#编译安装源码)
* [启动参数(shell)](#服务启动参数)
* [访问参数(web & api)](#服务访问参数和输出格式)
* [客户端IP获取逻辑](#客户端IP获取逻辑)
* [程序自动运行](#启用守护进程)

## 功能特点

1. 支持客户端IP和服务器IP（网卡）Web获取和展示
2. 支持代理协议、转发协议头中的用户真实IP获取
3. 适应各种网络环境，同时支持IPv4和IPv6的IP
4. 支持API方式访问，接口支持XML、JSON、HTML等多种返回格式
5. 支持自定义转发IP协议、以及扩展协议头中的IP **（在服务启动时指定扩展header名称）**

## 编译安装源码

```shell
# go install github.com/tersergo/showip@latest
```

## 服务启动参数

> $ showip -port=80 -path=/showip -header=X-Real-IP,X-Forwarded-For -via=X-Via -format=format -mode=mode

- **port**：服务响应端口(默认端口80)

- **path**：服务响应的路径(默认路径/showip)

- **header**：自定义转发协议header名称参数，指定优先获取IP相关header头名称（多个名称用半角,分隔）

### 以下启动参数支持重命名和关闭响应输出

- *format*：format请求参数名称重命名和关闭设置（默认名称format）
  
  > format=0时关闭格式参数响应操作

- *mode*：mode请求参数重命名和关闭设置（默认名称mode）
  
  > mode=0时关闭服务器IP信息的响应操作

- **via**：Response响应头重命名和关闭设置（默认名称X-Via）
  
  > via=0时会关闭Response头部的输出默认会在Response Header输出一个X-Via头参数，对应内容为响应的服务器IP
  
  ```text
  - Response Header
      X-Via: 2.1
  ```
  
  > 如服务器ip是192.168.2.1 X-Via头会输出 2.1 *为安全考虑默认是服务器IP是后2位(IPV6是后3位）*

```text
$ showip -via=X-Power-By

- Response Header
    X-Power-By: 2.1
```

> $ showip -via=0 则会关闭在Response Header的输出

## 服务访问参数和输出格式

> http://{service-name}[:port]/{path}?[format=text]&[obj=]&[mode=host]

- format: 请求输出格式参数（可选，默认参数text）可选参数text,json,array,xml,html
- obj: 返回对象名称（可选，默认为showip），输出html时为ul对象id名称，xml时为根节点名称
- mode: 扩展参数（默认空，返回客户端IP信息），只有当**mode=host**时返回结果切换为**服务器IP（网卡）信息**

*注意：format和mode参数名称可以通过其同名的启动参数来改变请求参数名称，或者关闭该参数响应*

### text：文本格式输出(默认）

format参数为空或format=text时，返回第一个有效IP的纯文本格式

> http://localhost/showip

```text
IP: 127.0.0.1
```

如果支持ipV6版本时，返回以下结果

```text
IP: ::1
```

> http://localhost/showip?mode=host

当请求参数包含mode=host时，返回结果为**服务器IP**，返回格式不变。

### array：数组格式输出

format=array 返回结果为获取到的所有IP的数组形式

> http://{service-name}/showip?format=array

```json
["192.168.2.13","xx.xx.xx.xx"]
```

> http://{service-name}/showip?format=array&mode=host

当请求参数包含mode=host时，返回结果**服务器IP**，返回格式不变。

### json：JSON格式输出

format=json 返回结果为客户端IP的json格式

> http://{service-name}/showip?format=json

```json
{
   "IP": "192.168.2.13",
   "RemoteAddress": "192.168.2.13",
   /* 如果请求包含以下头部协议，返回会包含下列内容
    （通过启动参数header自定义的头名称也出现在这里） */
   "X-Forwarded-For": "xx.xx.xx.xx",
   "X-Real-IP": "xx.xx.xx.xx",
   "Proxy-Client-IP": "xx.xx.xx.xx",
   "WL-Proxy-Client-IP": "xx.xx.xx.xx"
}
```

> http://{service-name}/showip?format=json&mode=host

当请求参数包含mode=host时，返回结果为**服务器IP**的json格式

```json
{
  "IP": "192.168.2.1",
  "IPV4": "192.168.2.1",
  /* 如果有ipV6地址时，会包含下列内容 */
  "IPV6": "ff::aa:bb:cc:dd"
}
```

### xml：XML格式输出

format=xml 返回结果为客户端IP的xml格式，xml的根节点ClientIP名称可以通过obj参数设置

> http://{service-name}/showip?format=xml

```xml
<?version="1.0" encoding="UTF-8" ?>
<ClientIP>
    <IP>192.168.2.13</IP>
    <RemoteAddress>192.168.2.13</RemoteAddress>
    <!-- 如果请求包含以下头部协议，返回会包含下列内容
    （通过启动参数header自定义的头名称也出现在这里） --> 
    <X-Forwarded-For>xx.xx.xx.xx</X-Forwarded-For>
    <X-Real-IP>xx.xx.xx.xx</X-Real-IP>
    <Proxy-Client-IP>xx.xx.xx.xx</Proxy-Client-IP>
    <WL-Proxy-Client-IP>xx.xx.xx.xx</WL-Proxy-Client-IP>
</ClientIP>
```

> http://{service-name}/showip?format=xml&mode=host

当请求参数包含mode=host时，返回结果为**服务器IP**的xml格式

```xml
<?version="1.0" encoding="UTF-8" ?>
<ServerIP>
    <IP>192.168.2.1</IP>
    <IPV4>192.168.2.1</IPV4>
    <!-- 如果有ipV6地址时，会包含下列内容 -->
    <IPV6>ff::aa:bb:cc:dd</IPV6>
</ServerIP>
```

### html：HTML格式输出

format=html 返回结果为html无序列表格式，html中ul的id名称可以通过请求obj参数来设置，样式固定为showip不会随指定obj名称改变

> http://{service-name}/showip?format=html&obj=showip

```html
<ul id="showip" class="showip">
    <li>192.168.2.13</li>
</ul>
```

> http://{service-name}/showip?format=html&mode=host

当请求参数包含mode=host时，返回结果为**服务器IP**，返回格式不变。

## 客户端IP获取逻辑

### 默认获取顺序

程序默认情况会按照以下顺序，优先从用户请求协议的header中，来获取用户的真实IP信息

1. X-Forwarded-For
2. X-Real-IP
3. Proxy-Client-IP
4. WL-Proxy-Client-IP
5. RemoteAddr

### 自定义获取IP头和顺序调整

如果要获取的IP信息不在以上头部设置里，如转发实现了 HTTP_CLIENT_IP 协议，可以通过启动参数header追加上*（多个参数用半角,分隔）*，响应时会优先获取HTTP_CLIENT_IP协议中的IP

```shell
$ showip -header=HTTP_CLIENT_IP
```

如果默认的获取IP头部顺序不符合要求，如需要优先获取X-Real-IP设置，可以通过启动参数header调整优先顺序，之后会按调整后的顺序来响应

```shell
$ showip -header=X-Real-IP,X-Forwarded-For,Proxy-Client-IP
```

## 启用守护进程

### 支持自启动服务

> conf/showip_unit.service

```shell
- 1. 复制showip_unit.service
cp conf/showip_unit.service /usr/lib/systemd/system/showip.service
- 2. 开启开机自动运行
systemctl enable showip
- 3. 启动和关闭 showip
systemctl start showip
systemctl stop showip
```

### 支持supervisor

> conf/showip_supervisor.conf

```shell
- 1. 复制showip_supervisor.conf
cp conf/showip_supervisor.conf /etc/supervisor/conf.d/showip.conf
- 2. 启动运行 showip
supervisorctl start showip
- 3. 关闭 showip
supervisorctl stop showip
```

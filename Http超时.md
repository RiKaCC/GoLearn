# Golang的超时设置

## 客户端超时设置
下面是一个最简单的Get请求：
```
resp, err := http.Get(url)
if err != nil {
  // handle with error
} 

defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
  // handle with error
}

fmt.Println(stirng(body))
```
这是一个没有设置超时时间的get请求，最简单的设置超时的方式：使用http.Client的Timeout字段， 这个超时的时间计算是包括从连接(Dail)到读完response body, 代码如下：
```
c := &http.Client {
  Timeout: 30 * time.Second
}

resp, err := c.Get(url)

... ...
```

下面是更详细的一些粒度：
http.Transport.TLSHandshakeTimeout -> TLS的握手时间
官方的解释如下：
```
// TLSHandshakeTimeout specifies the maximum amount of time waiting to
// wait for a TLS handshake. Zero means no timeout.
TLSHandshakeTimeout time.Duration
```

http.Transport.ResponseHeaderTimeout -> 限制读取response header的时间
官方解释如下：
```
// ResponseHeaderTimeout, if non-zero, specifies the amount of
// time to wait for a server's response headers after fully
// writing the request (including its body, if any). This
// time does not include the time to read the response body.
ResponseHeaderTimeout time.Duration
```

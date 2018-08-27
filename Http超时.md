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

### 简单的Timeout
这是一个没有设置超时时间的get请求，最简单的设置超时的方式：使用http.Client的Timeout字段， 这个超时的时间计算是包括从连接(Dail)到读完response body, 代码如下：
```
c := &http.Client {
  Timeout: 30 * time.Second
}

resp, err := c.Get(url)

... ...
```
看一下官方文档对Timeout的解释：
```
// Timeout specifies a time limit for requests made by this
// Client. The timeout includes connection time, any
// redirects, and reading the response body. The timer remains
// running after Get, Head, Post, or Do return and will
// interrupt reading of the Response.Body.
//
// A Timeout of zero means no timeout.
//
// The Client cancels requests to the underlying Transport
// using the Request.Cancel mechanism. Requests passed
// to Client.Do may still set Request.Cancel; both will
// cancel the request.
//
// For compatibility, the Client will also use the deprecated
// CancelRequest method on Transport if found. New
// RoundTripper implementations should use Request.Cancel
// instead of implementing CancelRequest.
       
Timeout time.Duration
```
可以看到以下几个关注点：
1. timeout的时间计算：包含了连接时间，任何重定向和读取response body的时间
2. 如果不设置该值，则表示没有timeout
3. 若果取消timeout，则会取消本次请求

### 详细的Timeout
下面是更详细的一些粒度：

```
t := &http.Transport {
  TLSHandshakeTimeout:   10 * time.Second,
  ResponseHeaderTimeout: 10 * time.Second,
}

c := &http.Client {
  TransPort: t,
}
```

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

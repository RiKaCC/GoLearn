# Client

## type Client

1. 客户端的传输通常具有内部状态（会缓存TCP连接），所以我们在使用的时候，应该是重用客户端而不是根据需要去创建客户端。
2. 并发的时候，多个goroutine可以安全的使用一个client.
3. 关于重定向的一些东西，以后再写

```
type Client struct {
        // Transport就是client在调用HTTP请求的时候，去设置的一些规则
        // 没有设置, 使用DefaultTransport
        Transport RoundTripper

        // CheckRedirect设置如何处理重定向
        // If CheckRedirect is not nil, the client calls it before
        // following an HTTP redirect. The arguments req and via are
        // the upcoming request and the requests made already, oldest
        // first. If CheckRedirect returns an error, the Client's Get
        // method returns both the previous Response (with its Body
        // closed) and CheckRedirect's error (wrapped in a url.Error)
        // instead of issuing the Request req.
        // As a special case, if CheckRedirect returns ErrUseLastResponse,
        // then the most recent response is returned with its body
        // unclosed, along with a nil error.
        //
        // If CheckRedirect is nil, the Client uses its default policy,
        // which is to stop after 10 consecutive requests.
        CheckRedirect func(req *Request, via []*Request) error

        // Jar specifies the cookie jar.
        //
        // The Jar is used to insert relevant cookies into every
        // outbound Request and is updated with the cookie values
        // of every inbound Response. The Jar is consulted for every
        // redirect that the Client follows.
        //
        // If Jar is nil, cookies are only sent if they are explicitly
        // set on the Request.
        Jar CookieJar

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
}
```

## func (*Client) Do
Do函数的原型如下：
```
func (c *Client) Do(req *Request) (*Response, error)
```
1. Do函数就是发送HTTP请求和接收HTTP响应。
2. 如果client设置的策略有问题，或者HTTP会话失败了，会返回error.非2XX的返回码并不会导致error。
3. 


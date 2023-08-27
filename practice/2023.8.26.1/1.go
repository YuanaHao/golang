//构造客户端
var client http.Client
//构造GET请求
reqList, err := http.NewRequest("GET", URL, nil)
//构造POST请求
jar, err := cookiejar.New(nil)
if err != nil {
	panic(err)
}
Info := "muser=" + muserid + "&" + "passwd=" + password
var data = strings.NewReader(Info)
req, err := http.NewRequest("POST", URL, data)
req.Header.Set("Connection", "keep-alive")
req.Header.Set("Pragma", "no-cache")
req.Header.Set("Cache-Control", "no-cache")
req.Header.Set("Upgrade-Insecure-Requests", "1")
req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

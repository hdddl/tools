package linkProxyDownloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// ProxyHandler 链接反向代理
func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	// 如果是GET请求直接给他返回首页内容
	err := r.ParseForm()
	if err != nil {
		_, _ = fmt.Fprint(w, "parameter worry")
		return
	}
	link := r.Form.Get("link")
	if link == "" {
		_, _ = fmt.Fprint(w, "can get correct parameter")
		return
	}
	log.Printf("Proxy request url: %s\n", link)

	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0")
	res, err := client.Do(req)
	// 处理返回的错误消息
	if err != nil {
		log.Printf("request forward err %v\n", err)
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	// 设置文件大小
	w.Header().Set("content-length", res.Header.Get("content-length"))
	// 设置文件是用来下载的
	w.Header().Set("Content-Disposition", res.Header.Get("content-Disposition"))
	w.Header().Set("Content-Type", res.Header.Get("content-Type"))
	// 返回http头
	w.WriteHeader(res.StatusCode)

	// 返回Body
	_, _ = io.Copy(w, res.Body)
	_ = res.Body.Close()
}

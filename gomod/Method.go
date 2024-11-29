package gomod

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func Method_GET(address string) any {
	baseURL := address

	// 定义GET参数对象
	params := map[string]string{
		// "key1": "value",
		// "key2": "value",
	}

	// 将参数对象转换为url.Values类型
	urlParams := url.Values{}
	for key, value := range params {
		urlParams.Add(key, value)
	}

	fmt.Printf(urlParams.Encode())
	// 拼接url和参数
	URL := fmt.Sprintf("%s?%s", baseURL, urlParams.Encode())

	// 创建HTTP请求
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("创建请求失败", err)
		return "创建请求失败"
	}

	// 定义请求头对象
	headers := map[string]string{
		// "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
	}

	// 批量设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req) // 发送请求
	if err != nil {
		fmt.Println("发送请求失败", err)
		return "发送请求失败"
	}
	defer resp.Body.Close()

	fmt.Println("响应状态码:", resp.StatusCode)
	return resp.StatusCode
}

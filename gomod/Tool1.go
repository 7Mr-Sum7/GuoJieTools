package gomod

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Result struct {
	ResultContent string `json:"resultcontent"`
	ResultLen     int    `json:"resultlen"`
}

func Lifestyle(address string) Result {
	baseUrl := address

	// 字符串转切片
	arraybaseurl := regexp.MustCompile(`[,\s\n]+`)
	arrayurl := arraybaseurl.Split(baseUrl, -1)

	// 定义一个空切片
	var result []string
	var mu sync.Mutex     // 用于保护 result 切片的并发写入
	var wg sync.WaitGroup // 用于等待所有 Goroutine 完成

	// 创建一个 HTTP 客户端
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, value := range arrayurl {
		// 跳过空字符串
		if value == "" {
			continue
		}

		wg.Add(1) //每启动一个Goroutine增加一个计数

		// 请求函数
		go func(url string) {
			defer wg.Done() // 确保 Goroutine 完成时调用 Done
			// 创建HTTP请求
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("创建请求失败", err)
				return // 失败后直接返回，跳过该 URL
			}
			// 发送请求
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("发送请求失败", err)
				return // 失败后直接返回，跳过该 URL
			}
			defer resp.Body.Close()

			// 检查响应状态码
			if resp.StatusCode == 200 || resp.StatusCode == 302 || resp.StatusCode == 301 {
				// 使用 Mutex 来确保对 result 切片的并发写入是安全的
				mu.Lock()
				result = append(result, url)
				mu.Unlock()
			} else {
				fmt.Println("请求失败", url)
			}
		}(value)
	}

	// 等待所有Goroutine 完成
	wg.Wait()

	// 将切片转换为字符串
	resultcontent := strings.Join(result, "\n")
	resultlen := len(result)

	return Result{ResultContent: resultcontent, ResultLen: resultlen}
}

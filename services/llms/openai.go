package llms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/IanVzs/Snowflakes/models"
	"github.com/IanVzs/Snowflakes/pkgs/logging"
	"github.com/gorilla/websocket"
)

const (
	openaiURL      = "http://10.121.24.249:3000/v1/chat/completions"       // 请根据实际情况更改为对应的API URL
	openaiAPIToken = "sk-xeWlzTZVhH5RnemyB3Cd02F7F2044235Ae0eB2250276F07e" // 替换为你的OpenAI API密钥
	streamEndpoint = "wss://api.openai.com/stream"                         // 流式接口
)

func QAOpenAITest(content string) string {
	logging.Debug("QA输入, Q: " + content)
	return "hi"
}

func QAOpenAI(content string) string {
	// 构建请求体
	var repSB strings.Builder
	data := models.PayloadOpenAI{
		Messages:         []models.MessageOpenAI{models.MessageOpenAI{Role: "assistant", Content: "你好我是你的AI小助手,谷歌"}, models.MessageOpenAI{Role: "user", Content: content}},
		MaxTokens:        2048,
		Temperature:      0.5,
		TopP:             0,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Model:            "gpt-3.5-turbo",
	}
	PayloadOpenAIBytes, err := json.Marshal(data)
	if err != nil {

	}
	// 创建POST请求
	req, err := http.NewRequest("POST", openaiURL, bytes.NewReader(PayloadOpenAIBytes))
	if err != nil {
		fmt.Printf("Error creating request: %s\n", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openaiAPIToken)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)

	}
	fmt.Println(string(body))
	// 解析响应数据
	var response models.ResponseChatCompletionOpenAI
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Error parsing response JSON: %s\n", err)

	}

	// 输出结果
	fmt.Printf("Response: %+v\n", response)
	for _, choice := range response.Choices {
		fmt.Printf("Text: %s\n", choice.Message.Content)
		repSB.WriteString(choice.Message.Content)
	}
	return repSB.String()
}

// 这个方法还没有测试过
func noTestFunction_wsOpenAIAPI() {
	// 设置 WebSocket 连接参数
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial("ws://"+streamEndpoint, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// 设置读取消息的回调函数
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)

			}
			fmt.Printf("Received message: %s\n", message)
			// 在这里处理接收到的消息
		}
	}()

	// 发送认证信息到服务器
	err = conn.WriteMessage(websocket.TextMessage, []byte("Bearer "+openaiAPIToken))
	if err != nil {
		log.Fatal("write:", err)
	}

	// 保持程序运行，等待接收消息
	select {}
}

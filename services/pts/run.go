package pts

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/IanVzs/Snowflakes/models"
	"github.com/IanVzs/Snowflakes/pkgs/app"
	"github.com/IanVzs/Snowflakes/pkgs/e"
	"github.com/IanVzs/Snowflakes/pkgs/logging"
	"github.com/gin-gonic/gin"
)

func requestUrl(url_str string, proxyurl string, statusCodeMap *sync.Map) {
	// 创建一个请求对象
	logging.Infof("开始发送请求")
	req, err := http.NewRequest("GET", url_str, nil)
	if err != nil {
		return
	}
	// 设置请求头
	// 这里可以设置任意需要的头部信息，比如Content-Type, Authorization等
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:127.0) Gecko/20100101 Firefox/127.0")
	req.Header.Set("Host", "app.appsflyer.com")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Priority", "u=1")
	logging.Info(proxyurl)
	proxyURL, err := url.Parse(proxyurl)
	if err != nil {
		return
	}
	// 创建自定义 Transport，设置代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// 创建 HTTP 客户端，使用自定义 Transport
	client := &http.Client{
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	// resp, err := http.Get(url)
	if err != nil {
		logging.Errorf("url请求失败: url: %s, 错误: %v", url_str, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	count, ok := statusCodeMap.Load(resp.StatusCode)
	if ok {
		statusCodeMap.Store(resp.StatusCode, count.(int)+1)
		logging.Infof("url返回码: %d, 返回值: %+v", resp.StatusCode, string(body))
	} else {
		statusCodeMap.Store(resp.StatusCode, 1)
		logging.Infof("url返回码: %d, 返回信息: %+v", resp.StatusCode, string(body))
	}
}

// @Tags PTS Page
// @测试接口质量前端页面
// @Produce html
// @Router /v1/pts/index [get]
func OpenHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "pts"})
}

// @Tags PTS Run
// @测试接口质量
// @Accept  json
// @Produce json
// @Param question body models.PtsParams true "Question to be answered"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/pts/run [post]
func Run(c *gin.Context) {
	var sm sync.Map
	var wg sync.WaitGroup
	var pts_params models.PtsParams
	statusMap := make(map[int]int)
	appG := app.Gin{C: c}
	if err := appG.C.ShouldBindJSON(&pts_params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logging.Info("params: %+v", pts_params)
	start := time.Now().Unix()
	_start := time.Now().Unix()
	count := 0
	for i := 1; i <= pts_params.Count; i++ {
		count = count + 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			requestUrl(pts_params.Url, pts_params.Proxyurl, &sm)
		}()
		_now := time.Now().Unix()
		delta := _now - _start
		if delta <= 1 && count >= pts_params.SecondCount {
			time.Sleep(time.Duration(1-delta) * time.Second)
			logging.Debugf("Sleep: %d", 1-delta)
			_start = time.Now().Unix()
			count = 0
		}
	}
	wg.Wait()
	done := time.Now().Unix()
	sm.Range(func(key, value interface{}) bool {
		logging.Debugf("status map k: %d", key.(int))
		statusMap[key.(int)] = value.(int)
		return true
	})
	logging.Infof("status map is: %+v", statusMap)
	appG.Response(200, e.SUCCESS, models.RespPts{Total: 0, Cost: int(done) - int(start), StatusMap: statusMap})
}

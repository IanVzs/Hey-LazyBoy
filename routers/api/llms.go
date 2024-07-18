package api

import (
	"fmt"

	"github.com/IanVzs/Snowflakes/models"
	"github.com/IanVzs/Snowflakes/pkgs/app"
	"github.com/IanVzs/Snowflakes/pkgs/e"
	"github.com/IanVzs/Snowflakes/pkgs/logging"
	"github.com/IanVzs/Snowflakes/services/llms"
	"github.com/gin-gonic/gin"
)

// @Tags QA
// @Summary Answer a question
// @Accept  json
// @Produce json
// @Param question body models.LLMRequest true "Question to be answered"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/qa [post]
func QA(c *gin.Context) {
	// 构建请求体
	var reqBody models.LLMRequest
	appG := app.Gin{C: c}
	err := appG.C.ShouldBindJSON(&reqBody)
	logging.Info(fmt.Sprintf("QA reqBody.Content: %s", reqBody.Content))
	if err != nil || reqBody.Content == "" {
		appG.Response(400, e.INVALID_PARAMS, "无效的请求参数")
		return
	}
	// rspContent := llms.QAOpenAI(reqBody.Content)
	rspContent := llms.QAOpenAITest(reqBody.Content)
	appG.Response(200, e.SUCCESS, rspContent)
}

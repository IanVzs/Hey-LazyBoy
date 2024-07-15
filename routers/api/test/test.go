package test

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/IanVzs/Snowflakes/pkgs/app"
	"github.com/IanVzs/Snowflakes/pkgs/e"
	"github.com/IanVzs/Snowflakes/pkgs/logging"
	"github.com/IanVzs/Snowflakes/pkgs/setting"
	"github.com/IanVzs/Snowflakes/pkgs/util"
	"github.com/IanVzs/Snowflakes/services/test_service"
)

var Count int = 0

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/test/get [get]
func GoTest(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	testService := test_service.TestData{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	tests, err := testService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TEST_FAIL, nil)
		return
	}

	count := len(tests)
	time.Sleep(time.Second * 1)
	Count += 1
	logging.Info("Count: ", Count)

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tests,
		"total": count,
	})
}

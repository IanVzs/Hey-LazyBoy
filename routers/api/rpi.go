package api

import (
	"bytes"
	"net/http"
	"os"
	"os/exec"

	"github.com/IanVzs/Snowflakes/pkgs/app"
	"github.com/IanVzs/Snowflakes/pkgs/e"
	"github.com/IanVzs/Snowflakes/pkgs/logging"
	"github.com/gin-gonic/gin"
)

func executeBashCommand(command string) (string, error) {
	// 请确保你对执行的命令有足够的限制和验证
	// 本例未做任何安全限制，请不要在生产环境中使用
	cmd := exec.Command("bash", "-c", command)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// 运行命令
	err := cmd.Run()
	return out.String(), err
}

// @Tags rpi
// @Summary 执行树莓派命令
// @Produce json
// @Param command query string true "Command to execute"
// @Param secret query string true "Secret"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /rpi/execute [get]
func RPiCMDExecute(c *gin.Context) {
	// 构建请求体
	appG := app.Gin{C: c}
	secret := appG.C.Query("secret")
	command := appG.C.Query("command")
	rpi_secret := os.Getenv("RPI_SECRET")
	logging.Infof("接收到的指令为: %s", command)

	if secret != rpi_secret {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, "check error")
		return
	}
	// 执行bash命令
	output, err := executeBashCommand(string(command))
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, err.Error())
		return
	}

	// 返回命令执行结果
	appG.Response(200, e.SUCCESS, output)
}

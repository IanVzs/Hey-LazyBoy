package user

import (
	"net/http"

	"github.com/IanVzs/Snowflakes/models"
	"github.com/IanVzs/Snowflakes/pkgs/app"
	"github.com/IanVzs/Snowflakes/pkgs/e"
	"github.com/IanVzs/Snowflakes/pkgs/logging"
	"github.com/IanVzs/Snowflakes/pkgs/util"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User
	appG := app.Gin{C: c}
	if err := appG.C.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 保存用户到数据库的逻辑
	// ...
	token, err := util.GenerateToken(user.Username, user.Password)
	if err != nil {
		logging.Errorf("无法生成令牌: 用户名: %s, 错误: %v", user.Username, err)
	}
	appG.Response(200, e.SUCCESS, models.RespLogin{Token: token})
}

func Info(c *gin.Context) {
	appG := app.Gin{C: c}
	token, ok := appG.C.Params.Get("token")
	if ok {
		logging.Info("获取到的Token: " + token)
	} else {
		logging.Info("没有Token: ")
	}
	appG.Response(200, e.SUCCESS, models.RespInfo{
		Roles:        []string{"admin"},
		Introduction: "I am a super administrator",
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:         "Super Admin",
	})
}

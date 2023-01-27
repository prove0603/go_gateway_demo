package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuangjie/go_gateway_demo/public"
	"time"
)

type AdminSessionInfo struct {
	ID        int       `json:"id"`
	UserName  string    `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
}

type AdminLoginInput struct {
	// form 代表从请求中接收的字段
	//todo validate换行的话会失效
	Username string `json:"username" form:"username" comment:"管理员用户名" example:"admin" validate:"required,is_valid_username"`
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`
}

// Code generated by hertz generator.

package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Ping 测试 handler
//
//	@Summary		Ping测试
//	@Description	测试 Description
//
//	@Tags			Ping
//
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/ping [get]
//	@success		200	{string}	map[string]string
func Ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}

// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/TremblingV5/DouTok/applications/api/biz/handler"
	api "github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/TremblingV5/DouTok/applications/api/initialize/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/errno"
)

// Register .
//
//	@Tags			User用户相关
//
//	@Summary		用户注册
//	@Description	添加一个用户到数据库中
//	@Param			req		body		api.DouyinUserRegisterRequest	true	"用户信息"
//	@Success		200		{object}	user.DouyinUserResponse
//	@Failure		default	{object}	api.DouyinUserRegisterResponse
//	@router			/douyin/user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildUserRegisterResp(errno.ErrBind))
		return
	}

	if len(req.Username) == 0 || len(req.Password) == 0 {
		handler.SendResponse(c, handler.BuildUserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Register(ctx, rpc.UserClient, &user.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildUserRegisterResp(errno.ConvertErr(err)))
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}

// GetUserById .
//
//	@Tags		User用户相关
//
//	@Summary	通过用户ID获取用户
//	@Description
//	@Param		req		query		api.DouyinUserRequest	true	"指明需要获取的用户的参数"
//	@Success	200		{object}	user.DouyinUserResponse
//	@Failure	default	{object}	api.DouyinUserResponse
//	@router		/douyin/user [GET]
func GetUserById(ctx context.Context, c *app.RequestContext) {
	// 如果是需要授权访问的接口，则在进入时已经被中间件从 body 中获取 token 解析完成了，这里无需额外解析
	var err error
	var req api.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildGetUserResp(errno.ErrBind))
		return
	}

	resp, err := rpc.GetUserById(ctx, rpc.UserClient, &user.DouyinUserRequest{
		UserId: req.UserId,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildGetUserResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}

// Login 确实有登录的接口，但是业务逻辑是在JWT中，写在这里是为了生成接口文档
//
//	@Tags			User用户相关
//
//	@Summary		用户登录
//	@Description	输入账号密码登录获取Token
//	@Param			req		body		api.DouyinUserLoginRequest	true	"用户信息"
//	@Success		200		{object}	user.DouyinUserResponse
//	@Failure		default	{object}	api.DouyinUserLoginResponse
//	@router			/douyin/user/login [POST]
func Login() {

}

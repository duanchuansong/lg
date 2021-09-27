package ecode

import "github.com/duanchuansong/stk/ecode"

var (
	OK           = ecode.New(1, "success")
	RequestErr   = ecode.New(20000, "Parameter error")
	UserIDErr    = ecode.New(30000, "用户身份验证失败。")
	DevIDErr     = ecode.New(30001, "开发者身份验证失败。")
	SignErr      = ecode.New(40000, "签名验证失败。")
	ServerErr    = ecode.New(50000, "服务异常。")
	GatewayErr   = ecode.New(50001, "网关异常。")
	UserExist    = ecode.New(60001, "user exist")
	UserNotExist = ecode.New(60002, "user not exist")
)

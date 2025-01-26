// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	user "github.com/Yzc216/gomall/app/user/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newRegisterArgs,
		newRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newLoginArgs,
		newLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ResetPassword": kitex.NewMethodInfo(
		resetPasswordHandler,
		newResetPasswordArgs,
		newResetPasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateUserInfo": kitex.NewMethodInfo(
		updateUserInfoHandler,
		newUpdateUserInfoArgs,
		newUpdateUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SetUserRole": kitex.NewMethodInfo(
		setUserRoleHandler,
		newSetUserRoleArgs,
		newSetUserRoleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetUserInfo": kitex.NewMethodInfo(
		getUserInfoHandler,
		newGetUserInfoArgs,
		newGetUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetUserInfoList": kitex.NewMethodInfo(
		getUserInfoListHandler,
		newGetUserInfoListArgs,
		newGetUserInfoListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"BanUser": kitex.NewMethodInfo(
		banUserHandler,
		newBanUserArgs,
		newBanUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.RegisterReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Register(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RegisterArgs:
		success, err := handler.(user.UserService).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *user.RegisterReq
}

func (p *RegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.RegisterReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.RegisterReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *user.RegisterReq

func (p *RegisterArgs) GetReq() *user.RegisterReq {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RegisterResult struct {
	Success *user.RegisterResp
}

var RegisterResult_Success_DEFAULT *user.RegisterResp

func (p *RegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.RegisterResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(user.RegisterResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *user.RegisterResp {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.RegisterResp)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RegisterResult) GetResult() interface{} {
	return p.Success
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.LoginReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *LoginArgs:
		success, err := handler.(user.UserService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *user.LoginReq
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.LoginReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(user.LoginReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *user.LoginReq

func (p *LoginArgs) GetReq() *user.LoginReq {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *LoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type LoginResult struct {
	Success *user.LoginResp
}

var LoginResult_Success_DEFAULT *user.LoginResp

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.LoginResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(user.LoginResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *user.LoginResp {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.LoginResp)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginResult) GetResult() interface{} {
	return p.Success
}

func resetPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.ResetPasswordReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).ResetPassword(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ResetPasswordArgs:
		success, err := handler.(user.UserService).ResetPassword(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ResetPasswordResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newResetPasswordArgs() interface{} {
	return &ResetPasswordArgs{}
}

func newResetPasswordResult() interface{} {
	return &ResetPasswordResult{}
}

type ResetPasswordArgs struct {
	Req *user.ResetPasswordReq
}

func (p *ResetPasswordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.ResetPasswordReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ResetPasswordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ResetPasswordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ResetPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ResetPasswordArgs) Unmarshal(in []byte) error {
	msg := new(user.ResetPasswordReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ResetPasswordArgs_Req_DEFAULT *user.ResetPasswordReq

func (p *ResetPasswordArgs) GetReq() *user.ResetPasswordReq {
	if !p.IsSetReq() {
		return ResetPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ResetPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ResetPasswordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ResetPasswordResult struct {
	Success *user.ResetPasswordResp
}

var ResetPasswordResult_Success_DEFAULT *user.ResetPasswordResp

func (p *ResetPasswordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.ResetPasswordResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ResetPasswordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ResetPasswordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ResetPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ResetPasswordResult) Unmarshal(in []byte) error {
	msg := new(user.ResetPasswordResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ResetPasswordResult) GetSuccess() *user.ResetPasswordResp {
	if !p.IsSetSuccess() {
		return ResetPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ResetPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.ResetPasswordResp)
}

func (p *ResetPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ResetPasswordResult) GetResult() interface{} {
	return p.Success
}

func updateUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UpdateUserInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateUserInfoArgs:
		success, err := handler.(user.UserService).UpdateUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateUserInfoArgs() interface{} {
	return &UpdateUserInfoArgs{}
}

func newUpdateUserInfoResult() interface{} {
	return &UpdateUserInfoResult{}
}

type UpdateUserInfoArgs struct {
	Req *user.UpdateUserInfoReq
}

func (p *UpdateUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserInfoArgs_Req_DEFAULT *user.UpdateUserInfoReq

func (p *UpdateUserInfoArgs) GetReq() *user.UpdateUserInfoReq {
	if !p.IsSetReq() {
		return UpdateUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserInfoResult struct {
	Success *user.UpdateUserInfoResp
}

var UpdateUserInfoResult_Success_DEFAULT *user.UpdateUserInfoResp

func (p *UpdateUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateUserInfoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserInfoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserInfoResult) GetSuccess() *user.UpdateUserInfoResp {
	if !p.IsSetSuccess() {
		return UpdateUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateUserInfoResp)
}

func (p *UpdateUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserInfoResult) GetResult() interface{} {
	return p.Success
}

func setUserRoleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.SetUserRoleReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).SetUserRole(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SetUserRoleArgs:
		success, err := handler.(user.UserService).SetUserRole(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SetUserRoleResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSetUserRoleArgs() interface{} {
	return &SetUserRoleArgs{}
}

func newSetUserRoleResult() interface{} {
	return &SetUserRoleResult{}
}

type SetUserRoleArgs struct {
	Req *user.SetUserRoleReq
}

func (p *SetUserRoleArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.SetUserRoleReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SetUserRoleArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SetUserRoleArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SetUserRoleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SetUserRoleArgs) Unmarshal(in []byte) error {
	msg := new(user.SetUserRoleReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SetUserRoleArgs_Req_DEFAULT *user.SetUserRoleReq

func (p *SetUserRoleArgs) GetReq() *user.SetUserRoleReq {
	if !p.IsSetReq() {
		return SetUserRoleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetUserRoleArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SetUserRoleArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SetUserRoleResult struct {
	Success *user.SetUserRoleResp
}

var SetUserRoleResult_Success_DEFAULT *user.SetUserRoleResp

func (p *SetUserRoleResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.SetUserRoleResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SetUserRoleResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SetUserRoleResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SetUserRoleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SetUserRoleResult) Unmarshal(in []byte) error {
	msg := new(user.SetUserRoleResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetUserRoleResult) GetSuccess() *user.SetUserRoleResp {
	if !p.IsSetSuccess() {
		return SetUserRoleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetUserRoleResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.SetUserRoleResp)
}

func (p *SetUserRoleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetUserRoleResult) GetResult() interface{} {
	return p.Success
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserInfoArgs:
		success, err := handler.(user.UserService).GetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserInfoArgs() interface{} {
	return &GetUserInfoArgs{}
}

func newGetUserInfoResult() interface{} {
	return &GetUserInfoResult{}
}

type GetUserInfoArgs struct {
	Req *user.GetUserInfoReq
}

func (p *GetUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoArgs_Req_DEFAULT *user.GetUserInfoReq

func (p *GetUserInfoArgs) GetReq() *user.GetUserInfoReq {
	if !p.IsSetReq() {
		return GetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserInfoResult struct {
	Success *user.GetUserInfoResp
}

var GetUserInfoResult_Success_DEFAULT *user.GetUserInfoResp

func (p *GetUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserInfoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoResult) GetSuccess() *user.GetUserInfoResp {
	if !p.IsSetSuccess() {
		return GetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserInfoResp)
}

func (p *GetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserInfoResult) GetResult() interface{} {
	return p.Success
}

func getUserInfoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserInfoListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserInfoList(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserInfoListArgs:
		success, err := handler.(user.UserService).GetUserInfoList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserInfoListArgs() interface{} {
	return &GetUserInfoListArgs{}
}

func newGetUserInfoListResult() interface{} {
	return &GetUserInfoListResult{}
}

type GetUserInfoListArgs struct {
	Req *user.GetUserInfoListReq
}

func (p *GetUserInfoListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserInfoListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserInfoListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserInfoListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserInfoListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoListArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoListArgs_Req_DEFAULT *user.GetUserInfoListReq

func (p *GetUserInfoListArgs) GetReq() *user.GetUserInfoListReq {
	if !p.IsSetReq() {
		return GetUserInfoListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserInfoListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserInfoListResult struct {
	Success *user.GetUserInfoListResp
}

var GetUserInfoListResult_Success_DEFAULT *user.GetUserInfoListResp

func (p *GetUserInfoListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserInfoListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserInfoListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserInfoListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserInfoListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoListResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserInfoListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoListResult) GetSuccess() *user.GetUserInfoListResp {
	if !p.IsSetSuccess() {
		return GetUserInfoListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoListResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserInfoListResp)
}

func (p *GetUserInfoListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserInfoListResult) GetResult() interface{} {
	return p.Success
}

func banUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.BanUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).BanUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *BanUserArgs:
		success, err := handler.(user.UserService).BanUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*BanUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newBanUserArgs() interface{} {
	return &BanUserArgs{}
}

func newBanUserResult() interface{} {
	return &BanUserResult{}
}

type BanUserArgs struct {
	Req *user.BanUserReq
}

func (p *BanUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.BanUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *BanUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *BanUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *BanUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *BanUserArgs) Unmarshal(in []byte) error {
	msg := new(user.BanUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var BanUserArgs_Req_DEFAULT *user.BanUserReq

func (p *BanUserArgs) GetReq() *user.BanUserReq {
	if !p.IsSetReq() {
		return BanUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *BanUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *BanUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type BanUserResult struct {
	Success *user.BanUserResp
}

var BanUserResult_Success_DEFAULT *user.BanUserResp

func (p *BanUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.BanUserResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *BanUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *BanUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *BanUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *BanUserResult) Unmarshal(in []byte) error {
	msg := new(user.BanUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *BanUserResult) GetSuccess() *user.BanUserResp {
	if !p.IsSetSuccess() {
		return BanUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *BanUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.BanUserResp)
}

func (p *BanUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *BanUserResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, Req *user.RegisterReq) (r *user.RegisterResp, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *user.LoginReq) (r *user.LoginResp, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ResetPassword(ctx context.Context, Req *user.ResetPasswordReq) (r *user.ResetPasswordResp, err error) {
	var _args ResetPasswordArgs
	_args.Req = Req
	var _result ResetPasswordResult
	if err = p.c.Call(ctx, "ResetPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserInfo(ctx context.Context, Req *user.UpdateUserInfoReq) (r *user.UpdateUserInfoResp, err error) {
	var _args UpdateUserInfoArgs
	_args.Req = Req
	var _result UpdateUserInfoResult
	if err = p.c.Call(ctx, "UpdateUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetUserRole(ctx context.Context, Req *user.SetUserRoleReq) (r *user.SetUserRoleResp, err error) {
	var _args SetUserRoleArgs
	_args.Req = Req
	var _result SetUserRoleResult
	if err = p.c.Call(ctx, "SetUserRole", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq) (r *user.GetUserInfoResp, err error) {
	var _args GetUserInfoArgs
	_args.Req = Req
	var _result GetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserInfoList(ctx context.Context, Req *user.GetUserInfoListReq) (r *user.GetUserInfoListResp, err error) {
	var _args GetUserInfoListArgs
	_args.Req = Req
	var _result GetUserInfoListResult
	if err = p.c.Call(ctx, "GetUserInfoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BanUser(ctx context.Context, Req *user.BanUserReq) (r *user.BanUserResp, err error) {
	var _args BanUserArgs
	_args.Req = Req
	var _result BanUserResult
	if err = p.c.Call(ctx, "BanUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

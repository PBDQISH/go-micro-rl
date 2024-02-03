package handler

import (
	"context"
	"go-micro-rl/server/model"
	"go-micro-rl/server/utils"
	user "go-micro-rl/service"
)

type User struct{}

func (u *User) Register(ctx context.Context, req *user.RegReq, rsp *user.Response) error {
	err := model.RegisterUser(req.Mobile, req.Password)
	if err != nil {
		rsp.Errno = utils.MSG_DBERR
		rsp.Errmsg = utils.RecodeText(utils.MSG_DBERR)
	} else {
		rsp.Errno = utils.MSG_OK
		rsp.Errmsg = utils.RecodeText(utils.MSG_OK)
	}
	return nil
}

func (u *User) Login(ctx context.Context, req *user.RegReq, rsp *user.Responsel) error {
	name, err := model.Login(req.Mobile, req.Password)
	if err != nil {
		rsp.Errno = utils.MSG_DBERR
		rsp.Errmsg = utils.RecodeText(utils.MSG_DBERR)
	} else {
		rsp.Errno = utils.MSG_OK
		rsp.Errmsg = utils.RecodeText(utils.MSG_OK)
	}
	rsp.Name = name
	return nil
}

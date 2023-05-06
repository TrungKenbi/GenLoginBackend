package logic

import (
	"context"

	"genlogin/app/user/api/internal/svc"
	"genlogin/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeLogic {
	return &MeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeLogic) Me(req *types.MeReq) (resp *types.MeReply, err error) {
	// todo: add your logic here and delete this line

	resp = &types.MeReply{
		CreatedAt: "",
		ID:        0,
		Email:     "",
		FirstName: "Trung",
	}

	return
}

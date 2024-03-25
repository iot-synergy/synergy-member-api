package member

import (
	"context"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByForeinIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMemberByForeinIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByForeinIdLogic {
	return &GetMemberByForeinIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetMemberByForeinIdLogic) GetMemberByForeinId(req *types.IDStringReq) (resp *types.MemberInfoResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetMemberByForeinId2(l.ctx, &mms.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.MemberInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: int(data.Code),
			Msg:  data.Msg,
		},
		Data: types.MemberInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Data.MemberInfo.Id,
				CreatedAt: data.Data.MemberInfo.CreatedAt,
				UpdatedAt: data.Data.MemberInfo.UpdatedAt,
			},
			Status:   data.Data.MemberInfo.Status,
			Username: data.Data.MemberInfo.Username,
			Nickname: data.Data.MemberInfo.Nickname,
			RankId:   data.Data.MemberInfo.RankId,
			Mobile:   data.Data.MemberInfo.Mobile,
			Email:    data.Data.MemberInfo.Email,
			Avatar:   data.Data.MemberInfo.Avatar,
		},
	}, nil
}

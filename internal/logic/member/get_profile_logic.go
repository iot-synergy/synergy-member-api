package member

import (
	"context"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	userId string
}

var MemberRankData = make(map[uint64]string)

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		userId: ctx.Value("userId").(string),
	}
}

// genRankCache used to generate cache for member rank to improve performance
func (l *GetProfileLogic) genRankCache() error {
	list, err := l.svcCtx.MmsRpc.GetMemberRankList(l.ctx, &mms.MemberRankListReq{
		Page:     1,
		PageSize: 1000,
	})
	if err != nil {
		return err
	}

	for _, v := range list.Data {
		MemberRankData[*v.Id] = *v.Name
	}

	return err
}
func (l *GetProfileLogic) GetProfile() (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.MmsRpc.GetMemberById(l.ctx, &mms.UUIDReq{
		Id: l.userId,
	})

	if err != nil {
		return nil, err
	}

	// get rank data
	if len(MemberRankData) == 0 {
		err = l.genRankCache()
		if err != nil {
			return nil, err
		}
	}

	resp = &types.LoginResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, "login.loginSuccessTitle")},
		Data: types.LoginInfo{
			UserId:   *result.Id,
			RankId:   result.GetRankCode(),
			Nickname: result.GetNickname(),
			RankName: l.svcCtx.Trans.Trans(l.ctx, MemberRankData[result.GetRankId()]),
			Avatar:   result.GetAvatar(),
			Expire:   uint64(result.GetExpiredAt()),
		},
	}
	return resp, nil
}

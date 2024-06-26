package memberrank

import (
	"context"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetReplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetReplyListLogic {
	return &AdminGetReplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetReplyListLogic) AdminGetReplyList(req *types.ReplyListReqVo) (resp *types.ReplyListRespVo, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("userId").(string)
	page := int64(req.Page)
	pageSize := int64(req.PageSize)
	list, err := l.svcCtx.MmsRpc.AdminGetReplyList(l.ctx, &mms.ReplyReq{
		AdminId:  &userId,
		Page:     &page,
		PageSize: &pageSize,
	})
	if err != nil {
		return &types.ReplyListRespVo{
			BaseMsgResp: types.BaseMsgResp{
				Code: -1,
				Msg:  err.Error(),
			},
		}, nil
	}
	vos := make([]types.ReplyRespVo, 0)
	for _, info := range list.ReplyList {
		vos = append(vos, types.ReplyRespVo{
			Id:         info.GetId(),
			CommentId:  info.GetCommentId(),
			Reply:      info.GetReply(),
			AdminId:    info.GetAdminId(),
			AdminName:  info.GetAdminName(),
			CreateTime: info.GetCreateTime(),
			UpdateTime: info.GetUpdateTime(),
		})
	}

	return &types.ReplyListRespVo{
		BaseMsgResp: types.BaseMsgResp{
			Code: 0,
			Msg:  "成功",
		},
		Data: types.ReplyListRespData{
			List:  vos,
			Total: int(*list.Count),
		},
	}, nil
}

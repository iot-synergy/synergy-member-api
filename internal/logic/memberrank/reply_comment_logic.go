package memberrank

import (
	"context"
	"time"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyCommentLogic {
	return &ReplyCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplyCommentLogic) ReplyComment(req *types.ReplyReqVo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("userId").(string)
	now := time.Now().Unix()
	_, err = l.svcCtx.MmsRpc.ReplyComment(l.ctx,
		&mms.ReplyInfo{
			Id:         new(int64),
			CommentId:  &req.CommentId,
			Reply:      &req.Reply,
			AdminId:    &userId,
			AdminName:  &req.AdminName,
			CreateTime: &now,
			UpdateTime: &now,
		},
	)

	if err != nil {
		return &types.BaseMsgResp{
			Code: -1,
			Msg:  err.Error(),
		}, err
	}

	return &types.BaseMsgResp{
		Code: 0,
		Msg:  "",
	}, err
}

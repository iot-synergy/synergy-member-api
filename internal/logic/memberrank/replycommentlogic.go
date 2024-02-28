package memberrank

import (
	"context"
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

func (l *ReplyCommentLogic) ReplyComment(req *types.ReplyReqVo) (resp string, err error) {
	// todo: add your logic here and delete this line
	comment, err := l.svcCtx.MmsRpc.ReplyComment(l.ctx,
		&mms.ReplyInfo{
			CommentId: &req.CommentId,
			Reply:     &req.Reply,
			AdminId:   &req.AdminId,
			AdminName: &req.AdminName,
		},
	)

	if err != nil {
		return "", err
	}

	return comment.Msg, err
}

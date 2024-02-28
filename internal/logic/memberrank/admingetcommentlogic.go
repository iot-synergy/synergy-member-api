package memberrank

import (
	"context"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetCommentLogic {
	return &AdminGetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetCommentLogic) AdminGetComment(req *types.CommentIdReqVo) (resp *types.CommentRespVo, err error) {
	// todo: add your logic here and delete this line
	comment, err := l.svcCtx.MmsRpc.AdminGetComment(l.ctx, &mms.CommentIdReq{Id: &req.Id})

	if err != nil {
		return nil, err
	}
	vos := make([]types.ReplyRespVo, 0)
	reply := comment.GetReply()
	for _, info := range reply {
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
	return &types.CommentRespVo{
		Id:          comment.GetId(),
		Title:       comment.GetTitle(),
		Content:     comment.GetContent(),
		MemberId:    comment.GetMemberId(),
		Create_time: comment.GetCreateTime(),
		Update_time: comment.GetUpdateTime(),
		Reply:       vos,
	}, err
}

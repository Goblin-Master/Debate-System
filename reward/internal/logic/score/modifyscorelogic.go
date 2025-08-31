package score

import (
	"Debate-System/reward/internal/repo"
	"context"
	"strconv"

	"Debate-System/reward/internal/svc"
	"Debate-System/reward/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyScoreLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	rewardRepo *repo.RewardRepo
}

func NewModifyScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyScoreLogic {
	return &ModifyScoreLogic{
		Logger:     logx.WithContext(ctx),
		ctx:        ctx,
		svcCtx:     svcCtx,
		rewardRepo: repo.NewRewardRepo(ctx, svcCtx),
	}
}

func (l *ModifyScoreLogic) ModifyScore(req *types.ModifyScoreReq) (resp *types.ModifyScoreResp, err error) {
	user_id, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, DEFAULT_ERROR
	}
	err = l.rewardRepo.ModifyScore(user_id, req.Score)
	if err != nil {
		return nil, MODIFY_SCORE_REEOR
	}
	resp = &types.ModifyScoreResp{
		Message: "修改积分成功",
	}
	return resp, nil
}

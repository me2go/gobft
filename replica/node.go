package replica

import (
	"context"

	"github.com/me2go/gobft/protos"
)

func New() *replica {
	return &replica{}
}

type replica struct {
}

func (r *replica) Sync(ctx context.Context,
	msg *protos.SyncMsg) (resp *protos.SyncReply, err error) {
	return
}

func (r *replica) PrePrepare(ctx context.Context,
	msg *protos.PrePrepareMsg) (resp *protos.Response, err error) {
	return
}

func (r *replica) Prepare(ctx context.Context,
	msg *protos.PrepareMsg) (resp *protos.Response, err error) {
	return
}

func (r *replica) Commit(ctx context.Context,
	msg *protos.CommitMsg) (resp *protos.Response, err error) {
	return
}

package keeper

import (
	"context"

	"m25/x/gold/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgTransferResponse{}, nil
}

package keeper

import (
	"context"

	"m25/x/gold/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransferTo(goCtx context.Context, msg *types.MsgTransferTo) (*types.MsgTransferToResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	nft, found := k.Keeper.GetNft(ctx, msg.NftId)
	if !found {
		panic("NFT not found")
	}

	if nft.Owner != msg.Creator {
		panic("Not authorized")
	}

	nft.Owner = msg.To

	k.Keeper.SetNft(ctx, nft)

	return &types.MsgTransferToResponse{}, nil
}

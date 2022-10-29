package keeper

import (
	"context"

	"m25/x/gold/types"

	nft_service "m25/x/gold/nft"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	newNft := nft_service.Mint(msg.Creator)

	storeNft := types.Nft{
		Index: newNft.GetIndex(),
		Owner: newNft.GetOwner(),
		Uri:   newNft.GetUri(),
	}

	k.Keeper.SetNft(ctx, storeNft)

	return &types.MsgMintResponse{
		NftId: storeNft.Index,
	}, nil
}

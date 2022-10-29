package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTransferTo = "transfer_to"

var _ sdk.Msg = &MsgTransferTo{}

func NewMsgTransferTo(creator string, nftId string, to string) *MsgTransferTo {
	return &MsgTransferTo{
		Creator: creator,
		NftId:   nftId,
		To:      to,
	}
}

func (msg *MsgTransferTo) Route() string {
	return RouterKey
}

func (msg *MsgTransferTo) Type() string {
	return TypeMsgTransferTo
}

func (msg *MsgTransferTo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferTo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferTo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgVerify{}

func NewMsgVerify(creator string, alias string, user string, kind string, target string) *MsgVerify {
	return &MsgVerify{
		Creator: creator,
		Alias:   alias,
		User:    user,
		Kind:    kind,
		Target:  target,
	}
}

func (msg *MsgVerify) Route() string {
	return RouterKey
}

func (msg *MsgVerify) Type() string {
	return "Verify"
}

func (msg *MsgVerify) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVerify) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVerify) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

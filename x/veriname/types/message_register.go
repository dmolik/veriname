package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRegister{}

func NewMsgRegister(creator string, alias string, user string, kind string, target string, payload string) *MsgRegister {
	return &MsgRegister{
		Creator: creator,
		Alias:   alias,
		User:    user,
		Kind:    kind,
		Target:  target,
		Payload: payload,
	}
}

func (msg *MsgRegister) Route() string {
	return RouterKey
}

func (msg *MsgRegister) Type() string {
	return "Register"
}

func (msg *MsgRegister) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegister) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegister) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

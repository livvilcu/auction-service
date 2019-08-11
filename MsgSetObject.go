package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

type MsgSetName struct {
	object  string
	owner sdk.AccAddress
}

func NewMsgSetObject(object string,  owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		object:  object,
		owner: owner,
	}
}

// ar trebui sa returneze numele modulului
func (msg MsgSetName) Route() string { return "RouterKey" }

// tipul ar trebui sa returneze actiunea
func (msg MsgSetName) Type() string { return "set_object" }

func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.owner.Empty() {
		return sdk.ErrNoSignatures("owner is not there")
	}
	return nil
}

// ?
func (msg MsgSetName) GetSignBytes() []byte {
	 // auth module?
	return sdk.MustSortJSON(auth.ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required - ?
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.owner}
}

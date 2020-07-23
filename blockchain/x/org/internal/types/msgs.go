package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MsgTypeAddEmployee    = "msg_add_employee"
	MsgTypeUpdateEmployee = "msg_update_employee"
)

type MsgAddEmployee struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Address    string `json:"address"`
	Skills     Skills `json:"skills"`

	From sdk.AccAddress `json:"from"`
}

func NewMsgAddEmployee(name, dept, address string, skills Skills, from sdk.AccAddress) MsgAddEmployee {
	return MsgAddEmployee{
		Name:       name,
		Department: dept,
		Address:    address,
		Skills:     skills,
		From:       from,
	}
}

var _ sdk.Msg = MsgAddEmployee{}

func (m MsgAddEmployee) Route() string {
	return RouterKey
}

func (m MsgAddEmployee) Type() string {
	return MsgTypeAddEmployee
}

func (m MsgAddEmployee) ValidateBasic() error {
	if len(m.Name) == 0 {
		return sdkerrors.Wrap(ErrInvalidInput, fmt.Sprintf("%s field is empty", m.Name))
	}

	return nil
}

func (m MsgAddEmployee) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddEmployee) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

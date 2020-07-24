package types

import (
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MsgTypeAddEmployee     = "msg_add_employee"
	MsgTypeUpdateEmployee  = "msg_update_employee"
	MsgTypeDeleteEmployee  = "msg_delete_employee"
	MsgTypeRestoreEmployee = "msg_restore_employee"
)

type MsgAddEmployeeInfo struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Address    string `json:"address"`
	Skills     Skills `json:"skills"`
	
	From sdk.AccAddress `json:"from"`
}

func NewMsgAddEmployeeInfo(name, dept, address string, skills Skills, from sdk.AccAddress) MsgAddEmployeeInfo {
	return MsgAddEmployeeInfo{
		Name:       name,
		Department: dept,
		Address:    address,
		Skills:     skills,
		From:       from,
	}
}

var _ sdk.Msg = MsgAddEmployeeInfo{}

func (m MsgAddEmployeeInfo) Route() string {
	return RouterKey
}

func (m MsgAddEmployeeInfo) Type() string {
	return MsgTypeAddEmployee
}

func (m MsgAddEmployeeInfo) ValidateBasic() error {
	if len(m.Name) == 0 {
		return sdkerrors.Wrap(ErrInvalidInput, fmt.Sprintf("%s field is empty", m.Name))
	}
	
	return nil
}

func (m MsgAddEmployeeInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgAddEmployeeInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// --------------------------------------
type MsgUpdateEmployeeInfo struct {
	Id         string         `json:"id"`
	Address    string         `json:"address"`
	Department string         `json:"department"`
	Skills     Skills         `json:"skills"`
	From       sdk.AccAddress `json:"from"`
}

func NewMsgUpdateEmployeeInfo(id, address, dept string, skills Skills, addr sdk.AccAddress) MsgUpdateEmployeeInfo {
	return MsgUpdateEmployeeInfo{
		Id:         id,
		Address:    address,
		Department: dept,
		Skills:     skills,
		From:       addr,
	}
}

var _ sdk.Msg = MsgUpdateEmployeeInfo{}

func (m MsgUpdateEmployeeInfo) Route() string {
	return RouterKey
}

func (m MsgUpdateEmployeeInfo) Type() string {
	return MsgTypeUpdateEmployee
}

func (m MsgUpdateEmployeeInfo) ValidateBasic() error {
	if len(m.Id) == 0 {
		return sdkerrors.Wrap(ErrInvalidInput, "id not found")
	}
	
	return nil
}

func (m MsgUpdateEmployeeInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgUpdateEmployeeInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// -----------------------------------------------
type MsgDeleteEmployeeInfo struct {
	Id     string         `json:"id"`
	Remove bool           `json:"remove"`
	From   sdk.AccAddress `json:"from"`
}

func NewMsgDeleteEmployeeInfo(id string, remove bool, addr sdk.AccAddress) MsgDeleteEmployeeInfo {
	return MsgDeleteEmployeeInfo{
		Id:     id,
		Remove: remove,
		From:   addr,
	}
}

var _ sdk.Msg = MsgDeleteEmployeeInfo{}

func (m MsgDeleteEmployeeInfo) Route() string {
	return RouterKey
}

func (m MsgDeleteEmployeeInfo) Type() string {
	return MsgTypeDeleteEmployee
}

func (m MsgDeleteEmployeeInfo) ValidateBasic() error {
	if len(m.Id) == 0 {
		return sdkerrors.Wrap(ErrInvalidInput, fmt.Sprintf("invalid id %s", m.Id))
	}
	
	return nil
}

func (m MsgDeleteEmployeeInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgDeleteEmployeeInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

// ---------------------------------------------

type MsgRestoreEmployeeInfo struct {
	Id   string `json:"id"`
	From sdk.AccAddress
}

func NewMsgRestoreEmployeeInfo(id string, addr sdk.AccAddress) MsgRestoreEmployeeInfo {
	return MsgRestoreEmployeeInfo{
		Id:   id,
		From: addr,
	}
}

var _ sdk.Msg = MsgRestoreEmployeeInfo{}

func (m MsgRestoreEmployeeInfo) Route() string {
	return RouterKey
}

func (m MsgRestoreEmployeeInfo) Type() string {
	return MsgTypeRestoreEmployee
}

func (m MsgRestoreEmployeeInfo) ValidateBasic() error {
	if len(m.Id) == 0 {
		return sdkerrors.Wrap(ErrInvalidInput, fmt.Sprintf("invalid id %s", m.Id))
	}
	return nil
}

func (m MsgRestoreEmployeeInfo) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m MsgRestoreEmployeeInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

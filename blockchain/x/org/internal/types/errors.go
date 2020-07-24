package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidInput = sdkerrors.Register(ModuleName, 11, "invalid input")
	ErrEmployee     = sdkerrors.Register(ModuleName, 12, "error employee")
)

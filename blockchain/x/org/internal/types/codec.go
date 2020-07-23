package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

var (
	amino = codec.New()

	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddEmployee{}, "Employee/msg_add_employee", nil)
	cdc.RegisterConcrete(Person{}, "Person", nil)
	cdc.RegisterConcrete(Employee{}, "Person/Employee", nil)

}

func init() {
	RegisterCodec(amino)
	codec.RegisterCrypto(amino)
	amino.Seal()
}

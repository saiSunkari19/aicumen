package org

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"github.com/saiSunkari19/aicumen/blockchain/x/org/internal/keeper"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, genState GenesisState) {

}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) GenesisState {
	return GenesisState{}
}

package keeper

import (
	"testing"
	
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"
	
	"github.com/saiSunkari19/aicumen/blockchain/x/org/internal/types"
)

func CreateTestInput(t *testing.T) (Keeper, sdk.Context) {
	keyOrg := sdk.NewKVStoreKey(types.ModuleName)
	
	mdb := db.NewMemDB()
	ms := store.NewCommitMultiStore(mdb)
	ms.MountStoreWithDB(keyOrg, sdk.StoreTypeIAVL, mdb)
	require.Nil(t, ms.LoadLatestVersion())
	
	cdc := MakeTestCodec()
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "chain-id"}, false, log.NewNopLogger())
	
	orgKeeper := NewKeeper(cdc, keyOrg)
	return orgKeeper, ctx
}

func MakeTestCodec() *codec.Codec {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)
	auth.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	return cdc
}

var TestEmployee = types.Employee{
	Person: types.Person{
		Name:    "Alice",
		Address: "UAE",
		Skills:  []string{"c", "java"},
	},
	ID:         "employee1",
	Department: "mg",
	Status:     1,
}

var (
	TestPrivKey1 = ed25519.GenPrivKey()
	TestPubkey1  = TestPrivKey1.PubKey()
	TestAddress1 = sdk.AccAddress(TestPubkey1.Address())
)

package store

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Get(ctx sdk.Context, sk sdk.StoreKey, cdc *codec.Codec, key []byte, proto interface{}) error {
	store := ctx.KVStore(sk)

	bz := store.Get(key)
	if bz == nil {
		return fmt.Errorf(" not found")
	}

	cdc.MustUnmarshalBinaryLengthPrefixed(bz, proto)
	return nil
}

func Set(ctx sdk.Context, sk sdk.StoreKey, cdc *codec.Codec, key []byte, val interface{}) {
	store := ctx.KVStore(sk)
	store.Set(key, cdc.MustMarshalBinaryLengthPrefixed(val))
}

func SetNotExists(ctx sdk.Context, sk sdk.StoreKey, cdc *codec.Codec, key []byte, val interface{}) error {
	if Has(ctx, sk, key) {
		return fmt.Errorf("already exist")
	}
	Set(ctx, sk, cdc, key, val)
	return nil
}

func SetExists(ctx sdk.Context, sk sdk.StoreKey, cdc *codec.Codec, key []byte, val interface{}) error {
	if !Has(ctx, sk, key) {
		return fmt.Errorf("not found")
	}
	Set(ctx, sk, cdc, key, val)
	return nil
}

func Has(ctx sdk.Context, sk sdk.StoreKey, key []byte) bool {
	store := ctx.KVStore(sk)
	return store.Has(key)
}

func Iterator(ctx sdk.Context, sk sdk.StoreKey, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(sk)
	iterator := sdk.KVStorePrefixIterator(store, prefix)

	return iterator
}

func Delete(ctx sdk.Context, sk sdk.StoreKey, key []byte) error {
	if !Has(ctx, sk, key) {
		return fmt.Errorf("not found")
	}

	store := ctx.KVStore(sk)
	store.Delete(key)
	return nil
}

package types

type GenesisState struct {
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func (g GenesisState) ValidateGenesis() error {
	return nil
}

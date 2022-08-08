package tokenregistry

import (
	"os"

	"github.com/teal-finance/solana-go"
)

var (
	mainnetProgramID = solana.MustPublicKeyFromBase58("CmPVzy88JSB4S223yCvFmBxTLobLya27KgEDeNPnqEub")
	testnetProgramID = solana.MustPublicKeyFromBase58("99999999999999999999999999999999999999999999")
	devnetProgramID  = solana.MustPublicKeyFromBase58("99999999999999999999999999999999999999999999")
)

func ProgramID() solana.PublicKey {
	if custom := os.Getenv("TOKEN_REGISTRY_PROGRAM_ID"); custom != "" {
		return solana.MustPublicKeyFromBase58(custom)
	}

	network := os.Getenv("SOL_NETWORK")

	switch network {
	case "mainnet":
		return mainnetProgramID
	case "testnet":
		return testnetProgramID
	case "devnet":
		return devnetProgramID
	default:
		return mainnetProgramID
	}
}

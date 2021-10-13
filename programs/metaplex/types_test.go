package metaplex

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/streamingfast/solana-go"
	"github.com/stretchr/testify/require"
)

func TestD(t *testing.T) {
	t.Skip("done")
	data, err := hex.DecodeString("04eb71d2f68370da513ad471ed6bb975a331931cb4e221aead8247ab8e4618d04fe96feb308986580b09abdc3182417ec643b338405dba3ec2734c6ef6cef22cf7200000005061727469616c206669677572650000000000000000000000000000000000000a00000000000000000000000000c800000068747470733a2f2f73756d6974726565732e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f476945795a355151384b353733535637664b584b6f6d415231537045434236465469485870515147387774692e6a736f6e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000101fd0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)

	m := &Metadata{}
	err = m.Decode(data)
	require.NoError(t, err)

	assert.Equal(t, &Metadata{
		Key:             MetadataV1,
		UpdateAuthority: solana.MustPublicKeyFromBase58("Gr5UanqwiKA54GGnw4b1bB5M8eatQzj6s6FQ9FeTze5C"),
		Mint:            solana.MustPublicKeyFromBase58("GiEyZ5QQ8K573SV7fKXKomAR1SpECB6FTiHXpQQG8wti"),
		Data: Data{
			Name:                 "Partial figure",
			Symbol:               "",
			URI:                  "https://sumitrees.s3.ap-northeast-1.amazonaws.com/GiEyZ5QQ8K573SV7fKXKomAR1SpECB6FTiHXpQQG8wti.json",
			SellerFeeBasisPoints: 0,
			Creators:             &[]Creator{},
		},
		PrimarySaleHappened: false,
		IsMutable:           true,
	}, m)

}

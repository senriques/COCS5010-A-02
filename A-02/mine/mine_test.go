package mine

import (
	"encoding/hex"
	"testing"

	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/block"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/hash"
)

func TestMineBlock(t *testing.T) {

	tests := []struct {
		bk               block.BlockType
		expectedSealHash string
		expectedNonce    uint64
	}{
		{
			bk: block.BlockType{
				Index:         0,
				Desc:          block.GenesisDesc,
				ThisBlockHash: []byte{},
				PrevBlockHash: []byte{},
				Nonce:         0,
				Seal:          []byte{},
			},
			expectedSealHash: "0000ae2cab130b4836988969f731c4f884ac4675790e5575a5161e5b96ab13d7",
			expectedNonce:    54586,
		},
		{
			bk: block.BlockType{
				Index:         1,
				Desc:          "My First Block",
				ThisBlockHash: []byte{},
				PrevBlockHash: []byte{},
				Nonce:         0,
				Seal:          []byte{},
			},
			expectedSealHash: "0000adc29a80f1f0df08c8687c013d179050f5d1b449599e4d1437e4fad23525",
			expectedNonce:    46734,
		},
		{
			bk: block.BlockType{
				Index:         1,
				Desc:          "My First Block",
				ThisBlockHash: []byte{},
				PrevBlockHash: MustDecodeString("136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e"),
				Nonce:         0,
				Seal:          []byte{},
			},
			expectedSealHash: "000013ce557332aaa68abe3b7bf1be856743a03689a802606a732e81713bb78c",
			expectedNonce:    4527,
		},
	}

	for ii, test := range tests {
		test.bk.ThisBlockHash = hash.HashOf(block.SerializeBlock(&test.bk))
		tests[ii] = test
	}

	for ii, test := range tests {
		bk := &test.bk
		MineBlock(bk)
		SealString := hex.EncodeToString(bk.Seal)
		if SealString != test.expectedSealHash {
			t.Errorf("Test %d, expected %s got %s\n", ii, test.expectedSealHash, SealString)
		}
		if bk.Nonce != test.expectedNonce {
			t.Errorf("Test %d, expected %d got %d\n", ii, test.expectedNonce, bk.Nonce)
		}
	}

}

func MustDecodeString(s string) []byte {
	rv, err := hex.DecodeString("136c53391115ab7ff717bd24e62dd0df2c270500d7194290169a83488022548e")
	if err != nil {
		panic(err)
	}
	return rv
}

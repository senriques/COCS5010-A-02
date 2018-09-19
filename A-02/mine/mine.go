package mine

import (
	"encoding/hex"
	"fmt"

	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/block"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/hash"
)

// MineBlock implements a proof of work mining system where the first 4 digits
//(2 bytes) of the hex value are 0. Difficulty can be increaesed by requiring
//more 0's or by requring some other pattern to apear in the resulting hash.
func MineBlock(bk *block.BlockType) (*block.BlockType){

	for {
	//serialize the current block -> hash it -> change to string, look for work criteria
		OutFromSerialize := block.SerializeForSeal(bk)
		HashOfHash := hash.HashOf(OutFromSerialize)
		theHashAsAString := hex.EncodeToString(HashOfHash)
		fmt.Printf("((Mining)) Hash for Block [%s] nonce [%8d]\r", theHashAsAString, bk.Nonce)
		// For the genesis block it requires 54586 iterations to calculate the proof of work.
		if (theHashAsAString[0:4] == "0000")	{
			fmt.Printf("\n*Work Criteria Met*\n")
			bk.Seal = HashOfHash
			return(bk)						// return the block back to originating function for write, etc.
			}	else	{
				bk.Nonce++
		}
	}
}

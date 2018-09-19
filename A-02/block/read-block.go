package block

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
//"encoding/hex"

	"github.com/pschlump/godebug"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/hash"

)

type IndexType struct {
	Index		      []string		 // Name of the genesis block file
	TxHashIndex	  TxHashIndexType //
}
type TxHashIndexType struct {
	TxHashIndex		hash.BlockHashType
	Index					uint64
	BlockHash			hash.BlockHashType
}

// ReadBlock reads in from a file a block and returns it.
func ReadBlock(fn string) (bk *BlockType, err error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Printf("Unable to read genesis block %s [at %s], %s\n", fn, godebug.LF(), err)
		return nil, err
	}
	bk = &BlockType{}
	err = json.Unmarshal(data, bk)
	if !IsGenesisBlock(bk) {
		fmt.Printf("Unable to read genesis block %s [at %s], %s\n", fn, godebug.LF(), err)
		return nil, err
	}
	return bk, nil
}

// WriteBlock read i;n a block from a file.
func WriteBlock(fn string, bk *BlockType) (err error) {
	data := IndentJSON(bk)
	err = ioutil.WriteFile(fn, []byte(data), 0644)
	return
}

// IndentJSON converts data to JSON format and returns it as a string.
func IndentJSON(v interface{}) string {
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	} else {
		return string(s)
	}
}

func ReadIndex(fn string) (bk *IndexType, err error) {

	// read index data from fn filepath, unmarshal .json input and return bk struct
	indexdata, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Printf("Unable to read genesis block %s [at %s], %s\n", fn, godebug.LF(), err)
		return nil, err
	}
	bk = &IndexType{}
	err = json.Unmarshal(indexdata, bk)
	return bk, nil
}

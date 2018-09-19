// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 2A Mining Assignment

package main

import (
	"flag"
	"fmt"
	"os"
//	"encoding/hex"

	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/cli"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/config"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/block"
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/mine"
)

type BasePath	struct 	{
	DataDir				string
}

func main() {

	var CreateGenesisFlag  = flag.Bool("create-genesis", false, "init command")
	var TestReadBlockFlag  = flag.Bool("test-read-block", false, "test read a block")
	var TestWriteBlockFlag = flag.Bool("test-write-block", false, "test write a block")
	var MineFlag = flag.Bool("Mine", false, "Mine")

	PPath := "C:/Go/src/github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-02/main/data/"

	flag.Parse() 					// Parse gpt create genesis block, read or write block
	fns := flag.Args()
	Cfg := flag.String("cfg", PPath + fns[0], "config file for this call")
	if Cfg == nil {
		fmt.Printf("configuration (CG, RB, WB), -*.json filepath, \n")
		os.Exit(1)
	}

// read from file path to store block(s)
	err := config.ReadConfig(*Cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read configuration: %s error %s\n", *Cfg, err)
		os.Exit(1)
	}
	gCfg := config.GetGlobalConfig()
	cc := cli.NewCLI(gCfg)

// Switch for input type to set flags: create, read or write blocks
switch fns[1]	{
	case "create-genesis":
		*CreateGenesisFlag = true
	case "CG":													// short-form for create-genesis
		*CreateGenesisFlag = true
	case "test-read-block":
		*TestReadBlockFlag = true
	case "TR":												// short-form for test-read-block
		*TestReadBlockFlag = true
	case "test-write-block":
		*TestWriteBlockFlag = true
	case "TW":												// short-form for test-write-block
		*TestWriteBlockFlag = true
	case "mine":
		*MineFlag = true
	default:		{
		fmt.Printf("--create-genesis, test-read-block, test-write-block\n")
		os.Exit(1)
		cc.Usage()
		}
	}

// run functions create, test read and test write based on Flags
	if *CreateGenesisFlag {										// Create a Genesis Block
		fmt.Printf("-cc.CreateGenesis\n")
		cc.CreateGenesis()
		}
	if *TestReadBlockFlag {										// Test ReadBlock function
		fmt.Printf("-cc.TestReadBlock: ")
		cc.TestReadBlock()
			fmt.Printf("\n")
		}
	if *TestWriteBlockFlag {									// Test WriteBlock function
		fmt.Printf("-cc.TestWriteBlock: ")
		cc.TestReadBlock()
		fmt.Printf("\n")
		}
	if *MineFlag {														// Mine for seal
		fmt.Printf("-cc.Mine\n")
		// read index.json to find block to change select the last one in the slice
		genblock, err := block.ReadIndex(*Cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read configuration: %s error %s\n", *Cfg, err)
			os.Exit(1)
			}
		// concatenate the filepath to open and read the corresponding block data
		PPath = PPath + genblock.Index[len(genblock.Index)-1] + ".json"
		BlockIn, err := block.ReadBlock(PPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in ReadBlock: %s error %s\n", *Cfg, err)
			os.Exit(1)
			}
		// mine for seal from loaded block, find it, and rewrite block w/ seal & Nonce
		SealedBlock := mine.MineBlock(BlockIn)
		block.WriteBlock(PPath, SealedBlock)
		}
}																					// end main

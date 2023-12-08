package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	files "github.com/ipfs/boxo/files"
	"github.com/ofman/filesharego"
	// This package is needed so that all the preloaded plugins are loaded automatically
)

func main() {

	var flagFilePath string
	flag.StringVar(&flagFilePath, "f", "", "a string path var") // filepath cli flag set

	var flagCid string
	flag.StringVar(&flagCid, "c", "", "a string cid var") // cid cli flag set

	flag.Parse()

	if flagCid != "" || flagFilePath != "" {
		fmt.Println("-- Getting an IPFS node running -- ")

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Spawn a node using a temporary path, creating a temporary repo for the run
		fmt.Println("Spawning Kubo node on a temporary repo")
		ipfsB, _, err := filesharego.SpawnEphemeral(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
		}

		fmt.Println("IPFS node is running")

		if flagCid != "" {
			filesharego.DownloadFromCid(ctx, ipfsB, flagCid)
		} else if flagFilePath != "" {
			someFile, err := filesharego.GetUnixfsNode(flagFilePath)
			if err != nil {
				panic(fmt.Errorf("could not get File: %s", err))
			}

			//for the future simplicity to download single files in the same directory. Opened ticked on ipfs here: https://github.com/ipfs/boxo/issues/520
			fileInfo, err := os.Stat(flagFilePath)
			if err != nil {
				panic(fmt.Errorf("could not get File stat info: %s", err))
			}

			// wrap file into directory with filename so ipfs shows file name later
			if !fileInfo.IsDir() {
				someFile = files.NewSliceDirectory([]files.DirEntry{
					files.FileEntry(filepath.Base(flagFilePath), someFile),
				})
			}
			filesharego.UploadFiles(ctx, ipfsB, someFile)
		}
	} else {
		fmt.Println("Use flags -f \"example.jpg\" or -c \"exampleCid\" to share files for example:\n./fsg -f \"example.jpg\"\nor to download files\n./fsg -c \"exampleCid\"")
	}
}

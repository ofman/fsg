package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/ofman/filesharego"
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
			filesharego.UploadFromPath(ctx, ipfsB, flagFilePath)
		}
	} else {
		fmt.Println("Use flags -f \"example.jpg\" or -c \"exampleCid\" to share files for example:\n./fsg -f \"example.jpg\"\nor to download files\n./fsg -c \"exampleCid\"")
	}
}

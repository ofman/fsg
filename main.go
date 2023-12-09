package main

import (
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
		if flagCid != "" {
			filesharego.DownloadFromCid(flagCid)
		} else if flagFilePath != "" {
			filesharego.UploadFiles(flagFilePath, false)
		}
	} else {
		fmt.Println("Use flags -f \"example.jpg\" or -c \"exampleCid\" to share files for example:\n./fsg -f \"example.jpg\"\nor to download files\n./fsg -c \"exampleCid\"")
	}
}

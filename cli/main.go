package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Jesse-Lucas1996/craftforge/cli/downloader"
)

func main() {
	forgeVersion := flag.String("forge", "", "Forge version to install")
	modURL := flag.String("mod", "", "URL of mod to download")
	modName := flag.String("modname", "mod", "Name of mod file")

	flag.Parse()

	if *forgeVersion == "" {
		fmt.Println("Error: Forge version is required.")
		os.Exit(1)
	}

	if err := downloader.DownloadForgeInstaller(*forgeVersion); err != nil {
		fmt.Printf("Error downloading Forge installer: %v\n", err)
		os.Exit(1)
	}

	if *modURL != "" {
		if err := downloader.DownloadMod(*modURL, *modName); err != nil {
			fmt.Printf("Error downloading mod: %v\n", err)
			os.Exit(1)
		}
	}

}

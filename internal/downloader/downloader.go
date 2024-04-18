// internal/downloader/downloader.go
package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadForgeInstaller(version string) error {
	url := fmt.Sprintf("https://files.minecraftforge.net/maven/net/minecraftforge/forge/%s/forge-%s-installer.jar", version, version)

	err := downloadFile(url, "forge-installer.jar")
	if err != nil {
		return fmt.Errorf("failed to download Forge installer: %v", err)
	}
	return nil
}

func DownloadMod(modURL string, modName string) error {
	err := downloadFile(modURL, modName+".jar")
	if err != nil {
		return fmt.Errorf("failed to download mod: %v", err)
	}
	return nil
}

func downloadFile(url string, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

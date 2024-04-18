package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func FindBuildVersion(version string) (string, string, error) {
	url := fmt.Sprintf("https://files.minecraftforge.net/net/minecraftforge/forge/index_%s.html", version)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch Forge builds: %v", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse HTML: %v", err)
	}

	var buildVersion string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "small" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					buildVersion = strings.TrimSpace(c.Data)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	parts := strings.Split(buildVersion, "-")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid build version format")
	}

	version = strings.TrimSpace(parts[0])
	build := strings.TrimSpace(parts[1])

	fmt.Printf("Found version: %s, build: %s\n", version, build)

	return version, build, nil
}

func DownloadForgeInstaller(version string) error {
	version, build, err := FindBuildVersion(version)
	if err != nil {
		return fmt.Errorf("failed to find Forge build version: %v", err)
	}

	url := fmt.Sprintf("https://maven.minecraftforge.net/net/minecraftforge/forge/%s-%s/forge-%s-%s-installer.jar", version, build, version, build)
	err = downloadFile(url, "forge-installer.jar")
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
	if response.StatusCode != 200 {
		return fmt.Errorf("failed to download file: %v", response.Status)

	}
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

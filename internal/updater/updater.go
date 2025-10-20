package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	githubRepo = "xqsit94/glm"
	apiURL     = "https://api.github.com/repos/" + githubRepo + "/releases/latest"
)

type ReleaseInfo struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Body    string `json:"body"`
	HTMLURL string `json:"html_url"`
}

type UpdateInfo struct {
	CurrentVersion string
	LatestVersion  string
	HasUpdate      bool
	ReleaseNotes   string
	ReleaseURL     string
}

func GetLatestVersion() (*ReleaseInfo, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch release info: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var release ReleaseInfo
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("failed to parse release info: %v", err)
	}

	return &release, nil
}

func CompareVersions(current, latest string) int {
	current = strings.TrimPrefix(current, "v")
	latest = strings.TrimPrefix(latest, "v")

	currentParts := strings.Split(current, ".")
	latestParts := strings.Split(latest, ".")

	maxLen := len(currentParts)
	if len(latestParts) > maxLen {
		maxLen = len(latestParts)
	}

	for i := 0; i < maxLen; i++ {
		var currentVal, latestVal int

		if i < len(currentParts) {
			currentVal, _ = strconv.Atoi(currentParts[i])
		}
		if i < len(latestParts) {
			latestVal, _ = strconv.Atoi(latestParts[i])
		}

		if latestVal > currentVal {
			return 1
		} else if latestVal < currentVal {
			return -1
		}
	}

	return 0
}

func DetectPlatform() (string, string, error) {
	osName := runtime.GOOS
	arch := runtime.GOARCH

	if osName != "darwin" && osName != "linux" {
		return "", "", fmt.Errorf("unsupported operating system: %s", osName)
	}

	if arch != "amd64" && arch != "arm64" {
		return "", "", fmt.Errorf("unsupported architecture: %s", arch)
	}

	return osName, arch, nil
}

func DownloadBinary(version, osName, arch string, progressCallback func(downloaded, total int64)) (string, error) {
	binaryName := fmt.Sprintf("glm-%s-%s", osName, arch)
	downloadURL := fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", githubRepo, version, binaryName)

	resp, err := http.Get(downloadURL)
	if err != nil {
		return "", fmt.Errorf("failed to download binary: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("download failed with status %d", resp.StatusCode)
	}

	tmpFile, err := os.CreateTemp("", "glm-update-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}
	defer tmpFile.Close()

	var downloaded int64
	total := resp.ContentLength

	buf := make([]byte, 32*1024)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			if _, writeErr := tmpFile.Write(buf[:n]); writeErr != nil {
				os.Remove(tmpFile.Name())
				return "", fmt.Errorf("failed to write to temp file: %v", writeErr)
			}
			downloaded += int64(n)
			if progressCallback != nil {
				progressCallback(downloaded, total)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Remove(tmpFile.Name())
			return "", fmt.Errorf("failed to download binary: %v", err)
		}
	}

	return tmpFile.Name(), nil
}

func VerifyBinary(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to stat binary: %v", err)
	}

	if info.Size() == 0 {
		return fmt.Errorf("downloaded binary is empty")
	}

	if err := os.Chmod(path, 0755); err != nil {
		return fmt.Errorf("failed to make binary executable: %v", err)
	}

	return nil
}

func InstallUpdate(newBinaryPath string) error {
	currentBinary, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get current binary path: %v", err)
	}

	currentBinary, err = filepath.EvalSymlinks(currentBinary)
	if err != nil {
		return fmt.Errorf("failed to resolve binary path: %v", err)
	}

	backupPath := currentBinary + ".old"
	if err := os.Rename(currentBinary, backupPath); err != nil {
		return fmt.Errorf("failed to backup current binary: %v", err)
	}

	if err := os.Rename(newBinaryPath, currentBinary); err != nil {
		os.Rename(backupPath, currentBinary)
		return fmt.Errorf("failed to install new binary: %v", err)
	}

	os.Remove(backupPath)

	return nil
}

func CheckForUpdate(currentVersion string) (*UpdateInfo, error) {
	release, err := GetLatestVersion()
	if err != nil {
		return nil, err
	}

	info := &UpdateInfo{
		CurrentVersion: currentVersion,
		LatestVersion:  release.TagName,
		ReleaseNotes:   release.Body,
		ReleaseURL:     release.HTMLURL,
	}

	comparison := CompareVersions(currentVersion, release.TagName)
	info.HasUpdate = comparison > 0

	return info, nil
}

func FormatReleaseNotes(notes string, maxLines int) string {
	lines := strings.Split(notes, "\n")
	if len(lines) <= maxLines {
		return notes
	}

	result := strings.Join(lines[:maxLines], "\n")
	return result + "\n..."
}

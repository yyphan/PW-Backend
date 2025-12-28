package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const MdRootFolderKey = "LOCAL_MD_ROOT"

// e.g. en/mySeries/part-1.md
func GetMarkdownRelaPath(lang, seriesSlug, postSlug string) string {
	return filepath.Join(lang, seriesSlug, fmt.Sprintf("%s.md", postSlug))
}

func WriteMarkdown(relativePath, content string) error {
	baseStoragePath := os.Getenv(MdRootFolderKey)
	fullPath := filepath.Join(baseStoragePath, relativePath)

	return writeFile(fullPath, content)
}

func ReadMarkdown(relativePath string) (string, error) {
	baseStoragePath := os.Getenv(MdRootFolderKey)
	fullPath := filepath.Join(baseStoragePath, relativePath)

	return readFile(fullPath)
}

func writeFile(absPath, content string) error {
	dir := filepath.Dir(absPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return os.WriteFile(absPath, []byte(content), 0644)
}

func readFile(absPath string) (string, error) {
	content, err := os.ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}

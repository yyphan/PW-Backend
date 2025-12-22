package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

const MdRootFolderKey = "MD_ROOT"

// e.g. en/mySeries/part-1.md
func GetMarkdownRelaPath(lang, seriesSlug, postSlug string) string {
	return filepath.Join(lang, seriesSlug, fmt.Sprintf("%s.md", postSlug))
}

func WriteFile(relativePath, content string) error {
	baseStoragePath := os.Getenv(MdRootFolderKey)
	fullPath := filepath.Join(baseStoragePath, relativePath)

	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

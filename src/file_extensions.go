package logging

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func buildFileName(parts ...string) string {
	builder := NewStringBuilder()
	for _, item := range parts {
		builder.Append(item)
	}
	return builder.String()
}
func getFileExtension(file string) string {
	return filepath.Ext(file)
}
func getFileName(path string) string {
	_, file := filepath.Split(path)
	return file
}
func getFileSize(file string) (int64, error) {
	info, err := os.Stat(file)

	return info.Size(), err
}
func getFileNameWithoutExtension(path string) string {
	_, file := filepath.Split(path)
	extension := getFileExtension(file)
	return strings.TrimSuffix(file, extension)
}
func getFileNameAndExtension(path string) (string, string, string) {
	dir, file := filepath.Split(path)
	extension := getFileExtension(file)
	return dir, strings.TrimSuffix(file, extension), extension
}

func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
func deleteFile(fileName string) bool {
	err := os.Remove(fileName)
	if err != nil {
		return false
	}

	return true
}
func deleteIfExists(fileName string) bool {
	if fileExists(fileName) {
		return deleteFile(fileName)
	}

	return false
}

func generateFileName(fullName string) string {
	dir, name, extension := getFileNameAndExtension(fullName)
	index := 1
	newName := filepath.Join(dir, buildFileName(name, "_", strconv.Itoa(index), extension))

	for fileExists(newName) {
		newName = filepath.Join(dir, buildFileName(name, "_", strconv.Itoa(index), extension))
		index++
	}

	return newName
}

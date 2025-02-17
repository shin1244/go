package fileSystem

import (
	"errors"
	"os"
)

type FileSystem interface {
	ReadFile(name string) (string, error)        // 이름을 받아서 파일 내용을 반환
	WriteFile(name string, content string) error // 이름과 내용을 받아서 파일 쓰기
	DeleteFile(name string) error                // 이름을 받아서 파일 삭제
}

type LocalFileSystem struct{}

func (l *LocalFileSystem) ReadFile(name string) (string, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (l *LocalFileSystem) WriteFile(name string, content string) error {
	return os.WriteFile(name, []byte(content), 0664)
}

func (l *LocalFileSystem) DeleteFile(name string) error {
	return os.Remove(name)
}

type MockFileSystem struct {
	files map[string]string
}

func NewMockFileSystem() *MockFileSystem {
	return &MockFileSystem{files: make(map[string]string)}
}

func (m *MockFileSystem) ReadFile(name string) (string, error) {
	content := m.files[name]
	if len(content) <= 0 {
		return "", errors.New("file not found")
	}
	return content, nil
}

func (m *MockFileSystem) WriteFile(name string, content string) error {
	m.files[name] = content
	return nil
}

func (m *MockFileSystem) DeleteFile(name string) error {
	delete(m.files, name)
	return nil
}

func ReadFile(fs FileSystem, name string) (string, error) {
	return fs.ReadFile(name)
}

func WriteFile(fs FileSystem, name string, content string) error {
	return fs.WriteFile(name, content)
}

func DeleteFile(fs FileSystem, name string) error {
	return fs.DeleteFile(name)
}

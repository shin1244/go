package fileSystem

import (
	"testing"
)

func TestLocalFileSystem(t *testing.T) {
	fs := LocalFileSystem{}

	content := "Hello, World!"
	err := fs.WriteFile("test.txt", content)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	readContent, err := fs.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if readContent != content {
		t.Fatalf("Content mismatch: %s != %s", readContent, content)
	}

	err = fs.DeleteFile("test.txt")
	if err != nil {
		t.Fatalf("Failed to delete file: %v", err)
	}

	_, err = fs.ReadFile("test.txt")
	if err == nil {
		t.Fatalf("File should be deleted")
	}
}

func TestMockFileSystem(t *testing.T) {
	fs := NewMockFileSystem()

	content := "Hello, World!"

	err := fs.WriteFile("test.txt", content)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	readContent, err := fs.ReadFile("test.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if readContent != content {
		t.Fatalf("Content mismatch: %s != %s", readContent, content)
	}

	err = fs.DeleteFile("test.txt")
	if err != nil {
		t.Fatalf("Failed to delete file: %v", err)
	}

	_, err = fs.ReadFile("test.txt")
	if err == nil {
		t.Fatalf("File should be deleted")
	}

}

func TestInterface(t *testing.T) {
	fs := &LocalFileSystem{}
	mfs := NewMockFileSystem()

	content := "Hello, World!"

	WriteFile(fs, "test.txt", content)
	WriteFile(mfs, "test.txt", content)

	readContent, err := ReadFile(fs, "test.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if readContent != content {
		t.Fatalf("Content mismatch: %s != %s", readContent, content)
	}

	err = DeleteFile(fs, "test.txt")
	if err != nil {
		t.Fatalf("Failed to delete file: %v", err)
	}

	readContent, err = ReadFile(mfs, "test.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if readContent != content {
		t.Fatalf("Content mismatch: %s != %s", readContent, content)
	}

	err = DeleteFile(mfs, "test.txt")
	if err != nil {
		t.Fatalf("Failed to delete file: %v", err)
	}

	_, err = ReadFile(mfs, "test.txt")
	if err == nil {
		t.Fatalf("File should be deleted")
	}

}

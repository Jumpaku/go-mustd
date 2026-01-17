package osmust_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpaku/go-mustd/osmust"
)

func TestReadWriteFile(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.txt")

	t.Run("WriteFile and ReadFile", func(t *testing.T) {
		data := []byte("hello world")
		osmust.WriteFile(filename, data, 0644)

		readData := osmust.ReadFile(filename)
		if string(readData) != "hello world" {
			t.Errorf("expected 'hello world', got %s", readData)
		}
	})
}

func TestCreateOpenFile(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.txt")

	t.Run("Create and Write", func(t *testing.T) {
		f := osmust.Create(filename)
		defer f.Close()

		n := f.WriteString("hello")
		if n != 5 {
			t.Errorf("expected 5 bytes written, got %d", n)
		}
	})

	t.Run("Open and Read", func(t *testing.T) {
		f := osmust.Open(filename)
		defer f.Close()

		buf := make([]byte, 5)
		n := f.Read(buf)
		if n != 5 {
			t.Errorf("expected 5 bytes read, got %d", n)
		}
		if string(buf) != "hello" {
			t.Errorf("expected 'hello', got %s", buf)
		}
	})
}

func TestMkdirRemove(t *testing.T) {
	tmpDir := t.TempDir()
	dirPath := filepath.Join(tmpDir, "testdir")

	t.Run("Mkdir and Remove", func(t *testing.T) {
		osmust.Mkdir(dirPath, 0755)

		info := osmust.Stat(dirPath)
		if !info.IsDir() {
			t.Error("expected directory")
		}

		osmust.Remove(dirPath)
		_, err := os.Stat(dirPath)
		if !os.IsNotExist(err) {
			t.Error("directory should not exist after Remove")
		}
	})
}

func TestMkdirAll(t *testing.T) {
	tmpDir := t.TempDir()
	dirPath := filepath.Join(tmpDir, "a", "b", "c")

	osmust.MkdirAll(dirPath, 0755)

	info := osmust.Stat(dirPath)
	if !info.IsDir() {
		t.Error("expected directory")
	}
}

func TestRename(t *testing.T) {
	tmpDir := t.TempDir()
	oldPath := filepath.Join(tmpDir, "old.txt")
	newPath := filepath.Join(tmpDir, "new.txt")

	osmust.WriteFile(oldPath, []byte("test"), 0644)
	osmust.Rename(oldPath, newPath)

	data := osmust.ReadFile(newPath)
	if string(data) != "test" {
		t.Errorf("expected 'test', got %s", data)
	}

	_, err := os.Stat(oldPath)
	if !os.IsNotExist(err) {
		t.Error("old file should not exist after Rename")
	}
}

func TestGetwd(t *testing.T) {
	dir := osmust.Getwd()
	if dir == "" {
		t.Error("Getwd returned empty string")
	}
}

func TestHostname(t *testing.T) {
	hostname := osmust.Hostname()
	if hostname == "" {
		t.Error("Hostname returned empty string")
	}
}

func TestUserHomeDir(t *testing.T) {
	home := osmust.UserHomeDir()
	if home == "" {
		t.Error("UserHomeDir returned empty string")
	}
}

func TestSetenvUnsetenv(t *testing.T) {
	key := "GO_MUSTD_TEST_VAR"
	value := "test_value"

	osmust.Setenv(key, value)
	if os.Getenv(key) != value {
		t.Errorf("expected %s, got %s", value, os.Getenv(key))
	}

	osmust.Unsetenv(key)
	if os.Getenv(key) != "" {
		t.Error("environment variable should be empty after Unsetenv")
	}
}

func TestChmod(t *testing.T) {
	tmpDir := t.TempDir()
	filename := filepath.Join(tmpDir, "test.txt")

	osmust.WriteFile(filename, []byte("test"), 0644)
	osmust.Chmod(filename, 0600)

	info := osmust.Stat(filename)
	if info.Mode().Perm() != 0600 {
		t.Errorf("expected permission 0600, got %o", info.Mode().Perm())
	}
}

func TestCreateTemp(t *testing.T) {
	tmpDir := t.TempDir()
	f := osmust.CreateTemp(tmpDir, "test-*")
	defer f.Close()
	defer osmust.Remove(f.Name())

	n := f.WriteString("temp content")
	if n != 12 {
		t.Errorf("expected 12 bytes written, got %d", n)
	}
}

func TestMkdirTemp(t *testing.T) {
	tmpDir := t.TempDir()
	dir := osmust.MkdirTemp(tmpDir, "test-*")
	defer osmust.RemoveAll(dir)

	info := osmust.Stat(dir)
	if !info.IsDir() {
		t.Error("expected directory")
	}
}

func TestSymlink(t *testing.T) {
	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "target.txt")
	link := filepath.Join(tmpDir, "link.txt")

	osmust.WriteFile(target, []byte("target"), 0644)
	osmust.Symlink(target, link)

	linkTarget := osmust.Readlink(link)
	if linkTarget != target {
		t.Errorf("expected %s, got %s", target, linkTarget)
	}

	info := osmust.Lstat(link)
	if info.Mode()&os.ModeSymlink == 0 {
		t.Error("expected symbolic link")
	}
}

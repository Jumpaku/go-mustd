package filepathmust_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpaku/go-mustd/pathmust/filepathmust"
)

func TestAbs(t *testing.T) {
	t.Run("relative path", func(t *testing.T) {
		abs := filepathmust.Abs(".")
		if abs == "" {
			t.Error("Abs returned empty string")
		}
		if !filepath.IsAbs(abs) {
			t.Errorf("expected absolute path, got %s", abs)
		}
	})
}

func TestRel(t *testing.T) {
	t.Run("valid paths", func(t *testing.T) {
		rel := filepathmust.Rel("/a/b", "/a/b/c/d")
		if rel != "c/d" && rel != filepath.Join("c", "d") {
			t.Errorf("expected 'c/d', got %s", rel)
		}
	})
}

func TestGlob(t *testing.T) {
	tmpDir := t.TempDir()

	// Create test files
	os.WriteFile(filepath.Join(tmpDir, "test1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "test2.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "other.dat"), []byte("test"), 0644)

	t.Run("glob pattern", func(t *testing.T) {
		pattern := filepath.Join(tmpDir, "*.txt")
		matches := filepathmust.Glob(pattern)
		if len(matches) != 2 {
			t.Errorf("expected 2 matches, got %d", len(matches))
		}
	})

	t.Run("invalid pattern panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Glob did not panic with invalid pattern")
			}
		}()
		filepathmust.Glob("[]")
	})
}

func TestMatch(t *testing.T) {
	t.Run("matching pattern", func(t *testing.T) {
		matched := filepathmust.Match("*.txt", "test.txt")
		if !matched {
			t.Error("expected match")
		}
	})

	t.Run("non-matching pattern", func(t *testing.T) {
		matched := filepathmust.Match("*.txt", "test.dat")
		if matched {
			t.Error("expected no match")
		}
	})

	t.Run("invalid pattern panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Match did not panic with invalid pattern")
			}
		}()
		filepathmust.Match("[]", "test.txt")
	})
}

func TestEvalSymlinks(t *testing.T) {
	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "target.txt")
	link := filepath.Join(tmpDir, "link.txt")

	os.WriteFile(target, []byte("test"), 0644)
	os.Symlink(target, link)

	t.Run("resolve symlink", func(t *testing.T) {
		resolved := filepathmust.EvalSymlinks(link)
		// On macOS, tmpDir might be under /var which is symlinked to /private/var
		// So we need to resolve the expected target as well
		expectedTarget := filepathmust.EvalSymlinks(target)
		if resolved != expectedTarget {
			t.Errorf("expected %s, got %s", expectedTarget, resolved)
		}
	})
}

func TestWalk(t *testing.T) {
	tmpDir := t.TempDir()

	// Create directory structure
	os.MkdirAll(filepath.Join(tmpDir, "a", "b"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "file.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "a", "file2.txt"), []byte("test"), 0644)

	t.Run("walk directory", func(t *testing.T) {
		count := 0
		filepathmust.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			count++
			return nil
		})
		if count < 4 {
			t.Errorf("expected at least 4 items, got %d", count)
		}
	})
}

func TestWalkDir(t *testing.T) {
	tmpDir := t.TempDir()

	// Create directory structure
	os.MkdirAll(filepath.Join(tmpDir, "a", "b"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "file.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "a", "file2.txt"), []byte("test"), 0644)

	t.Run("walk directory", func(t *testing.T) {
		count := 0
		filepathmust.WalkDir(tmpDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			count++
			return nil
		})
		if count < 4 {
			t.Errorf("expected at least 4 items, got %d", count)
		}
	})
}

func TestLocalize(t *testing.T) {
	t.Run("localize path", func(t *testing.T) {
		// This test is platform-dependent
		// Just verify it doesn't panic
		result := filepathmust.Localize("a/b/c")
		if result == "" {
			t.Error("Localize returned empty string")
		}
	})
}

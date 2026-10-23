package util

import (
	"os"
	"path/filepath"
	"testing"
)

// Cases from stow/t/join_paths.t
func TestJoinPaths(t *testing.T) {
	tests := []struct {
		in       []string
		want     string
		scenario string
	}{
		{[]string{"a/b/c", "d/e/f"}, "a/b/c/d/e/f", "simple"},
		{[]string{"a/b/c", "/d/e/f"}, "/d/e/f", "relative then absolute"},
		{[]string{"/a/b/c", "d/e/f"}, "/a/b/c/d/e/f", "absolute then relative"},
		{[]string{"/a/b/c", "/d/e/f"}, "/d/e/f", "two absolutes"},
		{[]string{"/a/b/c/", "/d/e/f/"}, "/d/e/f", "two absolutes with trailing /"},
		{[]string{"///a/b///c//", "/d///////e/f"}, "/d/e/f", "multiple /'s, absolute"},
		{[]string{"///a/b///c//", "d///////e/f"}, "/a/b/c/d/e/f", "multiple /'s, relative"},
		{[]string{"", "a/b/c"}, "a/b/c", "first empty"},
		{[]string{"a/b/c", ""}, "a/b/c", "second empty"},
		{[]string{"/", "a/b/c"}, "/a/b/c", "first is /"},
		{[]string{"a/b/c", "/"}, "/", "second is /"},
		{[]string{"../a1/b1/../c1/", "a2/../b2/e2"}, "../a1/c1/b2/e2", "relative with ../"},
		{[]string{"../a1/b1/../c1/", "/a2/../b2/e2"}, "/b2/e2", "absolute with ../"},
		{[]string{"../a1/../../c1", "a2/../../"}, "../..", "lots of ../"},
		{[]string{"./", "../a2"}, "../a2", "drop any \"./\""},
		{[]string{"./a1", "../../a2"}, "../a2", "drop any \"./foo\""},
		{[]string{"a/b/c", "."}, "a/b/c", ". on RHS"},
		{[]string{"a/b/c", ".", "d/e"}, "a/b/c/d/e", ". in middle"},
		{[]string{"0", "a/b"}, "0/a/b", "0 at start"},
		{[]string{"/0", "a/b"}, "/0/a/b", "/0 at start"},
		{[]string{"a/b/c", "0", "d/e"}, "a/b/c/0/d/e", "0 in middle"},
		{[]string{"a/b", "0"}, "a/b/0", "0 at end"},

		// checked against Stow::Util::join_paths
		{[]string{}, "", "no parts"},
		{[]string{"..foo/..", "a"}, "..foo/../a", "component starting with .. is kept"},
		{[]string{"a", "/.."}, "/", "/.. collapses to /"},
		{[]string{"/usr /", " /bin/"}, "/usr / /bin", "spaces are kept"},
	}

	for _, tc := range tests {
		got := JoinPaths(tc.in)
		if got != tc.want {
			t.Errorf("\nJoinPaths %s test failed:\nIn: %q\nGot: %q\nWant: %q", tc.scenario, tc.in, got, tc.want)
		}
	}
}

// Cases from stow/t/parent.t
func TestParent(t *testing.T) {
	tests := []struct {
		in       string
		want     string
		scenario string
	}{
		{"a/b/c", "a/b", "no leading or trailing /"},
		{"/a/b/c", "/a/b", "leading /"},
		{"a/b/c/", "a/b", "trailing /"},
		{"/////a///b///c///", "/a/b", "multiple /"},
		{"a", "", "empty parent"},

		// checked against Stow::Util::parent
		{"/a", "", "single absolute component"},
		{"/", "", "root"},
		{"", "", "empty"},
	}

	for _, tc := range tests {
		got := Parent([]string{tc.in})
		if got != tc.want {
			t.Errorf("\nParent %s test failed:\nIn: %q\nGot: %q\nWant: %q", tc.scenario, tc.in, got, tc.want)
		}
	}
}

// stow/t has no unit test for canon_path
func TestCanonPath(t *testing.T) {
	real, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	sub := filepath.Join(real, "sub")
	link := filepath.Join(real, "link")
	if err := os.Mkdir(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(sub, link); err != nil {
		t.Fatal(err)
	}

	t.Chdir(sub)

	tests := []struct {
		in       string
		want     string
		scenario string
	}{
		{real, real, "absolute"},
		{"..", real, "relative"},
		{link, sub, "symlink is resolved"},
		{".", sub, "current dir"},
	}

	for _, tc := range tests {
		got := CanonPath(tc.in)
		if got != tc.want {
			t.Errorf("\nCanonPath %s test failed:\nIn: %q\nGot: %q\nWant: %q", tc.scenario, tc.in, got, tc.want)
		}
	}

	// $PWD names the symlink, Perl's getcwd still returns the real dir
	t.Chdir(link)
	t.Setenv("PWD", link)
	if got := CanonPath("."); got != sub {
		t.Errorf("\nCanonPath symlinked $PWD test failed:\nGot: %q\nWant: %q", got, sub)
	}
}

// stow/t has no unit test for restore_cwd
func TestRestoreCwd(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(t.TempDir())

	RestoreCwd(dir)

	if got, _ := os.Getwd(); got != dir {
		t.Errorf("\nRestoreCwd test failed:\nGot: %q\nWant: %q", got, dir)
	}
}

// Cases from the adjust_dotfile() subtest in stow/t/dotfiles.t
func TestAdjustDotfile(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"file", "file"},
		{"dot-", "dot-"},
		{"dot-.", "dot-."},
		{"dot-file", ".file"},

		// checked against Stow::Util::adjust_dotfile
		{"dot-.testrc2", "dot-.testrc2"},
		{"dot-..", "dot-.."},
		{"dot-dot-x", ".dot-x"},
		{"dot-\n", ".\n"},
		{"testing/dot-testrc", "testing/dot-testrc"},
		{"", ""},
	}

	for _, tc := range tests {
		got := AdjustDotfile(tc.in)
		if got != tc.want {
			t.Errorf("\nAdjustDotfile(%q) test failed:\nGot: %q\nWant: %q", tc.in, got, tc.want)
		}
	}
}

// Cases from the unadjust_dotfile() subtest in stow/t/dotfiles.t
func TestUnAdjustDotfile(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"file", "file"},
		{".", "."},
		{"..", ".."},
		{".file", "dot-file"},

		// checked against Stow::Util::unadjust_dotfile
		{"..file", "dot-.file"},
		{"...", "dot-.."},
		{"", ""},
		{"testing/.testrc2", "testing/.testrc2"},
		{"./testing/testrc4", "dot-/testing/testrc4"},
		{"~/testing/.testrc7", "~/testing/.testrc7"},
		{".\n", ".\n"},
		{"..\n", "..\n"},
		{"\n.", "\n."},
	}

	for _, tc := range tests {
		got := UnAdjustDotfile(tc.in)
		if got != tc.want {
			t.Errorf("\nUnAdjustDotfile(%q) test failed:\nGot: %q\nWant: %q", tc.in, got, tc.want)
		}
	}
}

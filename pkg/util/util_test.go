package util

import "testing"

func TestJoinPaths(t *testing.T) {
	got1 := JoinPaths([]string{"/usr", "/bin"})
	got2 := JoinPaths([]string{"/usr /", " /bin/"})
	got3 := JoinPaths([]string{" //usr ", " /bin// "})
	got4 := JoinPaths([]string{".", "// usr //", "share", "../", "/bin/"})

	want := "/usr/bin"

	if got1 != want {
		t.Errorf("\nJoinPaths test 1 failed:\nGot: %q\nWant: %q\n", got1, want)
	}
	if got2 != want {
		t.Errorf("\nJoinPaths test 2 failed:\nGot: %q\nWant: %q\n", got2, want)
	}
	if got3 != want {
		t.Errorf("\nJoinPaths test 3 failed:\nGot: %q\nWant: %q\n", got3, want)
	}
	if got4 != want {
		t.Errorf("\nJoinPaths test 4 failed:\nGot: %q\nWant: %q\n", got4, want)
	}
}

func TestParent(t *testing.T) {
	in1 := []string{"usr", "bin"}
	in2 := []string{"usr"}
	in3 := []string{}
	got1 := Parent(in1)
	got2 := Parent(in2)
	got3 := Parent(in3)

	want1 := "/usr"
	want2 := "/"

	if got1 != want1 {
		t.Errorf("\nParent test failed:\nGot: %q\nWant: %q", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("\nParent test failed:\nGot: %q\nWant: %q", got2, want2)
	}
	if got3 != want2 {
		t.Errorf("\nParent test failed:\nGot: %q\nWant: %q", got3, want2)
	}
}

func TestCannonPath(t *testing.T) {
	got1 := CanonPath("/usr/bin")

	want1 := "/usr/bin"

	if got1 != want1 {
		t.Errorf("\nCannonPath test failed:\nGot: %q\nWant: %q", got1, want1)
	}
}

/*
func TestRestoreCwd(t *testing.T){
  var got1 string
}
*/

func TestAdjustDotfile(t *testing.T) {
	got1 := AdjustDotfile("/usr/bin/testing/dot-testrc")
	got2 := AdjustDotfile("testing/dot-.testrc2")
	got3 := AdjustDotfile("/home/testing/testrc3")
	got4 := AdjustDotfile("./testing/.testrc4")
	got5 := AdjustDotfile("./testing/dot-")
	got6 := AdjustDotfile("./testing/dot-.")
	got7 := AdjustDotfile("~/testing/dot-testrc7")

	want1 := "/usr/bin/testing/.testrc"
	want2 := "testing/.testrc2"
	want3 := "/home/testing/testrc3"
	want4 := "./testing/.testrc4"
	want5 := "./testing/dot-"
	want6 := "./testing/dot-."
	want7 := "~/testing/.testrc7"

	if got1 != want1 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got1, want1)
	}

	if got2 != want2 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got2, want2)
	}

	if got3 != want3 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got3, want3)
	}

	if got4 != want4 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got4, want4)
	}

	if got5 != want5 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got5, want5)
	}

	if got6 != want6 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got6, want6)
	}

	if got7 != want7 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got7, want7)
	}
}

func TestUnAdjustDotfile(t *testing.T) {
	got1 := UnAdjustDotfile("/usr/bin/testing/.testrc")
	got2 := UnAdjustDotfile("testing/.testrc2")
	got3 := UnAdjustDotfile("/home/testing/.testrc3")
	got4 := UnAdjustDotfile("./testing/testrc4")
	got5 := UnAdjustDotfile("./testing/dot-")
	got6 := UnAdjustDotfile("./testing/dot-.")
	got7 := UnAdjustDotfile("~/testing/.testrc7")

	want1 := "/usr/bin/testing/dot-testrc"
	want2 := "testing/dot-testrc2"
	want3 := "/home/testing/dot-testrc3"
	want4 := "./testing/testrc4"
	want5 := "./testing/dot-"
	want6 := "./testing/dot-."
	want7 := "~/testing/dot-testrc7"

	if got1 != want1 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got1, want1)
	}

	if got2 != want2 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got2, want2)
	}

	if got3 != want3 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got3, want3)
	}

	if got4 != want4 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got4, want4)
	}

	if got5 != want5 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got5, want5)
	}

	if got6 != want6 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got6, want6)
	}

	if got7 != want7 {
		t.Errorf("\nAdjustDotFile test failed:\nGot: %q\nWant: %q", got7, want7)
	}
}

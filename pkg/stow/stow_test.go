package stow

import "testing"

func TestJoinPaths(t *testing.T) {
	var got1 string = JoinPaths("/usr", "/bin")
	var got2 string = JoinPaths("/usr /", " /bin/")
	var got3 string = JoinPaths(" //usr ", " /bin// ")

	var want string = "/usr/bin"

	if got1 != want {
		t.Errorf("\nJoinPaths test failed:\nGot: %q\nWant: %q\n", got1, want)
	}
	if got2 != want {
		t.Errorf("\nJoinPaths test failed:\nGot: %q\nWant: %q\n", got2, want)
	}
	if got3 != want {
		t.Errorf("\nJoinPaths test failed:\nGot: %q\nWant: %q\n", got3, want)
	}
}

func TestParent(t *testing.T) {
	in1 := []string{"usr", "bin"}
	in2 := []string{"usr"}
	in3 := []string{}
	var got1 []string = Parent(in1)
	var got2 []string = Parent(in2)
	var got3 []string = Parent(in3)

	var want1 string = "/usr"
	var want2 string = "/"

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
	var got1 string = CannonPath("/usr/bin")

	var want1 string = "/usr/bin"

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
	var got1 string = AdjustDotFile("./testing/dot-testrc")
	var got2 string = AdjustDotFile("./testing/dot-.testrc2")
	var got3 string = AdjustDotFile("./testing/testrc3")
	var got4 string = AdjustDotFile("./testing/.testrc4")
	var got5 string = AdjustDotFile("./testing/dot-")
	var got6 string = AdjustDotFile("./testing/dot-.")

	var want1 string = "./testing/.testrc"
	var want2 string = "./testing/.testrc2"
	var want3 string = "./testing/testrc3"
	var want4 string = "./testing/.testrc4"
	var want4 string = "./testing/dot-"
	var want4 string = "./testing/dot-."

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
}

// Package util contains an implementation of the utility functions built into
// GNU Stow
//
// GoStow is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// GoStow is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see https://www.gnu.org/licenses/.
package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// JoinPaths concatenates given paths (path1, path2, ...), factors out
// redundant path elements: '//' => '/' and 'a/b/../c' => 'a/c' and
// returns a concatenated string.
func JoinPaths(paths []string) string {
	var filteredPaths []string

	for _, str := range paths {
		trimmed := strings.TrimSpace(str)
		trimmed = strings.Replace(trimmed, "/", "", -1)
		trimmed = strings.TrimSpace(trimmed)
		filteredPaths = append(filteredPaths, trimmed)
	}

	return filepath.Clean("/" + strings.Join(filteredPaths, "/"))
}

// Parent finds the parent of the given path.
// Parameter path [] => a slice containg components of the path
// Returns returns a path string
func Parent(path []string) string {
	if len(path) <= 1 {
		return "/"
	}
	return "/" + strings.Join(path[:len(path)-1], "/")
}

// CanonPath finds absolute canonical path of given path
// Parameters path is a string
// Returns absolute canonical path
// Throws fatal error if unable to get or change directory.
func CanonPath(path string) string {
	var cwd string
	var err error
	var canonPath string

	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("CanonPath cannot get current working directory: %q", err)
	}

	err = os.Chdir(path)
	if err != nil {
		log.Fatalf("CanonPath cannot change working dir to %q: %q", path, err)
	}

	canonPath, err = os.Getwd()
	if err != nil {
		log.Fatalf("CanonPath cannot get canonPath working directory: %q", err)
	}

	RestoreCwd(cwd)

	return canonPath
}

// RestoreCwd sets working dir to dir specified in parameters
// Parameters prev string, a string to the absolute dir path to set as the
// working directory.
// Does not return a value
// Throws fatal error if unable to change directory.
func RestoreCwd(prev string) {
	err := os.Chdir(prev)
	if err != nil {
		log.Fatalf("RestoreCwd believes your current directory, %q, has vanished: %q", prev, err)
	}
	return
}

// AdjustDotfile updates dotfile names
// Parameters: target string, the name of the file to be renamed
// Returns string of modified file name
func AdjustDotfile(target string) string {
	parts := strings.Split(target, "/")
	pathLen := len(parts) - 1
	fileName := parts[pathLen]

	switch fileName {
	case ".":
		return target
	case "..":
		return target
	case "dot-":
		return target
	case "dot-.":
		return target
	}

	if strings.HasPrefix(fileName, "dot-.") {
		parts = append(parts[:pathLen], strings.Replace(fileName, "dot-", "", 1))
	} else if strings.HasPrefix(fileName, "dot-") {
		parts = append(parts[:pathLen], strings.Replace(fileName, "dot-", ".", 1))
	}

	return strings.Join(parts, "/")
}

// UnAdjustDotfile reverts dotfile names to pre-stow format. Needed when
// unstowing with --compat and --dotfiles.
// Parameters target string, the name of the file to be renamed.
// Returns string of modified file name.
func UnAdjustDotfile(target string) string {
	parts := strings.Split(target, "/")
	pathLen := len(parts) - 1
	fileName := parts[pathLen]

	switch fileName {
	case ".":
		return target
	case "..":
		return target
	}

	if strings.HasPrefix(fileName, ".") {
		parts = append(parts[:pathLen], strings.Replace(fileName, ".", "dot-", 1))
	}

	return strings.Join(parts, "/")
}

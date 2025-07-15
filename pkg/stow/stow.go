// Package stow is the backend Go module for GoStow, a program for managing
// the installation of software packages, keeping them separate
// (C</usr/local/stow/emacs> vs. C</usr/local/stow/perl>, for example)
// while making them appear to be installed in the same place
// (C</usr/local>).
//
// GoStow doesn't store an extra state between runs, so there's no danger
// of mangling directories when file hierarchies don't match the
// database. Also, stow will never delete any files, directories, or
// links that appear in a stow directory, so it is always possible to
// rebuild the target tree.
//
// This is intended to be a 1:1 implementation of the original Stow source code
// found in the git repo https://git.savannah.gnu.org/git/stow.git.
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
package stow

// New
func New() {
	return
}

// GetVerbosity
func GetVerbosity() {
	return
}

// SetStowDir calculates relative path from target directory to stow directory.
// This will be commonly used as a prefix for constructing and
// recognising symlinks "installed" in the target directory which
// point to package files under the stow directory.
func SetStowDir() {
	return
}

// InitState
func InitState() {
	return
}

// PlanStow plans which symlink/directory creation/removal tasks need to be
// executed in order to stow the given packages.
// Returns a slice of conflicts.
func PlanStow(packages []string) []string {
	return [""]string
}

// WithinTargetDo executes code within target directory, preserving cwd.
// This is done to ensure that the consumer of the Stow interface doesn't
// have to worry about (a) what their cwd is, and (b) that their cwd
// might change.
// Parameter: code, Anonymous subroutine to execute within target dir.
func WithinTargetDo() {
	return
}

// StowContents stows the contents of the given directory.
//
// Parameter: stow_path - Relative path from current (i.e. target) directory to
// the stow dir containing the package to be stowed.
//
// Parameter: pkg_subdir - Subdirectory of the installation image in the package
// directory which needs stowing as a symlink which points to it.  This is
// relative to the top-level package directory.
//
// Parameter: target_subdir - Subdirectory of the target directory which either
// needs a symlink to the corresponding package subdirectory in the installation
// image, or if it's an existing directory, it's an unfolded tree which may need
// to be folded or recursed into.
func StowContents(stow_path string, pkg_subdir string, target_subdir string) {
	return
}

// StowNode - Stow the given node.
//
// Parameter: stow_path - Relative path from current (i.e. target) directory to
// the stow dir containing the node to be stowed.
//
// Parameter: stow_package - The package containing the node being stowed.
//
// Parameter: pkg_subpath - Subpath of the installation image in the package
// directory which needs stowing as a symlink which points to it.  This is
// relative to the top-level package directory.
//
// Parameter: target_subpath - Subpath of the target directory which either
// needs a symlink to the corresponding package subpathectory in the
// installation image, or if it's an existing directory, it's an unfolded tree
// which may need to be folded or recursed into.
func StowNode(stow_path string, stow_package string, pkg_subpath string, target_subpath string) {
	return
}

// ShouldSkipTarget - Determine whether target_subdir is a stow directory which
// should not be stowed to or unstowed from. This mechanism protects stow
// directories from being altered by stow, and is a necessary safety check
// because the stow directory could live beneath the target directory.
//
// Paramter: target_subdir - relative path to symlink target from the current
// directory.
//
// Return: true if target is a stow directory
func ShouldSkipTarget(target_subdir string) bool {
	return false
}

// MarkedStowDir - creates a .stow file in the target directory to mark that it
// is a stow directory.
//
// Parameter: dir - the directory to mark as a stow directory.
func MarkStowDir(dir string) {
	return
}

// LinkOwnedByPackage - Determine whether the given link points to a member of
// a stowed package.
//
// Parameter: target_subpath - Path to a symbolic link under current directory.
//
// Parameter: link_dest - Where that link points to.
//
// Return: stow_package if link is owned by stow, otherwise ”.
func LinkOwnedByPackage(target_subpath string, link_dest string) string {
	return
}

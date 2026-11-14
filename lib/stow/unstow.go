package stow

// PlanUnstow plans which symlink/directory creation/removal tasks need to be
// executed in order to unstow the given packages.  Any potential conflicts are
// then accessible via get_conflicts().
func PlanUnstow() {
	return
}

// UnstowContents - Unstow the contents of the given directory
//
// Parameter: stow_package - The package whose contents are being unstowed.
//
// Parameter: pkg_subdir - Subdirectory of the installation image in the
// package directory which may need a symlink pointing to it to be unstowed.
// This is relative to the top-level package directory.
//
// Parameter: target_subdir - Subdirectory of the target directory which either
// needs unstowing of a symlink to the corresponding package subdirectory in the
// installation image, or if it's an existing directory, it's an unfolded tree
// which may need to be recursed into.
func UnstowContents(stow_package string, pkg_subdir string, target_subdir string) {
	return
}

// UnstowNode - Unstow the given node.
//
// Parameter: stow_package - The package containing the node being unstowed.
//
// Parameter: pkg_subpath - Subpath of the installation image in the package
// directory which needs stowing as a symlink which points to it.  This is
// relative to the top-level package directory.
//
// Parameter: target_subpath - Subpath of the target directory which either
// needs a symlink to the corresponding package subpathectory in the
// installation image, or if it's an existing directory, it's an unfolded tree
// which may need to be folded or recursed into.
func UnstowNode(stow_package string, pkg_subpath string, target_subpath string) {
	return
}

// UnstowLinkNode -
//
// Parameter: stow_package - The package containing the node being unstowed.
//
// Parameter: pkg_subpath - Subpath of the installation image in the package
// directory which needs stowing as a symlink which points to it.  This is
// relative to the top-level package directory.
//
// Parameter: target_subpath - Subpath of the target directory which either
// needs a symlink to the corresponding package subpathectory in the
// installation image, or if it's an existing directory, it's an unfolded tree
// which may need to be folded or recursed into.
func UnstowLinkNode(stow_package string, pkg_subpath string, target_subpath string) {
	return
}

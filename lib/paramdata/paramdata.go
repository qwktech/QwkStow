package paramdata

var (
	options *Options
)

type Options struct {
	manditory    *ManditoryOptions
	nonManditory *NonManditoryOptions
}

type ManditoryOptions struct {
	dir    string
	target string
}

type NonManditoryOptions struct {
	conflicts    bool
	simulate     bool
	verbose      bool
	paranoid     bool
	compat       bool
	testMode     bool
	dotfiles     bool
	adopt        bool
	noFolding    bool
	ignoreList   []string
	overrideList []string
	deferList    []string
}

// ProcessOptions parses and process command line and .stowrc file options.
// Parameters: none
// Returns   : options *Options, pkgs_to_unstow []string, pkgs_to_stow []string,
// err error
// Comments  : checks @ARGV for valid package names
func ProcessOptions() (*Options, []string, []string, error) {
  options, pkgs_to_unstow, pkgs_to_stow, err := ParseOptions(os.Args)
  if err != nil {
    log.Fatalf()
	return options, pkgs_to_unstow, pkgs_to_stow, nil
}

func ParseOptions(arg []string) (*Options, []string, []string, error) {
	manditory := *ManditoryOptions{}
	nonManditory := *NonManditoryOptions{}
	options := *Options{manditory: manditory, nonManditory: nonManditory}

	return options, [""]string, [""]string, nil
}

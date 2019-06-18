package fsindex

// Default Settings
var (
	DefaultSettings = Settings{
		OmitRootNameFromPath:       false,
		StripFileExtensionFromName: true,
		UnknownCharsToDash:         false,
	}
)

// Settings will slightly alter how the `Refresh` method runs.
type Settings struct {
	// OmitRootNameFromPath will strip the root directory-name from indexed path targets.
	// For example if true, a path converted to "http path": path-in: "c:/mypath/mysubdir/my-target-path", path-out: "/".
	// If set to (default) false: path-in: "c:/mypath/mysubdir/my-target-path", path-out: "/my-target-path".
	OmitRootNameFromPath       bool
	StripFileExtensionFromName bool
	UnknownCharsToDash         bool // not uet supported.
	// this tells us weather or not to use full link-path
	// such as `http://[server:port]/` when generating JSON.
	HardLinks bool
}

// Model is the same as PathEntry but with Settings
type Model struct {
	PathEntry
	SimpleModel `json:"-"`
	Settings    `json:"-"`
}

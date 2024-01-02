package git

var Version = "dev"
var Commit = ""

func Information() map[string]any {
	return map[string]any{
		"version": Version,
		"commit":  Commit,
	}
}

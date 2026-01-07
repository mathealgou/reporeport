package configs

var DisallowedLibDirs = []string{
	"lib/",
	"lib\\",
	"vendor/",
	"vendor\\",
	"node_modules/",
	"node_modules\\",
	"venv/",
	"venv\\",
	"__pycache__/",
	"__pycache__\\",
	".venv/",
	".venv\\",
}

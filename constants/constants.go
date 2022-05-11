package constants

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/tran/dfs"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Styles
	PRIMARY_COLOR = lipgloss.Color("#1E90FF")
	CYAN_COLOR = lipgloss.Color("#00FFFF")
    GREEN_COLOR = "#04B575"
  	RED_COLOR = "#FF4141"
  	SUBTITLE_COLOR = lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"}
 	GRAY_COLOR = lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"}
 	SUCCESS_BACKGROUND = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color(GREEN_COLOR)).
		PaddingLeft(1).
		PaddingRight(1)
	FAIL_BACKGROUND = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color(RED_COLOR)).
		PaddingLeft(1).
		PaddingRight(1)
	SUCCESS_FOREGROUND = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(GREEN_COLOR))
	FAIL_FOREGROUND = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(RED_COLOR))

	// File Paths
	HomeDir, _ = dfs.GetHomeDirectory()
	BotwayConfigFile = filepath.Join(HomeDir, ".botway", "botway.json")
	BotwayConfig, Berr = ioutil.ReadFile(BotwayConfigFile)
	BotConfig, Oerr = ioutil.ReadFile(".botway.yaml")
	Guilds, Gerr = ioutil.ReadFile("guilds.json")

	BotwayDirPath = func () string {
		if runtime.GOOS == "windows" {
			return `$HOME\\.botway`
		} else {
			return `$HOME/.botway`
		}
	}
)

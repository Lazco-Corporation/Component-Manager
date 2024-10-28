package module

import (
	"errors"
	"io"
	"net/http"
	"regexp"

	"github.com/gookit/color"
	"github.com/hashicorp/go-version"
	"github.com/urfave/cli/v2"
)

func CheckRemoteVersion(ctx *cli.Context, muteUpToDate bool) (bool, error) {
	// URL of the raw content on GitHub
	url := "https://raw.githubusercontent.com/Lazco-Studio/Component-Manager/main/main.go"

	// Fetch the raw content
	resp, err := http.Get(url)
	if err != nil {
		return false, errors.New("no internet connection")
	}
	defer resp.Body.Close()

	// Read the content
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Define a regex pattern to find the version
	// This pattern looks for the string 'Version: "vX.X.X",' where X.X.X can be any version number
	pattern := `Version:\s*\"v([0-9]+(\.[0-9]+)*)\",`
	regex := regexp.MustCompile(pattern)

	// Find the first match in the content
	matches := regex.FindStringSubmatch(string(content))
	if len(matches) != 0 {
		remoteVersion, err := version.NewSemver(matches[1])
		if err != nil {
			return false, err
		}
		currentVersion, err := version.NewSemver(ctx.App.Version)
		if err != nil {
			return false, err
		}

		if currentVersion.LessThan(remoteVersion) {
			color.Yellowp("New version available: ")
			color.Redp(currentVersion.String())
			color.Yellowp(" → ")
			color.Greenln(remoteVersion.String())

			color.Magentap("Run \"")
			color.Cyanp("cm update")
			color.Magentaln("\" to update.")

			return false, nil
		}

		if !muteUpToDate {
			color.Greenln("Up to date.")
			return true, nil
		}
		return true, nil
	} else {
		color.Yellowln("Warning: could not get remote version.")
		return false, nil
	}
}

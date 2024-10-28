package command

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/urfave/cli/v2"

	"Component-Manager/module"
)

func DownloadScript() (string, error) {
	tempDir := os.TempDir()
	var scriptURL string

	if runtime.GOOS == "windows" {
		scriptURL = "https://short.on-cloud.tw/cm-install-script-windows"
	} else {
		scriptURL = "https://short.on-cloud.tw/cm-install-script"
	}
	scriptDir, err := module.Download(scriptURL, filepath.Join(tempDir, "cm-install-script"))
	if err != nil {
		return "", err
	}
	os.Chmod(scriptDir, 0755)

	return scriptDir, nil
}

func Update(ctx *cli.Context) error {
	scriptDir, err := DownloadScript()
	if err != nil {
		return err
	}
	cmd := exec.Command(scriptDir)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	os.Remove(scriptDir)

	return nil
}

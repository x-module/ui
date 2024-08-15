package widgets

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/gen2brain/beeep"
)

const (
	icon  = "/Users/lijin/go/ui/qrcodeSymbol.png"
	title = "Client"
)

type SystemNotice struct {
	iconPath string
	message  string
}

func NewSystemNotice() (*SystemNotice, error) {
	return &SystemNotice{
		//iconPath: filepath.Join(absolutePath, icon),
		iconPath: icon,
	}, nil
}

func (s *SystemNotice) Notice(message string) error {
	err := beeep.Notify(title, message, s.iconPath)
	if err != nil {
		return err
	}
	return nil
}

func getAbsolutePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("error getting executable path: %s", err.Error())
	}

	exSym, err := filepath.EvalSymlinks(ex)
	if err != nil {
		return "", fmt.Errorf("error getting filepath after evaluating sym links")
	}

	return path.Dir(exSym), nil
}

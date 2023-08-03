package systemos

import (
	windows "handy-translate/systemos/Windows"
	"runtime"
)

type SystemOS interface {
	Show()
}

func GetOS() SystemOS {
	switch runtime.GOOS {
	case "darwin":
		return nil
	case "windows":
		return new(windows.Windows)
	}
	return nil
}

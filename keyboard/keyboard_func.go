// +build !windows

package keyboard

import (
	"fmt"

	"golang/types"
)

func install(fn HookHandler, c chan<- types.KeyboardEvent) error {
	return fmt.Errorf("keyboard: not supported")
}

func uninstall() error {
	return fmt.Errorf("keyboard: not supported")
}

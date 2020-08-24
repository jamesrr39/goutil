//+build linux

package notification

import "os/exec"

func Send(title, summary string) error {
	return exec.Command("notify-send", title, summary).Run()
}

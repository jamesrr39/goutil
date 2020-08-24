//+build darwin

package notification

func Send(title, summary string) error {
	return exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s"`, summary, title).Run()
}

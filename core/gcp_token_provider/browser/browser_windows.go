//go:build windows

package browser

import "golang.org/x/sys/windows"

func openBrowser(url string, cmdOptions []CmdOption) error {
	return windows.ShellExecute(0, nil, windows.StringToUTF16Ptr(url), nil, nil, windows.SW_SHOWNORMAL)
}

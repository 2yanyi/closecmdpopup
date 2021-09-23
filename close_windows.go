package closecmdpopup

import "syscall"

var _GetConsoleWindow = syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
var _ShowWindow = syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")

func commandPopup(show bool) (zero int) {
	r1, _, _ := _GetConsoleWindow.Call()
	if r1 == 0 {
		return
	}
	if show {
		_, _, _ = _ShowWindow.Call(r1, 9)
	} else {
		_, _, _ = _ShowWindow.Call(r1, 0)
	}
	return
}

var _ = commandPopup(false)

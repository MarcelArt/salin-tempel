package tray

import (
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

func onQuitClicked() {
	systray.Quit()
}

func onQRCodeClicked() {
	open.Run("http://localhost:3000/qr")
}

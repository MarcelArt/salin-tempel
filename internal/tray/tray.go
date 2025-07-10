package tray

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcelArt/salin-tempel/internal/server"
	"github.com/getlantern/systray"
	"golang.design/x/clipboard"
)

var scissorIcon = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="white" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-scissors-icon lucide-scissors"><circle cx="6" cy="6" r="3"/><path d="M8.12 8.12 12 12"/><path d="M20 4 8.12 15.88"/><circle cx="6" cy="18" r="3"/><path d="M14.8 14.8 20 20"/></svg>`)

func onReady() {
	err := clipboard.Init()
	if err != nil {
		log.Fatalln(err.Error())
	}
	systray.SetIcon(scissorIcon)
	systray.SetTitle("Salin Tempel")
	mCurrentClipboard := systray.AddMenuItem("", "")
	mShowQR := systray.AddMenuItem("Show QR Code", "Show the QR code for the app")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-mShowQR.ClickedCh:
				onQRCodeClicked()
			case <-mQuit.ClickedCh:
				onQuitClicked()
				return
			}
		}
	}()

	go func() {
		for {
			currentClipboard := clipboard.Read(clipboard.FmtText)
			mCurrentClipboard.SetTitle(fmt.Sprintf("Clipboard: %s", string(currentClipboard)))
		}
	}()

	server.Run()
}

func onExit() {
	os.Exit(0)
}

func Run() {
	systray.Run(onReady, onExit)
}

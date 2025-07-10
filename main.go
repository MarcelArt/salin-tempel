package main

import (
	_ "github.com/MarcelArt/salin-tempel/docs" // Import the generated docs
	"github.com/MarcelArt/salin-tempel/internal/tray"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	tray.Run()
}

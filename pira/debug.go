package pira

import (
	"fmt"
)

func (app *App) PrintConfig() {
	fmt.Fprintf(app.Out, "%+v\n", app.Config)
}

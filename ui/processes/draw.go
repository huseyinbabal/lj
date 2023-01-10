package processes

import (
	"github.com/gdamore/tcell/v2"
)

// Draw draws this primitive onto the screen.
func (i *Processes) Draw(screen tcell.Screen) {
	i.refresh()
	i.Box.DrawForSubclass(screen, i)
	i.Box.SetBorder(false)
	x, y, width, height := i.GetInnerRect()
	i.table.SetRect(x, y, width, height)
	i.table.SetBorder(true)

	i.table.Draw(screen)
	x, y, width, height = i.table.GetInnerRect()
	// error dialog
	if i.errorDialog.IsDisplay() {
		i.errorDialog.SetRect(x, y, width, height)
		i.errorDialog.Draw(screen)
		return
	}
	// command dialog dialog
	if i.cmdDialog.IsDisplay() {
		i.cmdDialog.SetRect(x, y, width, height)
		i.cmdDialog.Draw(screen)
		return
	}
	// command input dialog
	if i.cmdInputDialog.IsDisplay() {
		i.cmdInputDialog.SetRect(x, y, width, height)
		i.cmdInputDialog.Draw(screen)
		return
	}
	// message dialog
	if i.messageDialog.IsDisplay() {
		i.messageDialog.SetRect(x, y, width, height+1)
		i.messageDialog.Draw(screen)
		return
	}
	// confirm dialog
	if i.confirmDialog.IsDisplay() {
		i.confirmDialog.SetRect(x, y, width, height)
		i.confirmDialog.Draw(screen)
		return
	}
	// progress dialog
	if i.progressDialog.IsDisplay() {
		i.progressDialog.SetRect(x, y, width, height)
		i.progressDialog.Draw(screen)

	}
	// top dialog
	if i.topDialog.IsDisplay() {
		i.topDialog.SetRect(x, y, width, height)
		i.topDialog.Draw(screen)
		return
	}
}

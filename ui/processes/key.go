package processes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
)

// InputHandler returns the handler for this primitive.
func (i *Processes) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return i.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		log.Debug().Msgf("view: processes event %v received", event.Key())
		if i.progressDialog.IsDisplay() {
			return
		}
		// command dialog handler
		if i.cmdDialog.HasFocus() {
			if cmdHandler := i.cmdDialog.InputHandler(); cmdHandler != nil {
				cmdHandler(event, setFocus)
			}
		}
		// input dialog handler
		if i.cmdInputDialog.HasFocus() {
			if cmdInputHandler := i.cmdInputDialog.InputHandler(); cmdInputHandler != nil {
				cmdInputHandler(event, setFocus)
			}
		}

		// message dialog handler
		if i.messageDialog.HasFocus() {
			if messageDialogHandler := i.messageDialog.InputHandler(); messageDialogHandler != nil {
				messageDialogHandler(event, setFocus)
			}
		}
		// confirm dialog handler
		if i.confirmDialog.HasFocus() {
			if confirmDialogHandler := i.confirmDialog.InputHandler(); confirmDialogHandler != nil {
				confirmDialogHandler(event, setFocus)
			}
		}
		// table handlers
		if i.table.HasFocus() {
			if event.Key() == tcell.KeyCtrlV || event.Key() == tcell.KeyEnter {
				if i.cmdDialog.GetCommandCount() <= 1 {
					return
				}
				i.selectedID, i.selectedProcess = i.getSelectedItem()
				i.cmdDialog.Display()
			}
			if tableHandler := i.table.InputHandler(); tableHandler != nil {
				tableHandler(event, setFocus)
			}
		}
		// error dialog handler
		if i.errorDialog.HasFocus() {
			if errorDialogHandler := i.errorDialog.InputHandler(); errorDialogHandler != nil {
				errorDialogHandler(event, setFocus)
			}
		}

		// container top dialog handler
		if i.topDialog.HasFocus() {
			if cntTopDialogHandler := i.topDialog.InputHandler(); cntTopDialogHandler != nil {
				cntTopDialogHandler(event, setFocus)
			}
		}
		setFocus(i)
	})
}

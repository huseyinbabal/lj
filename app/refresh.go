package app

import (
	"time"

	"github.com/rs/zerolog/log"
)

func (app *App) refresh() {
	log.Debug().Msg("app: starting refresh loop")
	tick := time.NewTicker(refreshInterval)
	for {
		select {
		case <-tick.C:
			/*connOK, connMsg := app.health.ConnOK()
			if connOK {
				if app.needInitUI {
					// init ui after reconnection
					app.initUI()
					app.needInitUI = false
					app.pages.SwitchToPage(app.currentPage)
					app.setFocus(app.currentPage)

				}
				eventTypes := app.health.GetEvents()
				// update events
				for _, evt := range eventTypes {
					app.updatePageData(evt)
				}
				if app.health.HasNewEvent() {
					app.system.SetEventMessage(app.health.GetEventMessages())
				}
			} else {
				// set init ui to true
				app.needInitUI = true
				app.clearUIData()
				app.pages.SwitchToPage(app.connection.GetTitle())
				app.connection.SetErrorMessage(connMsg)
				app.setFocus(app.connection.GetTitle())
			}*/
			app.initUI()
			app.needInitUI = false
			app.initInfoBar()
			//app.infoBar.UpdateConnStatus(true)
			app.Application.Draw()
		}
	}
}

func (app *App) setFocus(page string) {
	switch page {
	case app.processes.GetTitle():
		app.Application.SetFocus(app.processes)
	}
}

func (app *App) updatePageData(eventType string) {
	switch eventType {
	case "instance":
		app.processes.UpdateData()
	}

}

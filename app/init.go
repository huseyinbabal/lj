package app

import "github.com/rs/zerolog/log"

func (app *App) initUI() {
	//app.connection.Reset()
	//connOK, _ := app.health.ConnOK()
	//if connOK {
	app.processes.UpdateData()
	app.initInfoBar()
	//}
}

func (app *App) initInfoBar() {
	// update basic information
	//hostname, kernel, ostype := app.health.GetSysInfo()
	//app.infoBar.UpdateBasicInfo("a", "b", "c")

	// udpate memory and swap usage
	//memUsage, swapUsage := app.health.GetSysUsage()
	//app.infoBar.UpdateSystemUsageInfo(1, 2)

	java, err := app.GetJava()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get java information")
	} else {
		app.infoBar.UpdateJava(java.version)
	}
}

func (app *App) clearUIData() {
	app.processes.ClearData()
}

package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/huseyinbabal/lj/internal/jvm"
	"github.com/huseyinbabal/lj/ui/infobar"
	processes "github.com/huseyinbabal/lj/ui/processes"
	"github.com/huseyinbabal/lj/ui/searchbar"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	refreshInterval = 10000 * time.Millisecond
)

type App struct {
	*tview.Application
	pages           *tview.Pages
	infoBar         *infobar.InfoBar
	searchBar       *searchbar.SearchBar
	processes       *processes.Processes
	menu            *tview.TextView
	currentPage     string
	needInitUI      bool
	jvm             *jvm.Jvm
	searchBarActive bool
	flex            *tview.Flex
}

// NewApp returns new app
func NewApp() *App {
	log.Info().Msg("app: new application")
	jvm := jvm.NewJvm()
	app := App{
		Application:     tview.NewApplication(),
		pages:           tview.NewPages(),
		needInitUI:      false,
		jvm:             jvm,
		searchBarActive: false,
	}
	//app.health = health.NewEngine(refreshInterval)

	app.infoBar = infobar.NewInfoBar()
	app.searchBar = searchbar.NewSearchBar(func(key tcell.Key) {
		app.flex.RemoveItem(app.searchBar)
		resource := app.searchBar.InputField.GetText()
		app.searchBar.InputField.SetText("")
		app.searchBarActive = !app.searchBarActive
		app.showResources(resource)
	})
	app.processes = processes.NewProcesses(jvm)

	app.pages.AddPage(app.processes.GetTitle(), app.processes, true, false)

	return &app
}

func (app *App) showResources(text string) {
	log.Info().Msgf("app: text %s", text)

	switch text {
	case "processes":
		//processes page
		log.Info().Msgf("app: switching to %s view", app.processes.GetTitle())
		app.pages.SwitchToPage(app.processes.GetTitle())
		app.SetFocus(app.processes)
		app.processes.UpdateData()
		app.currentPage = app.processes.GetTitle()
	default:
		log.Info().Msgf("app: default %s", text)

	}
}

// Run starts the application loop.
func (app *App) Run() error {
	log.Info().Msg("app: run")

	app.flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(app.infoBar, infobar.InfoBarViewHeight, 0, false).
		AddItem(app.pages, 0, 1, false)

	// start health check and event parser
	//app.health.Start()

	// initial update
	app.initUI()

	// listen for user input
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//connOK, _ := app.health.ConnOK()
		//if !connOK {
		//return event
		//}
		switch event.Name() {

		case KeyNameColon:
			// Colon: search resource
			log.Debug().Msgf("app: switching to %v view", event)
			if app.searchBarActive {
				app.flex.RemoveItem(app.searchBar)
				app.SetFocus(app.pages)
			} else {
				app.flex.RemoveItem(app.pages)
				app.flex.AddItem(app.searchBar, infobar.SearchBarViewHeight, 0, false)
				app.flex.AddItem(app.pages, 0, 1, false)
				app.SetFocus(app.searchBar.InputField)
			}
			app.searchBarActive = !app.searchBarActive

			return nil

		}

		return event
	})
	app.currentPage = app.processes.GetTitle()
	a := app.processes.GetTitle()
	app.pages.SwitchToPage(a)

	// start refresh loop
	go app.refresh()

	if err := app.SetRoot(app.flex, true).SetFocus(app.processes).EnableMouse(false).Run(); err != nil {
		return err
	}
	return nil
}

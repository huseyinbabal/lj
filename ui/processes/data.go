package processes

import (
	"fmt"
	"github.com/huseyinbabal/lj/internal/jvm"
	"github.com/huseyinbabal/lj/ui/utils"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"strings"
)

// UpdateData retreives containers list data
func (i *Processes) UpdateData() {
	data, err := i.jvm.ListProcesses()
	if err != nil {
		log.Error().Msgf("view: containers %s", err.Error())
		i.errorDialog.SetText(err.Error())
		i.errorDialog.Display()
		return
	}
	if err != nil {
		log.Error().Msgf("view: containers %s", err.Error())
		i.errorDialog.SetText(err.Error())
		i.errorDialog.Display()
		return
	}
	i.processes.mu.Lock()
	i.processes.report = data
	i.processes.mu.Unlock()
}

func (i *Processes) getData() []jvm.Process {
	i.processes.mu.Lock()
	data := i.processes.report
	i.processes.mu.Unlock()
	return data
}

// ClearData clears table data
func (i *Processes) ClearData() {
	i.table.Clear()
	expand := 1
	fgColor := utils.Styles.PageTable.HeaderRow.FgColor
	bgColor := utils.Styles.PageTable.HeaderRow.BgColor
	for k := 0; k < len(i.headers); k++ {
		i.table.SetCell(0, k,
			tview.NewTableCell(fmt.Sprintf("[black::b]%s", strings.ToUpper(i.headers[k]))).
				SetExpansion(expand).
				SetBackgroundColor(bgColor).
				SetTextColor(fgColor).
				SetAlign(tview.AlignLeft).
				SetSelectable(false))
	}
	i.table.SetTitle(fmt.Sprintf("[::b]%s[0]", strings.ToUpper(i.title)))
}

type processesReporter struct {
	processes []jvm.Process
}

func (ic processesReporter) securityGroups() string {
	//TODO: domain methods for instance
	return ""
}

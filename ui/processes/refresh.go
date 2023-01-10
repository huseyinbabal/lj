package processes

import (
	"fmt"
	"github.com/huseyinbabal/lj/ui/utils"
	"github.com/rivo/tview"
	"strings"
)

func (i *Processes) refresh() {
	i.table.Clear()
	expand := 1
	alignment := tview.AlignLeft
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
	rowIndex := 1

	processes := i.getData()
	i.table.SetTitle(fmt.Sprintf("[::b]%s[%d]", strings.ToUpper(i.title), len(processes)))
	for k := 0; k < len(processes); k++ {
		processID := processes[k].ID
		name := processes[k].Name

		// ID column
		i.table.SetCell(rowIndex, 0,
			tview.NewTableCell(processID).
				SetExpansion(expand).
				SetAlign(alignment))

		// Name column
		i.table.SetCell(rowIndex, 1,
			tview.NewTableCell(name).
				SetExpansion(expand).
				SetAlign(alignment))

		rowIndex++
	}

}

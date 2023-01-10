package infobar

import (
	"fmt"
	"github.com/huseyinbabal/lj/ui/utils"

	"github.com/gdamore/tcell/v2"

	"github.com/rivo/tview"
)

// InfoBarViewHeight info bar height
const (
	InfoBarViewHeight   = 5
	SearchBarViewHeight = 2
	row1                = 0
	row2                = 1
	row3                = 2
	row4                = 3
	memCellRow          = 4
	swapCellRow         = 5
	connOK              = "\u2705"
	connERR             = "\u274C"
)

// InfoBar implements the info bar primitive
type InfoBar struct {
	*tview.Box
	table  *tview.Table
	title  string
	connOK bool
}

// NewInfoBar returns info bar view
func NewInfoBar() *InfoBar {
	table := tview.NewTable()
	headerColor := utils.GetColorName(utils.Styles.InfoBar.ItemFgColor)
	emptyCell := func() *tview.TableCell {
		return tview.NewTableCell("")
	}

	// empty column
	for i := 0; i < 5; i++ {
		table.SetCell(i, 0, emptyCell())
	}

	// valueColor := Styles.InfoBar.ValueFgColor
	table.SetCell(row1, 1, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "Java Version:")))
	table.SetCell(row1, 2, emptyCell())

	table.SetCell(row2, 1, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "UserId:")))
	table.SetCell(row2, 2, emptyCell())

	table.SetCell(row3, 1, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "Arn:")))
	table.SetCell(row3, 2, emptyCell())

	table.SetCell(row4, 1, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "Region:")))
	table.SetCell(row4, 2, emptyCell())

	/*	table.SetCell(memCellRow, 1, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "Memory usage:")))
		table.SetCell(memCellRow, 2, tview.NewTableCell(progressUsageString(0.00)))

		table.SetCell(swapCellRow, 1, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "Swap usage:")))
		table.SetCell(swapCellRow, 2, tview.NewTableCell(progressUsageString(0.00)))*/

	// empty column
	for i := 0; i < 5; i++ {
		table.SetCell(i, 3, emptyCell())
	}

	/*table.SetCell(row1, 4, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "Kernel version:")))
	table.SetCell(row1, 5, emptyCell())

	table.SetCell(row2, 4, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "API version:")))
	table.SetCell(row2, 5, emptyCell())

	table.SetCell(row3, 4, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "OCI runtime:")))
	table.SetCell(row3, 5, emptyCell())

	table.SetCell(memCellRow, 4, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "Conmon version:")))
	table.SetCell(memCellRow, 5, emptyCell())

	table.SetCell(swapCellRow, 4, tview.NewTableCell(fmt.Sprintf("[%s::]%s", headerColor, "Buildah version:")))
	table.SetCell(swapCellRow, 5, emptyCell())*/

	// infobar
	infoBar := &InfoBar{
		Box:    tview.NewBox(),
		title:  "infobar",
		table:  table,
		connOK: false,
	}
	return infoBar
}

// UpdateJava updates aws identity values
func (info *InfoBar) UpdateJava(version string) {
	info.table.GetCell(row1, 2).SetText(version)
}

// UpdateBasicInfo updates hostname, kernel and os type values
func (info *InfoBar) UpdateBasicInfo(hostname string, kernel string, ostype string) {
	info.table.GetCell(row2, 2).SetText(hostname)
	info.table.GetCell(row3, 2).SetText(ostype)
	info.table.GetCell(row1, 5).SetText(kernel)
}

// UpdateSystemUsageInfo updates memory and swap values
func (info *InfoBar) UpdateSystemUsageInfo(memUsage float64, swapUsage float64) {
	memUsageText := progressUsageString(memUsage)
	swapUsageText := progressUsageString(swapUsage)
	info.table.GetCell(memCellRow, 2).SetText(memUsageText)
	info.table.GetCell(swapCellRow, 2).SetText(swapUsageText)
}

// UpdateConnStatus updates connection status value
func (info *InfoBar) UpdateConnStatus(status bool) {
	info.connOK = status
	connStatus := ""
	if info.connOK {
		connStatus = fmt.Sprintf("%s STATUS_OK", connOK)

	} else {
		connStatus = fmt.Sprintf("%s STATUS_ERR", connERR)
	}
	info.table.GetCell(row1, 2).SetText(connStatus)
}

// Draw draws this primitive onto the screen.
func (info *InfoBar) Draw(screen tcell.Screen) {
	info.Box.DrawForSubclass(screen, info)
	info.Box.SetBorder(false)
	x, y, width, height := info.GetInnerRect()
	info.table.SetRect(x, y, width, height)
	info.table.SetBorder(false)
	info.table.Draw(screen)
}

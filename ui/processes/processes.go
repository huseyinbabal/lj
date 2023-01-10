package processes

import (
	"fmt"
	"github.com/huseyinbabal/lj/internal/jvm"
	"github.com/huseyinbabal/lj/ui/dialogs"
	"github.com/huseyinbabal/lj/ui/utils"
	"github.com/rivo/tview"
	"strings"
	"sync"
)

type Processes struct {
	*tview.Box
	title           string
	headers         []string
	table           *tview.Table
	errorDialog     *dialogs.ErrorDialog
	cmdDialog       *dialogs.CommandDialog
	cmdInputDialog  *dialogs.SimpleInputDialog
	confirmDialog   *dialogs.ConfirmDialog
	messageDialog   *dialogs.MessageDialog
	progressDialog  *dialogs.ProgressDialog
	topDialog       *dialogs.TopDialog
	processes       instancesReport
	selectedID      string
	selectedProcess string
	confirmData     string
	jvm             *jvm.Jvm
}

type instancesReport struct {
	mu     sync.Mutex
	report []jvm.Process
}

func NewProcesses(jvm *jvm.Jvm) *Processes {
	processes := &Processes{
		Box:            tview.NewBox(),
		title:          "processes",
		headers:        []string{"ID", "Name"},
		errorDialog:    dialogs.NewErrorDialog(),
		cmdInputDialog: dialogs.NewSimpleInputDialog(""),
		messageDialog:  dialogs.NewMessageDialog(""),
		progressDialog: dialogs.NewProgressDialog(),
		topDialog:      dialogs.NewTopDialog(),
		confirmDialog:  dialogs.NewConfirmDialog(),
		jvm:            jvm,
	}
	processes.topDialog.SetTitle("JVM Processes")
	processes.cmdDialog = dialogs.NewCommandDialog([][]string{
		{"openInJConsole", "Open process in JConsole"},
		{"kill", "Kill JVM process"},
	})

	fgColor := utils.Styles.PageTable.FgColor
	bgColor := utils.Styles.PageTable.BgColor
	processes.table = tview.NewTable()
	processes.table.SetTitle(fmt.Sprintf("[::b]%s[0]", strings.ToUpper(processes.title)))
	processes.table.SetBorderColor(bgColor)
	processes.table.SetTitleColor(fgColor)
	processes.table.SetBorder(true)
	fgColor = utils.Styles.PageTable.HeaderRow.FgColor
	bgColor = utils.Styles.PageTable.HeaderRow.BgColor

	for i := 0; i < len(processes.headers); i++ {
		processes.table.SetCell(0, i,
			tview.NewTableCell(fmt.Sprintf("[black::b]%s", strings.ToUpper(processes.headers[i]))).
				SetExpansion(1).
				SetBackgroundColor(bgColor).
				SetTextColor(fgColor).
				SetAlign(tview.AlignLeft).
				SetSelectable(false))
	}

	processes.table.SetFixed(1, 1)
	processes.table.SetSelectable(true, false)

	// set command dialog functions
	processes.cmdDialog.SetSelectedFunc(func() {
		processes.cmdDialog.Hide()
		processes.runCommand(processes.cmdDialog.GetSelectedItem())
	})
	processes.cmdDialog.SetCancelFunc(func() {
		processes.cmdDialog.Hide()
	})
	// set input cmd dialog functions
	processes.cmdInputDialog.SetCancelFunc(func() {
		processes.cmdInputDialog.Hide()
	})

	processes.cmdInputDialog.SetSelectedFunc(func() {
		processes.cmdInputDialog.Hide()
	})
	// set message dialog functions
	processes.messageDialog.SetSelectedFunc(func() {
		processes.messageDialog.Hide()
	})
	processes.messageDialog.SetCancelFunc(func() {
		processes.messageDialog.Hide()
	})

	// set container top dialog functions
	processes.topDialog.SetDoneFunc(func() {
		processes.topDialog.Hide()
	})

	processes.confirmDialog.SetCancelFunc(func() {
		processes.confirmDialog.Hide()
	})

	// set confirm dialogs functions
	processes.confirmDialog.SetSelectedFunc(func() {
		processes.confirmDialog.Hide()
		switch processes.confirmData {
		case "openInJConsole":
			processes.openInJConsole()
		}
	})
	processes.confirmDialog.SetCancelFunc(processes.confirmDialog.Hide)
	return processes
}

// GetTitle returns primitive title
func (i *Processes) GetTitle() string {
	return i.title
}

// HasFocus returns whether or not this primitive has focus
func (i *Processes) HasFocus() bool {
	if i.table.HasFocus() || i.errorDialog.HasFocus() {
		return true
	}
	if i.cmdDialog.HasFocus() || i.progressDialog.HasFocus() {
		return true
	}
	if i.topDialog.HasFocus() || i.messageDialog.IsDisplay() {
		return true
	}
	if i.confirmDialog.HasFocus() || i.cmdInputDialog.IsDisplay() {
		return true
	}
	return i.Box.HasFocus()
}

// Focus is called when this primitive receives focus
func (i *Processes) Focus(delegate func(p tview.Primitive)) {
	// error dialog
	if i.errorDialog.IsDisplay() {
		delegate(i.errorDialog)
		return
	}
	// command dialog
	if i.cmdDialog.IsDisplay() {
		delegate(i.cmdDialog)
		return
	}
	// command input dialog
	if i.cmdInputDialog.IsDisplay() {
		delegate(i.cmdInputDialog)
		return
	}
	// message dialog
	if i.messageDialog.IsDisplay() {
		delegate(i.messageDialog)
		return
	}
	// container top dialog
	if i.topDialog.IsDisplay() {
		delegate(i.topDialog)
		return
	}
	// confirm dialog
	if i.confirmDialog.IsDisplay() {
		delegate(i.confirmDialog)
		return
	}
	delegate(i.table)
}

func (i *Processes) getSelectedItem() (string, string) {
	var instanceId string
	var instanceType string
	if i.table.GetRowCount() <= 1 {
		return instanceId, instanceType
	}
	row, _ := i.table.GetSelection()
	instanceId = i.table.GetCell(row, 0).Text
	instanceType = i.table.GetCell(row, 1).Text
	return instanceId, instanceType
}

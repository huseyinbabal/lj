package processes

import "fmt"

func (i *Processes) runCommand(cmd string) {
	switch cmd {
	case "openInJConsole":
		i.openInJConsole()
	case "kill":
		i.kill()

	}
}

func (i *Processes) openInJConsole() {
	if i.selectedID == "" {
		i.errorDialog.SetText("there is no jvm process to openInJConsole")
		i.errorDialog.Display()
		return
	}
	err := i.jvm.OpenInJConsole(i.selectedID)
	if err != nil {
		i.errorDialog.SetText(fmt.Sprintf("Failed to open process with id %s in jconsole. %v", i.selectedID, err))
		i.errorDialog.Display()
		return

	}
}
func (i *Processes) kill() {
	if i.selectedID == "" {
		i.errorDialog.SetText("there is no jvm process to kill")
		i.errorDialog.Display()
		return
	}
	err := i.jvm.KillProcess(i.selectedID)
	if err != nil {
		i.errorDialog.SetText(fmt.Sprintf("Failed to kill process with id %s. %v", i.selectedID, err))
		i.errorDialog.Display()
		return

	}
}

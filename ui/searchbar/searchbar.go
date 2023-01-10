package searchbar

import (
	"github.com/gdamore/tcell/v2"
	"github.com/huseyinbabal/lj/ui/utils"
	"github.com/kyokomi/emoji/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"strings"
)

// SearchBar implements the search bar primitive
type SearchBar struct {
	*tview.InputField
	title string
}

// NewSearchBar returns search bar view
func NewSearchBar(f func(key tcell.Key)) *SearchBar {
	inputField := tview.NewInputField().
		SetLabel(emoji.Sprint(":mag_right:")).
		SetFieldWidth(100).
		SetDoneFunc(f)

	inputField.SetAutocompleteFunc(func(currentText string) (entries []string) {
		if len(currentText) == 0 {
			return
		}
		for _, word := range utils.JvmResources() {
			log.Info().Msgf("currtext:%s", currentText)
			if strings.HasPrefix(strings.ToLower(word), strings.ToLower(currentText)) {
				entries = append(entries, word)
			}
		}
		log.Info().Msgf("words: %s", strings.Join(entries, "-"))
		if len(entries) <= 0 {
			entries = nil
		}
		return
	})

	// searchBar
	searchBar := &SearchBar{
		title:      "searchbar",
		InputField: inputField,
	}
	return searchBar
}

// Draw draws this primitive onto the screen.
func (search *SearchBar) Draw(screen tcell.Screen) {
	search.InputField.Draw(screen)
}

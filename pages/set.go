package pages

import "fyne.io/fyne/v2"

func SetPage(win *fyne.Window, page string) {

	content := (*win).Content().(*fyne.Container)

	switch page {
	case "start":
		content.Objects[1] = Start(win)
	case "dashboard":
		content.Objects[1] = Dashboard(win)

	}

}

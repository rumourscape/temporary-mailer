package pages

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/rumourscape/temporary-mailer/mailer"
)

func Dashboard(win *fyne.Window) fyne.CanvasObject {

	title := canvas.NewText("Messages", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 30

	address := widget.NewButtonWithIcon(mailer.GetAddress(), theme.ContentCopyIcon(), func() {
		clip := (*win).Clipboard()
		clip.SetContent(mailer.GetAddress())
	})
	address.IconPlacement = widget.ButtonIconTrailingText

	password := widget.NewPasswordEntry()
	password.Text = mailer.GetPassword()
	password.OnChanged = func(s string) { password.SetText(mailer.GetPassword()) }
	//password.Disable()

	var split *container.Split
	mesCon := container.NewVBox()
	setMesCon(mesCon)

	del := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() { mailer.DeleteAccount(); SetPage(win, "start") })
	del.Importance = widget.DangerImportance

	refresh := widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), func() { setMesCon(mesCon) })

	logout := widget.NewButtonWithIcon("Logout", theme.LogoutIcon(), func() { mailer.Logout(); SetPage(win, "start") })
	logout.Importance = widget.WarningImportance

	trailer := container.NewHBox(container.NewCenter(refresh), container.NewCenter(logout), container.NewCenter(del))

	vCon := container.NewVBox(address, password)

	header := container.NewAdaptiveGrid(3, container.NewPadded(vCon), title, container.NewPadded(trailer))

	split = container.NewVSplit(header, container.NewVScroll(mesCon))
	split.Offset = 0.1

	return split
}

func setMesCon(mesCon *fyne.Container) {
	log.Println("Getting messages...")

	page := 1
	messages, err := mailer.GetMessages(page)
	if err != nil {
		log.Println(err)
	}

	if len(messages) == 0 {
		emptyLabel := canvas.NewText("No messages found", color.White)
		emptyLabel.Alignment = fyne.TextAlignCenter
		emptyLabel.TextStyle.Bold = true
		emptyLabel.TextSize = 20

		mesCon.RemoveAll()
		mesCon.Add(layout.NewSpacer())
		mesCon.Add(emptyLabel)
		mesCon.Add(layout.NewSpacer())
	} else {
		mesCon.RemoveAll()
		for _, message := range messages {
			card := widget.NewCard(
				message.Subject,
				message.From.Address+"\t"+message.CreatedAt.Format("2006-01-02 15:04:05"),
				widget.NewLabel(mailer.GetText(message.ID)),
			)
			mesCon.Add(container.NewCenter(card))
		}
	}
}

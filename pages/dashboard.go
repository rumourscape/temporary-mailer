package pages

import (
	"encoding/json"
	"image/color"
	"log"
	"net/http"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/donovanhide/eventsource"
	"github.com/felixstrobel/mailtm"

	"github.com/rumourscape/temporary-mailer/mailer"
)

func Dashboard(win *fyne.Window) fyne.CanvasObject {

	title := canvas.NewText("Inbox", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 20

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

	go sse(mesCon)

	popup := DeletePopUp(win)

	del := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() { popup.Show() })
	del.Importance = widget.DangerImportance

	refresh := widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), func() { setMesCon(mesCon) })

	logout := widget.NewButtonWithIcon("Logout", theme.LogoutIcon(), func() { mailer.Logout(); SetPage(win, "start") })
	logout.Importance = widget.WarningImportance

	trailer := container.NewHBox(container.NewCenter(refresh), container.NewCenter(logout), container.NewCenter(del))

	vCon := container.NewVBox(address, password)

	header := container.NewAdaptiveGrid(3, container.NewPadded(vCon), title, container.NewPadded(trailer), popup)

	split = container.NewVSplit(header, container.NewVScroll(mesCon))
	split.Offset = 0.1

	return split
}

func sse(mesCon *fyne.Container) {
	accountId := mailer.GetId()
	sseUrl, err := url.Parse("https://mercure.mail.tm/.well-known/mercure")
	if err != nil {
		log.Println(err)
		return
	}
	defTopic := "/accounts/{id}"
	topic := "/accounts/" + accountId

	params := url.Values{}
	params.Add("topic", defTopic)
	params.Add("topic", topic)
	sseUrl.RawQuery = params.Encode()

	request, err := http.NewRequest("GET", sseUrl.String(), nil)
	request.Header.Add("Authorization", "Bearer "+mailer.GetToken())

	if err != nil {
		log.Println(err)
		return
	}

	es, err := eventsource.SubscribeWithRequest("", request)
	if err != nil {
		log.Println("SSE ERR:", err, sseUrl.String())
		return
	}

	for {
		select {
		case c := <-es.Events:
			var message mailtm.Message
			err = json.Unmarshal([]byte(c.Data()), &message)
			if err != nil {
				log.Println(err)
			}

			if mesCon.Objects[1] == layout.NewSpacer() {
				mesCon.RemoveAll()
				setMesCon(mesCon)
			}

			card := widget.NewCard(
				message.Subject,
				message.From.Address+"\t"+message.CreatedAt.Format("2006-01-02 15:04:05"),
				widget.NewLabel(mailer.GetText(message.ID)),
			)
			//Insert card at the top
			copy(mesCon.Objects[1:], mesCon.Objects)
			mesCon.Objects[0] = card

		case err := <-es.Errors:
			log.Println(err)
		}
	}
}

func setMesCon(mesCon *fyne.Container) {
	log.Println("Fetching messages...")

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
			mesCon.Add(card)
		}
	}
}

func DeletePopUp(win *fyne.Window) *widget.PopUp {
	popup := widget.NewModalPopUp(container.NewVBox(), (*win).Canvas())
	popup.Resize(fyne.NewSize(400, 50))
	popup.Move(fyne.NewPos(200, 400))
	popup.Hide()

	label := widget.NewLabel("Are you sure you want to delete your account?")
	label.Alignment = fyne.TextAlignCenter

	confirm := widget.NewButton("YES", func() { go mailer.DeleteAccount(); popup.Hide(); SetPage(win, "start") })
	cancel := widget.NewButton("NO", func() { popup.Hide() })

	hCon := container.NewHBox(confirm, cancel)

	popup.Content.(*fyne.Container).Add(label)
	popup.Content.(*fyne.Container).Add(container.NewCenter(hCon))

	return popup
}

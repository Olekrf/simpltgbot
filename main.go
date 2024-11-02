package main
import("fmt"
 "os" 
 "github.com/mymmrac/telego" 
 tu "github.com/mymmrac/telego/telegoutil")
 func main(){
	botToken := ""
	bot, er := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	}
	updates, _ := bot.UpdatesViaLongPolling(nil)
	defer bot.StopLongPolling()
	for update :=range updates{
		if update.Message !=nil {
			chatID := tu.ID(update.Message.Chat.ID)

			_, _ = bot.CopyMessage(
				tu.CopyMessage(
					chatID,
					chatID,
					update.Message.MessageID,
				),
			)
		}
	}
 }
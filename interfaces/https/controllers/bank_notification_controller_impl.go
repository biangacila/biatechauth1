package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Define the structure of the notification
type Notification struct {
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	PackageName string    `json:"packageName"`
	Timestamp   time.Time `json:"timestamp"`
}

type BankNotificationControllerImpl struct {
}

func NewBankNotificationControllerImpl() *BankNotificationControllerImpl {
	return &BankNotificationControllerImpl{}
}

func (b BankNotificationControllerImpl) ReceiveNotificationGet(w http.ResponseWriter, r *http.Request) {

}

func (b BankNotificationControllerImpl) ReceiveNotificationPost(w http.ResponseWriter, r *http.Request) {
	var notification Notification

	// Parse the JSON body into the Notification struct
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Print the received notification
	fmt.Printf("\n\n-:) Received notification:\nTitle: %s\nText: %s\nPackage: %s\nTimestamp: %s\n",
		notification.Title, notification.Text, notification.PackageName, notification.Timestamp)

	// Send a response back to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Notification received successfully")
}

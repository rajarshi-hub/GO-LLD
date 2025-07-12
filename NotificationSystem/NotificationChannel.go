package main

import "fmt"

type NotificationChannel interface {
	sendNotification(userIds []string, message string) bool
	getRetryCount() int
}

const WHATSAPP_CHANNEL = "WHATSAPP_CHANNEL"
const PUSHNOTIFICATION_CHANNEL = "PUSHNOTIFICATION_CHANNEL"

type WhatsappChannel struct {
	sendGridToken  string
	whatsappClient string
	retryCount     int
}

func (wc *WhatsappChannel) sendNotification(userIds []string, message string) bool {
	// Use the whatsappClient Third party to send notification
	fmt.Printf("Message %v sent to whatsapp\n", message)
	return true
}
func (wc *WhatsappChannel) getRetryCount() int {
	return wc.retryCount
}

type PushNotificationChannel struct {
	apnToken   string
	apnClient  string
	retryCount int
}

func (pn *PushNotificationChannel) sendNotification(userIds []string, message string) bool {
	// sending message via APN
	fmt.Printf("Message %v not sent to Apple Push Notif due to congestion\n", message)
	return false
}

func (pn *PushNotificationChannel) getRetryCount() int {
	return pn.retryCount
}

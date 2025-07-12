package main

import "time"

type SendNotification struct {
	typeToChannel map[string]NotificationChannel
}

func (sn *SendNotification) initialize() {
	sn.typeToChannel = make(map[string]NotificationChannel)
	sn.typeToChannel[WHATSAPP_CHANNEL] = &WhatsappChannel{retryCount: 2}
	sn.typeToChannel[PUSHNOTIFICATION_CHANNEL] = &PushNotificationChannel{retryCount: 3}
}

func (sn *SendNotification) SendNotification(userIds []string, channels []string, mess string) {
	for _, channelName := range channels {
		curRetry := 0
		if channel, ok := sn.typeToChannel[channelName]; ok {
			res := false
			for !res && curRetry < channel.getRetryCount() {
				res = channel.sendNotification(userIds, mess)
				curRetry++
				time.Sleep(1000 * time.Millisecond)
			}
		}
	}
}

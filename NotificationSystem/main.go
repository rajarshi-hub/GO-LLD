package main

import "fmt"

func main() {
	notifSystem := SendNotification{}
	notifSystem.initialize()
	userSegmentSystem := userSegment{}
	_, err := userSegmentSystem.createStaticSegment([]string{"123"})
	if err != nil {
		fmt.Println(err)
	}
	userIds, err := userSegmentSystem.getUsersFromSegment("123")
	if err != nil {
		fmt.Println(err)
	}
	notifSystem.SendNotification(userIds, []string{WHATSAPP_CHANNEL}, "Happy Message")
	notifSystem.SendNotification(userIds, []string{PUSHNOTIFICATION_CHANNEL, WHATSAPP_CHANNEL}, "Happy Message")
	notifSystem.SendNotification(userIds, []string{PUSHNOTIFICATION_CHANNEL}, "Happy Message")

}

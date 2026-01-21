package helpers


import (
	"strings"
	"bytes"
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"
	// "crypto/rsa"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jolotech/jolo-mars/types"

	"fmt"
)

func GetNotificationStatusData(db *gorm.DB, userType string, key string, notificationType string, storeID *uint,) bool {

	var setting models.NotificationSetting

	query := db.Where("type = ?", userType).Where("key = ?", key)

	// Store-specific override
	if storeID != nil && userType == "store" {
		query = query.Where("store_id = ?", *storeID)
	}

	if err := query.First(&setting).Error; err != nil {
		return false
	}

	var status string

	switch notificationType {
	case "push_notification_status":
		status = setting.PushNotificationStatus
	case "email_notification_status":
		status = setting.EmailNotificationStatus
	case "sms_notification_status":
		status = setting.SmsNotificationStatus
	default:
		return false
	}

	return strings.ToLower(status) == "active"
}



// func SendPushNotifToDevice(fcmToken string, data map[string]interface{},webPushLink ...string,) error {

// 	clickAction := ""
// 	if len(webPushLink) > 0 {
// 		clickAction = webPushLink[0]
// 	}

// 	payload := map[string]interface{}{
// 		"message": map[string]interface{}{
// 			"token": fcmToken,
// 			"data": map[string]string{
// 				"title":       toString(data["title"]),
// 				"body":        toString(data["description"]),
// 				"image":       toString(data["image"]),
// 				"order_id":    toString(data["order_id"]),
// 				"type":        toString(data["type"]),
// 				"click_action": clickAction,
// 				"sound":       "notification.wav",
// 			},
// 			"notification": map[string]string{
// 				"title": toString(data["title"]),
// 				"body":  toString(data["description"]),
// 				"image": toString(data["image"]),
// 			},
// 			"android": map[string]interface{}{
// 				"notification": map[string]string{
// 					"channelId": "6ammart",
// 				},
// 			},
// 			"apns": map[string]interface{}{
// 				"payload": map[string]interface{}{
// 					"aps": map[string]string{
// 						"sound": "notification.wav",
// 					},
// 				},
// 			},
// 		},
// 	}

// 	body, _ := json.Marshal(payload)

// 	req, _ := http.NewRequest(
// 		"POST",
// 		"https://fcm.googleapis.com/v1/projects/YOUR_PROJECT_ID/messages:send",
// 		bytes.NewBuffer(body),
// 	)

// 	req.Header.Set("Authorization", "Bearer YOUR_FCM_SERVER_KEY")
// 	req.Header.Set("Content-Type", "application/json")

// 	_, err := http.DefaultClient.Do(req)
// 	return err
// }

// func toString(v interface{}) string {
// 	if v == nil {
// 		return ""
// 	}
// 	return fmt.Sprintf("%v", v)
// }



package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func SendPushNotifToDevice(
	sa FirebaseServiceAccount,
	fcmToken string,
	data map[string]interface{},
	webPushLink ...string,
) error {

	clickAction := ""
	if len(webPushLink) > 0 {
		clickAction = webPushLink[0]
	}

	accessToken, err := getAccessToken(sa)
	if err != nil {
		return err
	}

	payload := buildFCMPayload(fcmToken, data, clickAction)
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		"https://fcm.googleapis.com/v1/projects/"+sa.ProjectID+"/messages:send",
		bytes.NewBuffer(body),
	)

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	return err
}

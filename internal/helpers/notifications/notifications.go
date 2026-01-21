package helpers


import (
	"strings"
	"bytes"
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
	"github.com/jolotech/jolo-mars/internal/models"
	// "github.com/jolotech/jolo-mars/types"
)




func GetNotificationStatusData(userType string, key string, notificationType string, storeID *uint,) bool {

	var db *gorm.DB
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



func SendPushNotifToDevice(fcmToken string, data map[string]interface{}, webPushLink ...string,) error {

	sa, err := LoadFirebaseServiceAccountFromEnv()
	if err != nil {
		return err
	}

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

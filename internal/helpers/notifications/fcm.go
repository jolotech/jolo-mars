package helpers

import (
	"fmt"
	"os"
	"strings"
	"errors"
	"github.com/jolotech/jolo-mars/types"
)


func buildFCMPayload(fcmToken string, data map[string]interface{}, clickAction string,) map[string]interface{} {

	return map[string]interface{}{
		"message": map[string]interface{}{
			"token": fcmToken,
			"data": map[string]string{
				"title":       toString(data["title"]),
				"body":        toString(data["description"]),
				"image":       toString(data["image"]),
				"order_id":    toString(data["order_id"]),
				"type":        toString(data["type"]),
				"click_action": clickAction,
				"sound":       "notification.wav",
			},
			"notification": map[string]string{
				"title": toString(data["title"]),
				"body":  toString(data["description"]),
				"image": toString(data["image"]),
			},
			"android": map[string]interface{}{
				"notification": map[string]string{
					"channelId": "6ammart",
				},
			},
			"apns": map[string]interface{}{
				"payload": map[string]interface{}{
					"aps": map[string]string{
						"sound": "notification.wav",
					},
				},
			},
		},
	}


	// return map[string]interface{}{
	// 	"message": map[string]interface{}{
	// 		"token": fcmToken,
	// 		"data": map[string]string{
	// 			"title": toString(data["title"]),
	// 			"body":  toString(data["description"]),
	// 			"type":  toString(data["type"]),
	// 			"sound": "notification.wav",
	// 		},
	// 		"notification": map[string]string{
	// 			"title": toString(data["title"]),
	// 			"body":  toString(data["description"]),
	// 		},
	// 		"android": map[string]interface{}{
	// 			"notification": map[string]string{
	// 				"channelId": "6ammart",
	// 			},
	// 		},
			
	// 	},
	// }
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}



func LoadFirebaseServiceAccountFromEnv() (*types.FirebaseServiceAccount, error) {
	projectID := os.Getenv("FIREBASE_PROJECT_ID")
	clientEmail := os.Getenv("FIREBASE_CLIENT_EMAIL")
	privateKey := os.Getenv("FIREBASE_PRIVATE_KEY")

	if projectID == "" || clientEmail == "" || privateKey == "" {
		return nil, errors.New("firebase env variables not set")
	}

	// Convert \n to actual newlines (CRITICAL)
	privateKey = strings.ReplaceAll(privateKey, `\n`, "\n")

	return &types.FirebaseServiceAccount{
		ProjectID:   projectID,
		ClientEmail: clientEmail,
		PrivateKey:  privateKey,
	}, nil
}

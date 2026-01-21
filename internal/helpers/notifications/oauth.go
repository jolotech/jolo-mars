package helpers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"github.com/jolotech/jolo-mars/types"
	// "strings"
)

func getAccessToken(sa *types.FirebaseServiceAccount) (string, error) {
	jwtToken, err := createJWT(sa)
	if err != nil {
		return "", err
	}

	resp, err := http.PostForm(
		"https://oauth2.googleapis.com/token",
		url.Values{
			"grant_type": {"urn:ietf:params:oauth:grant-type:jwt-bearer"},
			"assertion":  {jwtToken},
		},
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		AccessToken string `json:"access_token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result.AccessToken, err
}

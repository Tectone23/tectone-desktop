package zoho

import (
	"fmt"
	"net/http"
	"strings"
	"tectone/tectone-desktop/internal/model"
)

const zohoApiUrl = "https://accounts.zoho.com/oauth/v2/token"

var accessTokenQueryParamsFormat = "?code=%s&grant_type=authorization_code&client_id=%s&client_secret=%s&redirect_uri=https://www.tectone23.com/"
var refreshTokenQueryParamsFormat = "?refresh_token=%s&client_id=%s&client_secret=%s&scope=%s&redirect_uri=https://www.tectone23.com/&grant_type=refresh_token"

type Zoho struct {
	clientId          string
	clientSecret      string
	authorizationCode string
	departmentId      string
	refreshToken      string
	scope             string
}

func (z *Zoho) getRefreshToken() {}

func (z *Zoho) SubmitFeedback(f *model.Feedback) error {
	url := z.generateAccessTokenURL()
	_, err := http.Post(url, "application/json", nil)

	return err
}

func (z *Zoho) generateRefreshTokenURL() string {
	rtformat := fmt.Sprintf(refreshTokenQueryParamsFormat, z.refreshToken, z.clientId, z.clientSecret, z.scope)
	return strings.Join([]string{zohoApiUrl, rtformat}, "")

}

func (z *Zoho) generateAccessTokenURL() string {
	atformat := fmt.Sprintf(accessTokenQueryParamsFormat, z.authorizationCode, z.clientId, z.clientSecret)
	return strings.Join([]string{zohoApiUrl, atformat}, "")
}

func New() *Zoho {

	return &Zoho{
		clientId:          "1000.PWUUWBSI1Q3ND0UTW8EQBO3EIDVBPI",
		clientSecret:      "298aaa51f3c55cb36ca0037716c6309e4ec7688f1f",
		authorizationCode: "",
		departmentId:      "164109000000251406",
		refreshToken:      "",
		scope:             "Desk.tickets.CREATE",
	}

}

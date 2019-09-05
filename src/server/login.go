package server

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type LoginResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func Login(keys *JWTKeys) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.FormValue("userId")
		password := c.FormValue("password")

		if len(userId) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "ログイン名が入力されていません。")
		}
		if len(password) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "パスワードが入力されていません。")
		}

		// 本来はここで、データベースに格納されたユーザアカウントとパスワードをチェックする。
		// デモではID:test, password:aaaaaaaaとする。
		if userId != "test" || password != "aaaaaaaa" {
			return echo.NewHTTPError(http.StatusUnauthorized, "ログインに失敗しました。")
		}

		now := time.Now()

		accessToken, err := keys.NewAccessToken(userId, now)
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "ログインに失敗しました。")
		}

		refreshToken, err := keys.NewRefreshToken(userId, now)
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "ログインに失敗しました。")
		}

		// ここでtokenを発行した時刻をデータベースなどに記録しておくと
		// APIリクエストの際にその時刻より古いtokenをエラーにすることで二重ログインの防止ができる

		response := LoginResponseData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		err = c.JSON(http.StatusOK, response)
		return err
	}
}

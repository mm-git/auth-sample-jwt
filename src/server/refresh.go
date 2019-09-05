package server

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type RefreshResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func Refresh(keys *JWTKeys) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := GetCustomClaims(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "bad token")
		}

		if claims.Subject != RefreshToken {
			return echo.NewHTTPError(http.StatusUnauthorized, "not refresh token")
		}

		// ここでtokenに含まれるuserIdが本当にデータベースに存在するか念のためチェックする。
		// またデータベースに記録されたtoken発行時刻よりRefreshTokenの時刻の方が古い場合、
		// RefreshTokenの使用を1回のみに制限するのであれば、ここでエラーにする。。

		userId := claims.UserId
		now := time.Now()
		accessToken, err := keys.NewAccessToken(userId, now)
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "failed to refresh token")
		}

		refreshToken, err := keys.NewRefreshToken(userId, now)
		if err != nil {
			return echo.NewHTTPError(http.StatusServiceUnavailable, "failed to refresh token")
		}

		// ここでtokenを再発行した時刻をデータベースなどに記録する

		response := RefreshResponseData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		err = c.JSON(http.StatusOK, response)
		return err
	}
}

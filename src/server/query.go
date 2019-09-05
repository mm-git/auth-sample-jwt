package server

import (
	"auth-sample-jwt/src/graph"
	"github.com/99designs/gqlgen/handler"
	"github.com/labstack/echo"
	"net/http"
)

func Query() echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := GetCustomClaims(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "bad token")
		}

		if claims.Subject != AccessToken {
			return echo.NewHTTPError(http.StatusUnauthorized, "not access token")
		}

		// ここでtokenに含まれるuserIdが本当にデータベースに存在するか念のためチェックする。
		// またデータベースに記録されたtoken発行時刻よりAccessTokenの時刻の方が古い場合、二重ログインされている。
		// 二重ログインを禁止する場合はここでエラーにする

		// gqlgenの各クエリーの内部で使用するデータ(アカウント名、データベースインスタンス)をここでセットして渡す。
		// userId := claims.UserId
		config := graph.Config{
			Resolvers: &graph.Resolver{
				// Account: userId,
				// DB: db
			},
		}

		graphQLHandler := handler.GraphQL(graph.NewExecutableSchema(config))
		graphQLHandler(c.Response(), c.Request())

		return nil
	}
}

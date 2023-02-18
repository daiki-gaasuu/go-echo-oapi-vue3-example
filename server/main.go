package main

import (
	"net/http"
	oapi "server/generated"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type apiController struct{}

// OpenAPI で定義された (GET /hello) の実装
func (a apiController) GetGreeting(ctx echo.Context, params oapi.GetGreetingParams) error {
	payload := "hello go"
	if params.Name != nil {
		payload = *params.Name
	}
	return ctx.JSON(http.StatusOK, payload)
}

func main() {
	// Echo のインスタンス作成
	e := echo.New()

	// OpenApi 仕様に沿ったリクエストかバリデーションをするミドルウェアを設定
	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}
	e.Use(oapiMiddleware.OapiRequestValidator(swagger))
	// ロガーのミドルウェアを設定
	e.Use(middleware.Logger())
	// APIがエラーで落ちてもリカバーするミドルウェアを設定
	e.Use(middleware.Recover())

	// OpenAPI の仕様を満たす構造体をハンドラーとして登録する
	api := apiController{}
	oapi.RegisterHandlers(e, api)

	// 3000ポートで Echo サーバ起動
	e.Logger.Fatal(e.Start(":3000"))
}

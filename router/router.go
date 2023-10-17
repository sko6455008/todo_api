// APIを受け取りmodelに渡す
package router

import (
	"os"

	"github.com/labstack/echo/v4/middleware"

	_ "net/http"

	"github.com/labstack/echo/v4"
)

// EchoはAPIを受け取る箇所を手軽にしてくれるライブラリ
// Routingを設定する関数　引数はecho.echo型であり、戻り値はerror型
func SetRouter(e *echo.Echo) error {
	// APIが叩かれた時にログを出す
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	// 予想外のエラーが発生した際でも、サーバーを落とさないようにする
	e.Use(middleware.Recover())
	// CORSに対応する
	e.Use(middleware.CORS())

	// APIを書く場所
	api := e.Group("/api")
	{
		apiTasks := api.Group("/tasks")
		{
			apiTasks.GET("", GetTasksHandler)
			apiTasks.POST("", AddTaskHandler)
			apiTasks.PUT("/:taskID", ChangeFinishedTaskHandler)
			apiTasks.DELETE("/:taskID", DeleteTaskHandler)
		}
	}

	// 8000番のポートを開く
	err := e.Start(":8000")
	return err
}

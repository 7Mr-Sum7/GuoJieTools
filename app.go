package main

import (
	"GuoJieTools/gomod"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// RespDate 数据响应结构体
type RespDate struct {
	Code int          `json:"code"` // 响应码
	Data any `json:"data"` // 响应数据
	Msg  string       `json:"msg"`  // 响应消息
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Lifestyle(address string) RespDate {
	response := gomod.Lifestyle(address)

	return RespDate{Code: 200, Data: response, Msg: "success"}
}

func (a *App) GET(address string) any {
	url := address
	StatusCode := gomod.Method_GET(url)
	return StatusCode
}

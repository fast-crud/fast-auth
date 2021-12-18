package main

import "github.com/fast-crud/fast-auth/boot"

// @title GF-VUE-ADMIN Swagger Docs
// @version 1.0.1
// @description This is a GF-VUE-ADMIN Server
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	boot.Initialize()
	boot.Server.Initialize() // 初始化gf服务器
}

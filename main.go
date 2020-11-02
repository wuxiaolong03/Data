package main

import (
	"DataCertPlatform/blockchain"
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
)

func main() {

	block0 := blockchain.CreateGenesisBlock()
	block1 := blockchain.NewBlock(block0+1,block0.Hasha,[]byte("a"))
	fmt.println(block0,block1)
	return

	//连接数据库
	db_mysql.Connect()
	//静态资源文件映射。
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","static/css")
	beego.SetStaticPath("/img","static/img")

	beego.Run()//阻塞


}

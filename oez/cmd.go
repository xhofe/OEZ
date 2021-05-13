package oez

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	h bool
	i bool
	c string
)

func init() {
	flag.BoolVar(&h, "h", false, "show help")
	flag.BoolVar(&i, "i", false, "install application")
	flag.StringVar(&c, "c", "oez.yml", "config file")
}

func Run() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	if i {
		install()
		return
	}
	server()
}

func install() {
	if !InitConf(c) {
		return
	}
	log.Println("配置文件创建成功,请修改数据库信息.")
}

func printASC() {
	log.Print(`
 ________  _______   ________     
|\   __  \|\  ___ \ |\_____  \    
\ \  \|\  \ \   __/| \|___/  /|   
 \ \  \\\  \ \  \_|/__   /  / /   
  \ \  \\\  \ \  \_|\ \ /  /_/__  
   \ \_______\ \_______\\________\
    \|_______|\|_______|\|_______|

`)
}

func server() {
	printASC()
	if !ReadConf(c) {
		return
	}
	if !Init() {
		return
	}
	baseServer := "0.0.0.0:" + Config.Common.Listen
	log.Printf("Starting server @ %s", baseServer)
	if !Config.Common.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	InitGin(r)
	err := r.Run(baseServer)
	if err != nil {
		log.Print("Server failed start.\n" + err.Error())
	}
}

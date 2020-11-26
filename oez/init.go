package oez

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func InitConf(conf string) bool {
	if !Exists(conf) {
		confContent:=`common:
  debug: true
  listen: 5233
  title: oh,easy!
  chars: `+RandomStr(CHARS)+`
database:
  type: mysql
  user: oez
  password: 123456
  host: 127.0.0.1
  port: 3306
  name: oez
  tablePrefix: oez_`
		f,err:=CreatNestedFile(conf)
		if err !=nil {
			log.Printf("无法创建配置文件,%s",err.Error())
			return false
		}
		_,err=f.WriteString(confContent)
		if err !=nil {
			log.Printf("无法写入配置文件,%s",err.Error())
			return false
		}
		return true
	}
	log.Println("配置文件已存在，无需install")
	return false
}

var DB *gorm.DB

func Init() bool {
	log.Println("初始化数据库连接...")
	var (
		db *gorm.DB
		err error
	)
	switch Config.Database.Type {
	case "mysql":
		db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			Config.Database.User,
			Config.Database.Password,
			Config.Database.Host,
			Config.Database.Port,
			Config.Database.Name)),
			&gorm.Config{
			NamingStrategy:schema.NamingStrategy{
				TablePrefix: Config.Database.TablePrefix,
			},
		})
	default:
		log.Printf("不支持数据库类型: %s", Config.Database.Type)
	}
	if err !=nil {
		log.Printf("数据库连接失败,%s",err.Error())
		return false
	}
	DB=db

	log.Println("开始进行数据库初始化...")
	err = DB.AutoMigrate(&Url{})
	if err !=nil {
		log.Printf("数据库迁移失败,%s",err.Error())
		return false
	}
	return true
}

func InitGin(r *gin.Engine)  {
	if !Config.Common.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(cors.Default())
	r.LoadHTMLGlob("static/*")
	//r.StaticFile("/static/favicon.ico","static/favicon.ico")
	r.GET("/", Index)
	//r.GET("/static/favicon.ico",Favicon)
	r.POST("/",CreateUrl)
	r.GET("/:id",GetUrl)
	r.GET("/:id/json",GetUrlJSON)
}
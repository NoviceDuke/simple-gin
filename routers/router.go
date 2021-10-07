package routers

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"log"
	"nueip/eshop/controller"
	"nueip/eshop/pkg/setting"
	"nueip/eshop/repository"
	"nueip/eshop/service"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	CategoryController controller.CategoryController
	db                 *gorm.DB
)

func initDB() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func initHandler() {
	CategoryController = controller.CategoryController{
		CategorySrv: service.CategoryService{
			Repo: &repository.CategoryRepository{
				DB: db,
			},
		}}
}

func init() {
	initDB()
	initHandler()
}

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/category/:id", CategoryController.GetCategory)
		apiv1.PUT("/category/:id", CategoryController.EditCategory)
		apiv1.POST("/category", CategoryController.CreateCategory)
	}
	return r
}

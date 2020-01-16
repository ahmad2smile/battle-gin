package main

import (
	"fmt"
	"github.com/ahmad2smile/battle-gin/context"
	"github.com/ahmad2smile/battle-gin/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/battle_gin?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Db Error: ", err)
	}

	context.Setup(db)

	defer db.Close()

	router := gin.Default()

	controllers.GamesController(router)

	router.Run()
}

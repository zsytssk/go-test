package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 初始化 GORM 数据库连接
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()

	// 在中间件中设置上下文值
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// 处理路由
	r.GET("/", func(c *gin.Context) {
		// 从上下文中获取数据库连接
		db, ok := c.Value("db").(*gorm.DB)
		if !ok {
			c.JSON(500, gin.H{"error": "failed to get database connection"})
			return
		}
		// 这里可以使用 db 进行数据库操作
		c.JSON(200, gin.H{"message": "success"})
	})

	r.Run()
}

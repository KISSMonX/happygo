package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

const uri = "mongodb://con_log:@dds-.mongodb.rds.aliyuncs.com:3717,dds-.mongodb.rds..com:3717/con_log?replicaSet=mgset-"

func createCollectionWithSessionPool(name string) {
	session, err := mgo.Dial(uri)
	if err != nil {
		fmt.Println("连接MongoDB数据库失败", err.Error())
	}
	session.SetMode(mgo.Monotonic, true) // NOTE: 无法创建集合
	//default is 4096
	session.SetPoolLimit(100)

	collectionInfo := &mgo.CollectionInfo{
		DisableIdIndex:   false,
		ForceIdIndex:     false,
		Capped:           false,
		ValidationLevel:  "strict",
		ValidationAction: "error",
	}

	m := session.Copy()

	fmt.Println("用连接池, 新集合名:", name, "数据库:", "console_log")

	// 网络集合和索引
	err = m.DB("console_log").C(name).Create(collectionInfo)
	if err != nil {
		fmt.Println("用连接池创建集合失败: ", err.Error())
		// return
	}

}

func main() {
	createCollectionWithSessionPool("直接调用_使用连接池")

	r := gin.Default()
	r.GET("/new", func(c *gin.Context) {
		createCollectionWithSessionPool("接口调用_使用连接池")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

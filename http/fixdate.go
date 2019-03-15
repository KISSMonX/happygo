package main

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// globalSession 全局连接池
var globalSession *mgo.Session
var globalDBName string

// StarkMgo 服务 mgo 实例
type StarkMgo struct {
	Session *mgo.Session
	DBName  string
}

// Init 初始化连接信息
func Init() error {
	session, err := mgo.Dial("mongodb://console_log:xH3KyMxxfuU8gV93PtWAgffc9H7ADUXm@dds-bp123b105f4173e41.mongodb.rds.aliyuncs.com:3717,dds-bp123b105f4173e42.mongodb.rds.aliyuncs.com:3717/console_log?replicaSet=mgset-9488657")
	if err != nil {
		fmt.Println("连接MongoDB数据库失败", err.Error())
		return err
	}
	globalSession = session
	globalSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	globalSession.SetPoolLimit(4096)

	globalDBName = "console_log"

	return nil
}

// CloneMgoSession 获取数据库实例
func CloneMgoSession() *mgo.Session {
	return globalSession.Copy()
}

// New 新建实例
func New() *StarkMgo {
	mgo := &StarkMgo{}

	mgo.Session = globalSession.Copy()
	if mgo.Session == nil {
		fmt.Println("新建MongoDB实例失败")
		return nil
	}

	mgo.DBName = globalDBName

	return mgo
}

// Close 销毁实例
func (m *StarkMgo) Close() {
	if m.Session != nil {
		m.Session.Close()
	}
}

// GetSession 获取会话
func (m *StarkMgo) GetSession() *mgo.Session {
	return m.Session
}

// MgoDB 获取数据库
func (m *StarkMgo) MgoDB() *mgo.Database {
	return m.Session.DB(m.DBName)
}

func main() {

	if err := Init(); err != nil {
		panic(err)
	}
	m := New()
	defer m.Close()

	t := time.Now()

	err := m.MgoDB().C("roomlog_201811").Find(bson.M{"sessionId": bson.M{"$regex": "^201811"}, "operation": "Login"})
	fmt.Println("ERROR:", err, "日志总数:", len(logs), "耗时:", time.Since(t))

}

package main

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
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
	globalSession.SetPoolLimit(100)

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

type countListModel struct {
	RealUserCount  int       `json:"realUserCount" bson:"realUserCount"`
	FakeUserCount  int       `json:"fakeUserCount" bson:"fakeUserCount"`
	TotalUserCount int       `json:"totalUserCount" bson:"totalUserCount"`
	TimeStamp      time.Time `json:"timeStamp" bson:"timeStamp"`
}

type lobbyCountLatestInfoListMongoModel struct {
	SessionRoomID string         `json:"sessionRoomId" bson:"sessionRoomId"`
	SessionID     string         `json:"sessionId" bson:"sessionId"`
	Latest        countListModel `json:"latest" bson:"latest"`
	UpdateTime    time.Time      `json:"updateTime" bson:"updateTime"`
}

type lobbyCountRespModel struct {
	SessionRoomID  string    `json:"sessionRoomId"`
	SessionID      string    `json:"sessionId"`
	RealUserCount  int       `json:"realUserCount"`
	FakeUserCount  int       `json:"fakeUserCount"`
	TotalUserCount int       `json:"totalUserCount"`
	UpdateTime     time.Time `json:"updateTime"`
}

func main() {
	if err := Init(); err != nil {
		panic(err)
	}

	yyyymmddhh := ""

	m := New()
	defer m.Close()

	lobbyRoomsLatestInfo := []lobbyCountLatestInfoListMongoModel{}
	results := []lobbyCountRespModel{}
	var start, end time.Time
	var err error

	timeLoc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	if yyyymmddhh != "" {
		start, err = time.Parse("2006010215", yyyymmddhh)
		if err != nil {
			fmt.Println("Lobby=>解析输入日期格式失败:", yyyymmddhh, err)
			return
		}
		end = start.Add(time.Hour)
	} else {
		y, m, d := time.Now().Date()
		start = time.Date(y, m, d, 0, 0, 0, 0, timeLoc)
		end = start.Add(24 * time.Hour)
	}

	err = m.MgoDB().C("lobby_room_member").Find(bson.M{"updateTime": bson.M{"$gt": start, "$lt": end}}).
		Select(bson.M{"sessionRoomId": 1, "sessionId": 1, "updateTime": 1, "latest": bson.M{"$arrayElemAt": bson.M{"infoList": -1}}}).
		All(&lobbyRoomsLatestInfo)
	if err != nil {
		fmt.Println("Lobby=>获取指定时间课程人数失败:", err, start, end)
		return
	}

	for _, room := range lobbyRoomsLatestInfo {
		tmp := lobbyCountRespModel{
			SessionRoomID:  room.SessionRoomID,
			SessionID:      room.SessionID,
			RealUserCount:  room.Latest.RealUserCount,
			FakeUserCount:  room.Latest.FakeUserCount,
			TotalUserCount: room.Latest.TotalUserCount,
			UpdateTime:     room.UpdateTime.Add(8 * time.Hour),
		}

		results = append(results, tmp)
	}

	fmt.Println(results)

}

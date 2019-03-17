package nodes

import "time"

type Node struct {
	NodeID      string   `json:"nodeid" form:"nodeid"`
	Slug        string   `json:"slug"`
	Coin        string   `json:"coin" form:"coin"`
	RPCUser     string   `json:"rpcuser" form:"rpcuser"`
	RPCPassword string   `json:"rpcpassword" form:"rpcpassword"`
	Address     string   `json:"address" form:"address"`
	IP          string   `json:"ip" form:"ip"`
	Port        int64    `json:"port" form:"port"`
	Published   bool     `json:"published" form:"published"`
	BitNode     bool     `json:"bitnode" form:"bitnode"`
	Status      NodeStat `json:"stat"`
}

type NodeMap struct {
	Coin          string  `json:"coin" form:"coin"`
	IP            string  `json:"ip" form:"ip"`
	Country_short string  `"country_short" form:"country_short"`
	Country_long  string  `json:"country_long" form:"country_long"`
	Region        string  `json:"region" form:"region"`
	City          string  `json:"city" form:"city"`
	Latitude      float32 `json:"latitude" form:"latitude"`
	Longitude     float32 `json:"longitude" form:"longitude"`
	Zipcode       string  `json:"zipcode" form:"zipcode"`
	Timezone      string  `json:"timezone" form:"timezone"`
}

type NodeStat struct {
	Time time.Time   `json:"time"`
	Data interface{} `json:"data"`
}

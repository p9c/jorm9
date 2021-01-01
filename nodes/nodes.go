package nodes

import (
	"encoding/json"
	"fmt"
	"github.com/p9c/jorm9/jdb"
	"github.com/p9c/jorm9/utils"
	"strconv"
	"strings"
)

func GetNodeStatus() {
	// nd := cs.MainJDB.MJDBGetAll("nodes")
	// for nd.Next() {
	// 	var node Node
	// 	err := nd.Decode(&node)
	// 	if err != nil {
	// 	}
	// 	if node.BitNode {
	// 		c := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	// 		//c := tools.NewClient("duo", "pass", "127.0.0.1", 11066)
	// 		if c == nil {
	// 			fmt.Println("Error node status write")
	// 		}
	// 		params := []string{}
	// 		getinfo, err := c.MakeRequest("getinfo", params)
	// 		if err != nil {
	// 			fmt.Println("Error node status getinfo", err)
	// 		}
	// 		getpeerinfo, err := c.MakeRequest("getpeerinfo", params)
	// 		if err != nil {
	// 			fmt.Println("Error node status getpeerinfo", err)
	// 		}
	// 		// fmt.Println("getinfo", getinfo)
	// 		var nodec Node
	// 		nodec = node
	// 		nodec.GetInfo = getinfo
	// 		nodec.GetPeerInfo = getpeerinfo
	// 		cs.MainJDB.MJDBSet("nodes", node.GPSID, nodec)
	// 	}
	// }
}

func GetNodes() {
	cnodes, err := jdb.DB.ReadAll("nodes")
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, nd := range cnodes {
		var node Node
		if err := json.Unmarshal([]byte(nd), &node); err != nil {
			fmt.Println("Error", err)
		}
		if node.BitNode {
			var peers []interface{}
			fmt.Println("Load Nodes", node.NodeID)

			gpeers := node.JNGetPeerInfo()
			// fmt.Println("Load Nodddddddddddddddddddddddddddddddddddddddes", gpeers)
			if gpeers != nil {
				switch gpeers.(type) {
				case []interface{}:
					peers = gpeers.([]interface{})
					var peer map[string]interface{}
					for _, gpeer := range peers {
						peer = gpeer.(map[string]interface{})
						peerAddress := peer["addr"].(string)

						if node.Address != peerAddress {
							fmt.Println("Load Nodddddddddddddddddddddddddddddddddddddddes", peerAddress)

							splitted := strings.Split(peerAddress, ":")
							peerPort := splitted[len(splitted)-1]
							var peerIP string
							for i := range splitted {
								if i == len(splitted)-1 {
									break
								}
								peerIP += splitted[i]
								if i < len(splitted)-2 {
									peerIP += ":"
								}
								pP, err := strconv.ParseInt(peerPort, 10, 64)
								if err != nil {
								}
								var nodec Node

								nodec.NodeID = peerIP + node.Coin
								nodec.Slug = utils.MakeSlug(peerIP + node.Coin)
								nodec.Coin = node.Coin
								nodec.Address = peerAddress
								nodec.IP = peerIP
								nodec.Port = pP

								nodec.BitNode = false
								// cs.MainJDB.MJDBSet("nodes", GPSID, nodec)
								//jdb.JDB.Write("nodes", nodec.Slug, nodec)

								fmt.Println("Load Nodes")
							}
						}
					}
				}
			}
		}
	}
}

// func GetNodesByCoin(coin string) (nodes []Node) {
// 	cnodes := jdb.GetAllJDB("nodes", Node{})
// 	for _, nd := range cnodes {
// 		node := nd.(Node)
// 		if node.Coin == coin {
// 			nodes = append(nodes, node)
// 		}
// 	}
// 	return
// }

func GetNodesByCoin(coin string) (nodes []Node) {
	cnodes, err := jdb.DB.ReadAll("nodes")
	if err != nil {
		fmt.Println("Error", err)
	}
	for _, nd := range cnodes {
		var node Node
		if err := json.Unmarshal([]byte(nd), &node); err != nil {
			fmt.Println("Error", err)
		}
		if node.Coin == coin {
			nodes = append(nodes, node)
		}
	}
	return
}

func GetBitNodes(coin string) (node Node) {
	cnodes := GetNodesByCoin(coin)
	for _, nd := range cnodes {
		if nd.BitNode {
			node = nd
		}
	}
	return
}

//func GetBitNodeStatus() {
//	//cnodes, err := jdb.JDB.ReadAll("nodes")
//	//if err != nil {
//	//	fmt.Println("Error", err)
//	//}
//	for _, nd := range cnodes {
//		var node Node
//		if err := json.Unmarshal([]byte(nd), &node); err != nil {
//			fmt.Println("Error", err)
//		}
//		if node.BitNode {
//			fmt.Println("Load Node", node.NodeID)
//
//			ns := node.JNGetInfo()
//
//			stat := NodeStat{
//				Time: time.Now(),
//				Data: ns,
//			}
//			node.Status = stat
//
//			jdb.JDB.Write("nodes", node.Slug, node)
//
//		}
//	}
//}

//
//func GetLastBlocks() {
//	var lastblocks []interface{}
//	//ncoins, err := jdb.JDB.ReadAll("coins")
//	//if err != nil {
//	//	fmt.Println("Error", err)
//	//}
//	//for _, nc := range ncoins {
//	//	var coin coins.Coin
//	//	if err := json.Unmarshal([]byte(nc), &coin); err != nil {
//	//		fmt.Println("Error", err)
//	//	}
//
//
//		//if coin.BitNode {
//		//	fmt.Println("Load cossssin", coin.Name)
//
//			node := GetBitNodes(coin.Slug)
//
//			blockcount := node.JNGetBlockCount()
//			minusblockcount := int(blockcount - 10)
//			for ibh := minusblockcount; ibh <= blockcount; {
//				//ib := strconv.Itoa(ibh)
//				blk := node.JNGetBlockByHeight(ibh)
//				lastblocks = append(lastblocks, blk)
//				ibh++
//				fmt.Println("Load lopopopo", ibh)
//
//			}
//			fmt.Println("Load lastblocks", lastblocks)
//
//			//jdb.JDB.Write("chains", coin.Slug, lastblocks)
//
//		//}
//	//}
//}
//
//

package nodes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	ip2location "github.com/ip2location/ip2location-go"
)

//func ABitNodes(w http.ResponseWriter, r *http.Request) {
//	var bns []Node
//	cnodes, err := jdb.JDB.ReadAll("nodes")
//	if err != nil {
//		fmt.Println("Error", err)
//	}
//	for _, nd := range cnodes {
//		var node Node
//		if err := json.Unmarshal([]byte(nd), &node); err != nil {
//			fmt.Println("Error", err)
//		}
//		if node.BitNode {
//			fmt.Println("Load Node", node.NodeID)
//
//			bns = append(bns, node)
//		}
//	}
//	bitNodes := map[string]interface{}{
//		"d": bns,
//		// "coin": gCoin,
//	}
//	bn, err := json.Marshal(bitNodes)
//	if err != nil {
//		fmt.Println("Error encoding JSON")
//		return
//	}
//	w.Write([]byte(bn))
//}

func ANodesMap(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	var nodesIPs []NodeMap
	nds := GetNodesByCoin(coin)

	// var sj []interface{}
	for _, node := range nds {

		var NodeMap NodeMap

		ip2location.Open("./tools/IP2LOCATION-LITE-DB11.BIN")
		results := ip2location.Get_all(node.IP)
		NodeMap.Coin = coin
		NodeMap.IP = node.IP

		NodeMap.Country_short = results.Country_short
		NodeMap.Country_long = results.Country_long
		NodeMap.Region = results.Region
		NodeMap.City = results.City
		NodeMap.Latitude = results.Latitude
		NodeMap.Longitude = results.Longitude
		NodeMap.Zipcode = results.Zipcode
		NodeMap.Timezone = results.Timezone
		ip2location.Close()
		nodesIPs = append(nodesIPs, NodeMap)
	}

	mapNodes := map[string]interface{}{
		"nodes": nodesIPs,
	}
	nodes, err := json.Marshal(mapNodes)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(nodes))
}

//func ABitNode(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	// coin := vars["coin"]
//	nodeid := vars["nodeid"]
//	var node Node
//	if err := jdb.JDB.Read("nodes", nodeid, &node); err != nil {
//		fmt.Println("Error", err)
//	}
//	// gCoin := jdb.GetOneJDB("coins", coin, coins.Coin{})
//	info := node.JNGetInfo()
//
//	bitNode := map[string]interface{}{
//		"d": info,
//		// "coin": gCoin,
//	}
//	bn, err := json.Marshal(bitNode)
//	if err != nil {
//		fmt.Println("Error encoding JSON")
//		return
//	}
//	w.Write([]byte(bn))
//}
func ALastBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	var lastblocks []interface{}
	node := GetBitNodes(coin)
	fmt.Println("coincoincoincoin", coin)
	fmt.Println("nodenodenodenode", node)

	blockcount := node.JNGetBlockCount()
	minusblockcount := int(blockcount - 20)
	for ibh := minusblockcount; ibh <= blockcount; {
		//ib := strconv.Itoa(ibh)
		blk := node.JNGetBlockByHeight(ibh)
		lastblocks = append(lastblocks, blk)
		ibh++
	}
	lb := map[string]interface{}{
		"d": lastblocks,
	}
	out, err := json.Marshal(lb)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ABlock(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//coin := vars["coin"]

	//node := GetBitNodes(coin)
	//lastblock := node.JNGetBlockCount()

	bl := map[string]interface{}{
		//"d": lastblock,
	}
	out, err := json.Marshal(bl)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ABlockHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	bh := vars["blockheight"]
	node := GetBitNodes(coin)
	bhi, _ := strconv.Atoi(bh)
	block := node.JNGetBlockByHeight(bhi)
	bl := map[string]interface{}{
		"d": block,
	}
	out, err := json.Marshal(bl)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ABHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	bh := vars["blockheight"]
	node := GetBitNodes(coin)
	bhi, _ := strconv.Atoi(bh)
	block := node.JNGetBlockTxAddr(bhi)
	bl := map[string]interface{}{
		"d": block,
	}
	out, err := json.Marshal(bl)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func AHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	bh := vars["blockhash"]
	node := GetBitNodes(coin)
	block := node.JNGetBlock(bh)
	bl := map[string]interface{}{
		"d": block,
	}
	out, err := json.Marshal(bl)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ATx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	txid := vars["txid"]
	node := GetBitNodes(coin)
	tX := node.JNGetTx(txid)

	tx := map[string]interface{}{
		"d": tX,
	}
	out, err := json.Marshal(tx)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ARawMemPool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	node := GetBitNodes(coin)
	rawMemPool := node.JNGetRawMemPool()
	rmp := map[string]interface{}{
		"d": rawMemPool,
	}
	out, err := json.Marshal(rmp)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func AMiningInfo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	coin := vars["coin"]
	node := GetBitNodes(coin)

	miningInfo := node.JNGetMiningInfo()

	mi := map[string]interface{}{
		"d": miningInfo,
	}
	out, err := json.Marshal(mi)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func AInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	node := GetBitNodes(coin)

	info := node.JNGetInfo()

	in := map[string]interface{}{
		"d": info,
	}
	out, err := json.Marshal(in)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func APeers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	node := GetBitNodes(coin)
	info := node.JNGetPeerInfo()
	pi := map[string]interface{}{
		"d": info,
	}
	out, err := json.Marshal(pi)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func AAddress(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// coin := vars["coin"]
	// addr := vars["addr"]
	// node := GetBitNodes(coin)

	// aD := node.JNGetAddr(addr)

	// fmt.Println("343434343444", aD)
	// sty, err := json.Marshal(aD)
	// if err != nil {
	// 	fmt.Println(string(sty), err.Error())
	// }

	address := map[string]interface{}{
		// "d": aD,
	}
	out, err := json.Marshal(address)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

package nodes

import (
	"fmt"

	"github.com/parallelcointeam/jorm/tools"
)

// var _ JormNode = &Node{}

// type JormNode interface {
// 	JNGetBlockCount() int
// 	// JNGetBlockByHash(blockhash string) interface{}
// 	JNGetBlock(blockhash string) interface{}
// 	JNGetBlockTxAddr(blockheight int) interface{}
// 	JNGetBlockByHeight(blockheight int) interface{}
// 	JNGetTx(txid string) interface{}
// 	JNGetAddr(cs *CSystem, addr string) interface{}
// 	JNGetRawMemPool() interface{}
// 	JNGetMiningInfo() interface{}
// 	JNGetInfo() interface{}
// }

func (node *Node) JNGetBlockCount() int {
	jrc := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	gbc, err := jrc.MakeRequest("getblockcount", bparams)
	if err != nil {
		fmt.Println("Error node call: ", err)

	}
	getFBlC := gbc.(float64)
	var getBlCount int = int(getFBlC)
	return getBlCount
}

func (node *Node) JNGetBlock(blockhash string) interface{} {
	jrc := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []string{blockhash}
	block, err := jrc.MakeRequest("getblock", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block Error", err)
	}
	return block
}

func (node *Node) JNGetBlockTxAddr(blockheight int) interface{} {
	jrc := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	//bh, err := strconv.Atoi(blockheight)
	//bparams := []int{bh}
	bparams := []int{blockheight}
	blockHash, err := jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block Tx Addr Error", err)
	}
	var block interface{}
	var txs []interface{}
	if blockHash != nil {
		block = node.JNGetBlock((blockHash).(string))
	}
	iblock := make(map[string]interface{})
	iblock = block.(map[string]interface{})

	itxs := iblock["tx"].([]interface{})
	//txs := itxs.([]string)

	for _, itx := range itxs {
		var txid string
		txid = itx.(string)

		verbose := int(1)
		var grtx []interface{}
		grtx = append(grtx, txid)
		grtx = append(grtx, verbose)
		rtx, err := jrc.MakeRequest("getrawtransaction", grtx)
		if err != nil {
			fmt.Println("Jorm Node Get Block Tx Addr Tx Error", err)
		}
		txs = append(txs, rtx)

	}

	//fmt.Println("blockblockblockblockblock", txs)

	blocktxaddr := map[string]interface{}{
		"b": block,
		"t": txs,
	}
	return blocktxaddr
}
func (node *Node) JNGetBlockByHeight(blockheight int) interface{} {
	jrc := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	//bh, err := strconv.Atoi(blockheight)
	//bparams := []int{bh}
	bparams := []int{blockheight}
	blockHash, err := jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block By Height Error", err)
	}
	var block interface{}
	if blockHash != nil {
		block = node.JNGetBlock((blockHash).(string))
	}
	return block
}
func (node *Node) JNGetTx(txid string) interface{} {
	jrc := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	verbose := int(1)
	var grtx []interface{}
	grtx = append(grtx, txid)
	grtx = append(grtx, verbose)
	rtx, err := jrc.MakeRequest("getrawtransaction", grtx)
	if err != nil {
		fmt.Println("Jorm Node Get Tx Error", err)
	}
	// if rtx != nil {
	// 	rawtx = rtx.(Tx)
	// }
	return rtx
}

// func (node *Node) JNGetAddr(addr string) interface{} {

// aD := exJDB.EJDBGetAddr(addr)
// if aD.Addr == "" {
// 	return nil
// }
// 	 return aD
// }
func (node *Node) JNGetRawMemPool() interface{} {
	jrc := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getrawmempool", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Raw Mem Pool Error", err)
	}
	return get
}

func (node *Node) JNGetMiningInfo() interface{} {
	jrc := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getmininginfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Mining Info Error", err)
	}
	return get
}

func (node *Node) JNGetInfo() interface{} {
	jrc := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Info Error", err)
	}
	return get
}
func (node *Node) JNGetPeerInfo() interface{} {
	jrc := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getpeerinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}

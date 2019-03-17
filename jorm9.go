package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/jorm9/coins"
	"github.com/parallelcointeam/jorm9/nodes"
	"github.com/parallelcointeam/jorm9/rts"
	"github.com/robfig/cron"
)

var r = mux.NewRouter()

func main() {
	port := "8999"

	cr := cron.New()
	cr.AddFunc("@every 22s", func() {
		fmt.Println("Radi kron")
		 nodes.GetBitNodeStatus()
		//go nodes.GetLastBlocks()
	})
	cr.Start()

	r.PathPrefix("/s/").Handler(http.StripPrefix("/s/", http.FileServer(http.Dir("./tpl/static/"))))
	// r.PathPrefix("/amp/").Handler(http.StripPrefix("/amp/", http.FileServer(http.Dir("./tpl/amp/"))))

	r.HandleFunc("/", rts.LoginPage) // GET
	r.HandleFunc("/admin/login", rts.Login).Methods("POST")
	r.HandleFunc("/admin/", rts.Home).Methods("GET")
	r.HandleFunc("/admin/", rts.Admin).Methods("POST")

	r.HandleFunc("/logout", rts.Logout).Methods("POST")

	r.HandleFunc("/coins/", coins.Coins).Methods("GET")
	r.HandleFunc("/coins/{slug}", coins.CoinPage).Methods("GET")
	// r.HandleFunc("/coinadd", coins.CoinAdd).Methods("GET")
	r.HandleFunc("/coin", coins.CoinPost).Methods("POST")

	r.HandleFunc("/nodes/", nodes.Nodes).Methods("GET")
	r.HandleFunc("/nodes/{slug}", nodes.NodePage).Methods("GET")
	r.HandleFunc("/nodeadd", nodes.NodeAdd).Methods("GET")
	r.HandleFunc("/node", nodes.NodePost).Methods("POST")

	r.HandleFunc("/api/ini", coins.IniCoins).Methods("GET")

	r.HandleFunc("/api/coins", coins.APICoins).Methods("GET")
	r.PathPrefix("/api/coins/").Handler(http.StripPrefix("/api/coins/", http.FileServer(http.Dir("/web/comhttp/jdb/coins"))))
	// r.PathPrefix("/i/").Handler(http.StripPrefix("/i/", http.FileServer(http.Dir("/web/comhttp/imgs"))))

	// r.HandleFunc("/a/c/a", rts.ACoins).Methods("GET")
	// r.HandleFunc("/", rts.AmpFrontHandler) // GET

	r.HandleFunc("/n", nodes.ABitNodes).Methods("GET")
	r.HandleFunc("/n/{coin}", nodes.ANodesMap).Methods("GET")
	r.HandleFunc("/n/{coin}/{nodeid}", nodes.ABitNode).Methods("GET")

	r.HandleFunc("/e/{coin}/last", nodes.ALastBlock).Methods("GET")
	r.HandleFunc("/e/{coin}/b", nodes.ABlock).Methods("GET")
	r.HandleFunc("/e/{coin}/block/{blockheight}", nodes.ABlockHeight).Methods("GET")
	r.HandleFunc("/e/{coin}/b/{blockheight}", nodes.ABHeight).Methods("GET")
	r.HandleFunc("/e/{coin}/hash/{blockhash}", nodes.AHash).Methods("GET")

	r.HandleFunc("/e/{coin}/tx/{txid}", nodes.ATx).Methods("GET")
	r.HandleFunc("/e/{coin}/rmp", nodes.ARawMemPool).Methods("GET")
	r.HandleFunc("/e/{coin}/gmi", nodes.AMiningInfo).Methods("GET")
	r.HandleFunc("/e/{coin}/info", nodes.AInfo).Methods("GET")
	r.HandleFunc("/e/{coin}/peer", nodes.APeers).Methods("GET")
	r.HandleFunc("/e/{coin}/addr/{addr}", nodes.AAddress).Methods("GET")

	r.HandleFunc("/ttt", nodes.TTT).Methods("GET")

	fmt.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))

}

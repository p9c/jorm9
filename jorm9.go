package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/p9c/jorm9/nodes"
	"github.com/robfig/cron"
)

var r = mux.NewRouter()

func main() {
	port := "8999"

	cr := cron.New()
	cr.AddFunc("@every 22s", func() {
		fmt.Println("Radi kron")
		//nodes.GetBitNodeStatus()
		//go nodes.GetLastBlocks()
	})
	cr.Start()

	r.PathPrefix("/s/").Handler(http.StripPrefix("/s/", http.FileServer(http.Dir("./tpl/static/"))))
	// r.PathPrefix("/amp/").Handler(http.StripPrefix("/amp/", http.FileServer(http.Dir("./tpl/amp/"))))

	//r.HandleFunc("/n", nodes.ABitNodes).Methods("GET")
	//r.HandleFunc("/n/{coin}", nodes.ANodesMap).Methods("GET")
	//r.HandleFunc("/n/{coin}/{nodeid}", nodes.ABitNode).Methods("GET")

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

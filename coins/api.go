package coins

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parallelcointeam/jorm9/jdb"
)

func APICoins(w http.ResponseWriter, r *http.Request) {
	jDB, err := jdb.JDB.ReadAll("coins")
	if err != nil {
		fmt.Println("Error", err)
	}
	var coinsA []APICoin
	for _, f := range jDB {
		var jdbFound Coin
		var jdbAPI APICoin
		if err := json.Unmarshal([]byte(f), &jdbFound); err != nil {
			fmt.Println("Error", err)
		}
		jdbAPI.Name = jdbFound.Name
		jdbAPI.Algo = jdbFound.Algo
		jdbAPI.Symbol = jdbFound.Symbol
		jdbAPI.Slug = jdbFound.Slug
		jdbAPI.BitNode = jdbFound.BitNode

		jdbAPI.Img = jdbFound.Imgs.Img16

		coinsA = append(coinsA, jdbAPI)
	}
	// coinsAa, err := json.Marshal(coinsA)
	// if err != nil {
	// 	fmt.Println("Error encoding JSON")
	// 	return
	// }

	coinsB := API{
		Data: coinsA,
	}

	coins, err := json.Marshal(coinsB)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}

	w.Write([]byte(coins))
}

// func APICoinCall(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	slug := vars["slug"]
// 	coinA := jdb.GetOneJDB("coins", slug, Coin{})

// 	coinB := API{
// 		Data: coinA,
// 	}
// 	coin, err := json.Marshal(coinB)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON")
// 		return
// 	}
// 	w.Write((coin))
// }

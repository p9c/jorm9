package coins

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/p9c/jorm9/utils"
)

// type NodeServices interface {
// 	GetNodeStatus()
// }

func GetCoins() {
	coinresultMapA, err := getCCD("https://min-api.cryptocompare.com/data/all/coinlist")
	if err != nil {
		fmt.Println("Esssrror", err)
		return
	}
	coinresultMap := make(map[string]interface{})
	coinresultMap = coinresultMapA.(map[string]interface{})

	data := coinresultMap["Data"]
	if err != nil {
		fmt.Println("Error", err)
	}
	switch mapCoins := data.(type) {
	case map[string]interface{}:
		for d := range mapCoins {
			var coin Coin
			var cmap = mapCoins[d].(map[string]interface{})
			slug := utils.MakeSlug(cmap["CoinName"].(string))

			//var exco models.Coin

			//			eXCo := jdb.GetAllJDB("coins", Coin{})
			//			exco := eXCo.(Coin)
			//			if slug != exco.Slug {
			imgurl := fmt.Sprint(cmap["ImageUrl"])
			coin.CCID = cmap["Id"].(string)
			coin.Name = cmap["CoinName"].(string)
			coin.Symbol = cmap["Symbol"].(string)
			coin.Algo = cmap["Algorithm"].(string)
			coin.Slug = slug

			compareDataMapA, err := getCCD("https://www.cryptocompare.com/api/data/coinsnapshotfullbyid/?id=" + coin.CCID)
			if err != nil {
				fmt.Println("Error", err)
			}
			compareDataMap := make(map[string]interface{})
			compareDataMap = compareDataMapA.(map[string]interface{})

			data := compareDataMap["Data"]
			if err != nil {
				fmt.Println("Error", err)
			}

			var dcmap = data.(map[string]interface{})
			//cmap := make(map[string]interface{})
			cdcmap := dcmap["General"].(map[string]interface{})

			if cdcmap["Description"] != nil {
				coin.CData.Description = cdcmap["Description"].(string)
			}
			if cdcmap["WebsiteUrl"] != nil {
				coin.CData.WebSite = cdcmap["WebsiteUrl"].(string)
			}
			if cdcmap["TotalCoinSupply"] != nil {
				coin.CData.TotalCoinSupply = cdcmap["TotalCoinSupply"].(string)
			}
			if cdcmap["DifficultyAdjustment"] != nil {
				coin.CData.DifficultyAdjustment = cdcmap["DifficultyAdjustment"].(string)
			}
			if cdcmap["DifficultyAdjustment"] != nil {
				coin.CData.DifficultyAdjustment = cdcmap["DifficultyAdjustment"].(string)
			}
			if cdcmap["ProofType"] != nil {
				coin.CData.ProofType = cdcmap["ProofType"].(string)
			}
			if cdcmap["StartDate"] != nil {
				coin.CData.StartDate = cdcmap["StartDate"].(string)
			}
			if cdcmap["Twitter"] != nil {
				coin.CData.Twitter = cdcmap["Twitter"].(string)
			}

			codexDataMapA, err := getCCD("https://coincodex.com/api/coincodex/get_coin/" + coin.Symbol)
			if err != nil {
				fmt.Println("Error", err)
			}
			if codexDataMapA != nil {

				codexDataMap := codexDataMapA.(map[string]interface{})

				if codexDataMap != nil {

					if codexDataMap["release_date"] != nil {
						coin.CData.StartDate = fmt.Sprint(cdcmap["release_date"])
					}

					// if codexDataMap["whitepaper"] != nil {
					// 	coin.CData.WhitePaper = cdcmap["whitepaper"].(string)
					// }

					// if codexDataMap["social"] != nil {
					// 	social := codexDataMap["social"].(map[string]interface{})
					// 	if err != nil {
					// 		fmt.Println("Error", err)
					// 	}
					// 	fmt.Println("socialsocialsocialsocial", codexDataMap)

					// 	if social["github"] != nil {
					// 		coin.CData.Github = social["github"].(string)
					// 	}
					// 	if social["reddit"] != nil {
					// 		coin.CData.Reddit = social["reddit"].(string)
					// 	}
					// 	if social["message board"] != nil {
					// 		coin.CData.Board = social["message board"].(string)
					// 	}
					// }
				}

			}
			fmt.Println("Coin >>>", coin.Name)
			fmt.Println("ImgUrl >>>", imgurl)
			fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
			coin.Imgs, _ = getIMG("https://cryptocompare.com" + imgurl)
			if err != nil {
				fmt.Println("Problem Insert", err)

			}

			//jdb.JDB.Write("coins", slug, coin)

			// fmt.Println("Insertededed", slug)
			// fmt.Println("Insertedeimgurlimgurlded", imgurl)
			// fmt.Println("Insertedededimgurimgurimgurl", img)
		}
	}
	// }

}

func getCCD(completeURL string) (interface{}, error) {
	resp, err := http.Get(completeURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	mapBody, err := ioutil.ReadAll(resp.Body)

	var mapBodyInterface interface{}
	json.Unmarshal(mapBody, &mapBodyInterface)
	return mapBodyInterface, nil
}

func getIMG(url string) (Imgs, error) {
	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	//fmt.Println("encodedencodedencodedencodedencoded", encoded)
	// encoded := base64.StdEncoding.EncodeToString(content)

	img16, _ := utils.IMGResize(content, utils.Options{Width: 16, Height: 16})
	img32, _ := utils.IMGResize(content, utils.Options{Width: 32, Height: 32})
	img64, _ := utils.IMGResize(content, utils.Options{Width: 64, Height: 64})
	img128, _ := utils.IMGResize(content, utils.Options{Width: 128, Height: 128})
	img256, _ := utils.IMGResize(content, utils.Options{Width: 256, Height: 256})
	imgs := Imgs{
		Img16:  base64.StdEncoding.EncodeToString(img16),
		Img32:  base64.StdEncoding.EncodeToString(img32),
		Img64:  base64.StdEncoding.EncodeToString(img64),
		Img128: base64.StdEncoding.EncodeToString(img128),
		Img256: base64.StdEncoding.EncodeToString(img256),
	}

	return imgs, nil
}

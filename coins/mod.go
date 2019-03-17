package coins

type Coin struct {
	GPSID  string   `json:"id"`
	CCID   string   `json:"ccid" form:"ccid"`
	Name   string   `json:"name" form:"name"`
	Symbol string   `json:"symbol" form:"symbol"`
	Slug   string   `json:"slug" form:"slug"`
	Algo   string   `json:"algo" form:"algo"`
	CData  CoinData `json:"cdata"`
	Imgs   Imgs     `json:"imgs" `
	//Nodes  []explorer.Node `json:"node"`
	//Explorer  Explorer
	BitNode   bool `json:"bitnode" form:"bitnode"`
	Published bool `json:"published" form:"published"`
}

type CoinData struct {
	Name                 string `json:"name"`
	Description          string `json:"desc"`
	WebSite              string `json:"web"`
	TotalCoinSupply      string `json:"total"`
	DifficultyAdjustment string `json:"diff"`
	BlockRewardReduction string `json:"rew"`
	ProofType            string `json:"proof"`
	StartDate            string `json:"start"`
	Twitter              string `json:"tw"`
	Explorer             string `json:"explorer"`
	Board                string `json:"board"`
	Reddit               string `json:"reddit"`
	Github               string `json:"github"`
	WhitePaper           string `json:"wp"`
}
type Imgs struct {
	Img16  string `json:"img16"`
	Img32  string `json:"img32"`
	Img64  string `json:"img64"`
	Img128 string `json:"img128"`
	Img256 string `json:"img256"`
}
type CoinAMP struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Slug    string `json:"slug"`
	Algo    string `json:"algo"`
	BitNode bool   `json:"bitnode"`
}
type API struct {
	Data interface{} `json:"data"`
}
type APICoin struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Slug    string `json:"slug"`
	Algo    string `json:"algo"`
	Img     string `json:"img"`
	BitNode bool   `json:"bitnode"`
}

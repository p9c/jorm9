package coins

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"

	"github.com/parallelcointeam/jorm9/hlp"
	"github.com/parallelcointeam/jorm9/jdb"
	"github.com/parallelcointeam/jorm9/rts"
	"github.com/parallelcointeam/jorm9/utils"
)

func Coins(w http.ResponseWriter, r *http.Request) {
	userName := rts.GetUserName(r)
	if !hlp.IsEmpty(userName) {
		coins := jdb.GetAllJDB("coins", Coin{})
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/coins.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/main.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", coins)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func CoinPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	userName := rts.GetUserName(r)
	if !hlp.IsEmpty(userName) {
		coin := jdb.GetOneJDB("coins", slug, Coin{})
		// fmt.Println("asasasasasasas", coin)

		tmpl, _ := template.New("").ParseFiles("./tpl/admin/coin.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/main.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", coin)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func CoinPost(w http.ResponseWriter, r *http.Request) {
	userName := rts.GetUserName(r)
	if !hlp.IsEmpty(userName) {
		r.ParseForm()

		slug := r.FormValue("slug")
		name := r.FormValue("name")
		// excerpt := r.FormValue("excerpt")
		// content := r.FormValue("content")
		// image := r.FormValue("image")

		slug = utils.MakeSlug(name)

		var coin Coin = Coin{
			// ID:    id,
			// Title: title,
			Slug: slug,
			// Excerpt: excerpt,
			// Content: content,
			// Image:   image,
		}
		fmt.Println("asssssss", slug)

		jdb.JDB.Write("coins", slug, coin)
		http.Redirect(w, r, "/admin/coin/"+slug, 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func IniCoins(w http.ResponseWriter, r *http.Request) {
	GetCoins()
}

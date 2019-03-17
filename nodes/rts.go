package nodes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
	"github.com/parallelcointeam/jorm9/hlp"
	"github.com/parallelcointeam/jorm9/jdb"
	"github.com/parallelcointeam/jorm9/rts"
	"github.com/parallelcointeam/jorm9/utils"
)

func TTT(w http.ResponseWriter, r *http.Request) {
	GetNodes()

}
func Nodes(w http.ResponseWriter, r *http.Request) {
	userName := rts.GetUserName(r)
	if !hlp.IsEmpty(userName) {
		nodes := jdb.GetAllJDB("nodes", Node{})
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/nodes.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/main.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", nodes)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func NodePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	userName := rts.GetUserName(r)
	if !hlp.IsEmpty(userName) {
		node := jdb.GetOneJDB("nodes", slug, Node{})
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/node.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/main.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", node)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
func NodeAdd(w http.ResponseWriter, r *http.Request) {
	userName := rts.GetUserName(r)
	if !hlp.IsEmpty(userName) {
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/nodeadd.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/main.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", "")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func NodePost(w http.ResponseWriter, r *http.Request) {
	userName := rts.GetUserName(r)
	if !hlp.IsEmpty(userName) {
		r.ParseForm()

		nodeid := r.FormValue("nodeid")
		coin := r.FormValue("coin")
		rpcuser := r.FormValue("rpcuser")
		rpcpassword := r.FormValue("rpcpassword")
		address := r.FormValue("address")
		ip := r.FormValue("ip")
		port, _ := strconv.ParseInt(r.FormValue("port"), 10, 64)

		slug := utils.MakeSlug(nodeid)

		var node Node = Node{
			NodeID:      nodeid,
			Slug:        slug,
			Coin:        coin,
			RPCUser:     rpcuser,
			RPCPassword: rpcpassword,
			Address:     address,
			IP:          ip,
			Port:        port,
		}
		fmt.Println("asssssss", node)

		jdb.JDB.Write("nodes", slug, node)
		http.Redirect(w, r, "/nodes/"+slug, 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

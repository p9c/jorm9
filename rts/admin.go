package rts

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/parallelcointeam/jorm9/hlp"
	"github.com/parallelcointeam/jorm9/jdb"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		// var site mod.Site
		// if err := JDB.Read("site", "meta", &site); err != nil {
		// 	fmt.Println("Error", err)
		// }
		tmpl, _ := template.New("base").ParseFiles("./tpl/admin/admin.gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/main.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", "admin")
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func AdminPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lang := vars["lang"]
	page := vars["page"]
	userName := GetUserName(r)
	var pageData interface{}
	if !hlp.IsEmpty(userName) {
		// switch page {
		// case "home":
		// 	pageData = mod.Home{}
		// case "about":
		// 	pageData = mod.About{}
		// case "download":
		// 	pageData = mod.Download{}
		// case "roadmap":
		// 	pageData = mod.RoadMap{}
		// case "contact":
		// 	pageData = mod.Contact{}
		// default:
		// 	pageData = mod.Site{}
		// }
		if err := jdb.JDB.Read(lang, page, &pageData); err != nil {
			fmt.Println("Error", err)
		}
		data := struct {
			Lang string      `json:"lang"`
			Page interface{} `json:"page"`
		}{
			lang,
			pageData,
		}
		tmpl, _ := template.New("").ParseFiles("./tpl/admin/"+page+".gohtml", "./tpl/admin/base.gohtml", "./tpl/admin/main.gohtml", "./tpl/admin/nav.gohtml")
		tmpl.ExecuteTemplate(w, "base", data)

	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func Admin(w http.ResponseWriter, r *http.Request) {
	userName := GetUserName(r)
	if !hlp.IsEmpty(userName) {
		// r.ParseForm()
		// lang := r.FormValue("lang")
		// title := r.FormValue("title")
		// subtitle := r.FormValue("subtitle")
		// welcome := r.FormValue("welcome")
		// about := r.FormValue("about")
		// features := r.FormValue("features")
		// feature1 := r.FormValue("feature1")
		// feature1txt := r.FormValue("feature1txt")
		// feature2 := r.FormValue("feature2")
		// feature2txt := r.FormValue("feature2txt")
		// feature3 := r.FormValue("feature3")
		// feature3txt := r.FormValue("feature3txt")
		// feature4 := r.FormValue("feature4")
		// feature4txt := r.FormValue("feature4txt")
		// gallery := r.FormValue("gallery")
		// specs := r.FormValue("specs")
		// spec1 := r.FormValue("spec1")
		// spec2 := r.FormValue("spec2")
		// spec3 := r.FormValue("spec3")
		// spec4 := r.FormValue("spec4")
		// moto1 := r.FormValue("moto1")
		// moto2 := r.FormValue("moto2")

		// var HOME mod.Home = mod.Home{
		// 	Title:       title,
		// 	SubTitle:    subtitle,
		// 	Welcome:     welcome,
		// 	About:       about,
		// 	Features:    features,
		// 	Feature1:    feature1,
		// 	Feature1txt: feature1txt,
		// 	Feature2:    feature2,
		// 	Feature2txt: feature2txt,
		// 	Feature3:    feature3,
		// 	Feature3txt: feature3txt,
		// 	Feature4:    feature4,
		// 	Feature4txt: feature4txt,
		// 	Gallery:     gallery,
		// 	Specs:       specs,
		// 	Spec1:       spec1,
		// 	Spec2:       spec2,
		// 	Spec3:       spec3,
		// 	Spec4:       spec4,
		// 	Moto1:       moto1,
		// 	Moto2:       moto2,
		// }
		// JDB.Write(lang, "home", HOME)
		// http.Redirect(w, r, "/admin/"+lang+"/home", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

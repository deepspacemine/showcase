package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("secret-password"))

	htmlTemplatesFiles = []string{
		"support.html",
		"userpage.html",
		"home.html",
		"userlogin.html",
	}
	//tmpls *template.Template
	tmpls = template.Must(template.ParseFiles(htmlTemplatesFiles...))
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	t := tmpls.Lookup("home.html")
	t.Execute(w, nil)
}
func supportPageHandler(w http.ResponseWriter, r *http.Request) {
	t := tmpls.Lookup("support.html")
	t.Execute(w, nil)
}

func userpagePageHandler(w http.ResponseWriter, r *http.Request) {

	//user form
	valusername := r.FormValue("username")
	valpassword := r.FormValue("password")
	setpage := 0

	fmt.Println("value", valusername, valpassword)

	//welcomemessage := "Welcome " + valusername
	faillogin := "Login failed " + valusername
	//logout"

	//addedturn := string(added)
	listofData := []string{"Test", "Car", "go", "it", "vestige"}

	//cookie
	//session
	//checkpassword
	session, _ := store.Get(r, "session")
	if r.Method == "POST" {

		if r.FormValue("username") == "cull@example.com" && r.FormValue("password") == "makethefuture" {
			setpage = 3
			session.Values["logged_in"] = r.FormValue("username")
			fmt.Println(" password pass")

		} else if r.FormValue("username") == "" && r.FormValue("password") == "" {
			setpage = 0
		} else {
			setpage = 2
		}
	}
	session.Save(r, w)
	//io.WriteString(w,)
	//displaypage

	switch setpage {
	case 3:
		t := tmpls.Lookup("userlogin.html")
		t.ExecuteTemplate(w, "userlogin.html", listofData)
		fmt.Println("I ran 3")

	case 2:
		t := tmpls.Lookup("userpage.html")
		fmt.Println("I fail run")
		t.Execute(w, faillogin)

	case 0:

		fmt.Println("current", session)
		t := tmpls.Lookup("userpage.html")
		t.Execute(w, "enter password")
	}
	if r.Method == "POST" {
		if r.FormValue("logout") == "logout" {
			session, _ := store.Get(r, "session")
			delete(session.Values, "logged_in")
			session.Save(r, w)
			t := tmpls.Lookup("userpage.html")
			t.Execute(w, "enter password")

			setpage = 2
		}
	}

}
func main() {

	http.Handle("/css/",
		http.StripPrefix("/css",
			http.FileServer(http.Dir("./css"))))

	http.Handle("/images/",
		http.StripPrefix("/images",
			http.FileServer(http.Dir("./images"))))
	
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/userpage.html", userpagePageHandler)
	http.HandleFunc("/support.html", supportPageHandler)

	log.Fatalln(http.ListenAndServe(":1234", nil))
}

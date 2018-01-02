package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
)

type Testlist struct {
	Idslice   []int
	Slicename []string
}

var tl1 = Testlist{
	Idslice:   []int{0, 1, 2},
	Slicename: []string{"user0", "user1", "user2", "user3", "user4"},
}

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
	//listofData := []string{"Test", "Car", "go", "it", "vestige"}

	//cookie
	//session
	//checkpassword
	session, _ := store.Get(r, "session")
	if r.Method == "POST" {

		if r.FormValue("username") == "cull@example.com" && r.FormValue("password") == "makethefuture" {
			setpage = 3
			session.Values["logged_in"] = r.FormValue("username")
			fmt.Println(" password pass", session.Values)

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
		//t := tmpls.Lookup("userlogin.html")

		//t.ExecuteTemplate(w, "userlogin.html", tl1)
		fmt.Println("I ran 3")
		http.Redirect(w, r, "/userlogin.html", 301)

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
	http.HandleFunc("/userlogin.html", userloginPageHandler)
	log.Fatalln(http.ListenAndServe(":1234", nil))
}

// remove string
func dSlicer(slice1 *Testlist, userRemove int) Testlist {
	if userRemove < len(slice1.Slicename) {
		if userRemove == 0 {
			slice1.Slicename = slice1.Slicename[1:]

		} else if userRemove+1 == len(slice1.Slicename) {

			slice1.Slicename = slice1.Slicename[:len(slice1.Slicename)-1]

		} else {
			fmt.Println("remove middle", userRemove, len(slice1.Slicename))

			low := slice1.Slicename[:userRemove-1]
			high := slice1.Slicename[userRemove:]
			slice1.Slicename = append(low, high...)
		}
	}

	return *slice1
}

// add user
func slicer(slice1 *Testlist, userAdd string) Testlist {
	if len(slice1.Slicename) < 199 {

		slice1.Slicename = append(slice1.Slicename, userAdd)
	}
	return *slice1
}

func userloginPageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if session.Values["logged_in"] == "cull@example.com" {
		removetest := 200
		t := tmpls.Lookup("userlogin.html")
		if r.Method != "GET" && r.FormValue("pressbut") != "" {
			removetest, _ := strconv.Atoi(r.FormValue("pressbut"))
			tl1 = dSlicer(&tl1, removetest)
			fmt.Println(" post 1", removetest, tl1)
			t.ExecuteTemplate(w, "userlogin.html", tl1)
		} else if r.Method != "GET" && r.FormValue("newuser") != "" {
			adduser := r.FormValue("newuser")
			tl1 = slicer(&tl1, adduser)
			fmt.Println("\n before", tl1, adduser)
			t.ExecuteTemplate(w, "userlogin.html", tl1)
		} else {

			t.ExecuteTemplate(w, "userlogin.html", tl1)
		}
		fmt.Println("\n before", tl1, removetest)
	} else {
		http.Redirect(w, r, "/", 301)

	}
}


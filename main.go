package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"

	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

type customer struct {
	FirstName, LastName string
	Email, Password     string
	Mobile              int
}

func dbConns() (db *sql.DB) {
	er := godotenv.Load(".env")

	if er != nil {
		log.Fatalf("Error loading .env file")
	}
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func home(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/social.html")
	t.Execute(res, nil)
}

func beginauth(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req) // authentication dengan provider
}

func completeauth(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req) //mendapatkan data yang di autentikasi (name, id, profile)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	db := dbConns()

	userid := user.UserID
	email := user.Email
	firstname := user.FirstName
	lastname := user.LastName
	provider := user.Provider
	name := user.Name
	p, err := db.Prepare("INSERT INTO oauths(userid,email,firstname,lastname,provider) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	p.Exec(userid, email, firstname, lastname, provider)
	fmt.Println("customer userId : ", userid)
	fmt.Println("customer email : ", email)
	fmt.Println("customer FistName : ", firstname)
	fmt.Println("customer LastName : ", lastname)
	fmt.Println("customer data provider : ", provider)
	fmt.Println("customer raw data provider :", name)
	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(res, user)

}

func main() {
	key := "Secret-session-key"
	maxAge := 86400 * 30
	isProd := false

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	er := godotenv.Load(".env")

	if er != nil {
		log.Fatalf("Error loading .env file")
	}

	goth.UseProviders(
		google.New(os.Getenv("GLE_KEY"), os.Getenv("GLE_SECRET"), os.Getenv("GLE_CALLBACKURL"), "email", "profile"),
		facebook.New(os.Getenv("FB_KEY"), os.Getenv("FB_SECRET"), os.Getenv("FB_CALLBACKURL"), "email"),
	)

	p := pat.New()

	p.Get("/auth/{provider}/callback", completeauth)
	p.Get("/auth/{provider}", beginauth)
	p.HandleFunc("/", home)
	log.Println("listening on http://localhost:8000")
	p.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8000", p)

}

package main
import (

	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	_ "github.com/lib/pq"
	"usewebapps.com/views"
	"usewebapps.com/module"
)
//var homeTemplate *template.Template
//var contactTemplate *template.Template

var homeView *views.View
var contactView *views.View
var signupView   *views.View

const (
host = "192.168.56.101"
user = "postgres"
password = "postgres"
dbname = "calhounio_demo"

)

func main() {

	var err error
	//homeTemplate, err = template.ParseFiles("views/home.gohtml")
	//contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	//calhounio_demo

	homeView = views.NewView("bootstrap","views/home.gohtml")
	contactView = views.NewView("bootstrap","views/contact.gohtml")
	signupView = views.NewView("bootstrap","views/signup.gohtml")

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",host, user, password, dbname)

	module.SelectDATA(psqlInfo)

	module.CheckIfError(err)




	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3002", r)
}



func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Println(homeView.Layout)
	homeView.Render(w, nil)
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	contactView.Render(w, nil)
}
func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	signupView.Render(w, nil)
}

func calcul(args ... int) int {
	total := 0
	for _, v := range args {

		total += v
	}
return total
}
// Create a webpage that serves at localhost:8080 and will display the name
// in the url when the url is localhost:8080/name - use req.URL.Path to do this

package main
import (
	"io"
	"log"
	"net/http"
)

func handler(res http.ResponseWriter,req *http.Request){
	key := "n"
	val := req.FormValue(key)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<form method="GET">
		 <input type="text" name="n">
		 <input type="submit">
		</form>`+ val)
}

func main(){
	http.HandleFunc("/",handler)
	log.Println("Listening...")
	http.ListenAndServe(":8080",nil)
}
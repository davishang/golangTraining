// Store a uuid in a cookie
// Value: uuid
// Store the uuid
// key:uuid
// value: your name
// Retrieve the uuid from the cookie
// Retrieve the uuid & value from memcache
// Import packages:
// "github.com/nu7hatch/gouuid"
// "google.golang.org/appengine"
// "google.golang.org/appengine/memcache"

package main

import (
	"fmt"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/retrieve", noConfusion)
}

func index(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	//test
	//fmt.Fprint(res, "Test")
	id, _ := uuid.NewV4()

	// set cookie
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id.String(),
		// Secure: true,
		HttpOnly: true,
	}
	http.SetCookie(res, cookie)

	// set memcache
	ctx := appengine.NewContext(req)
	item1 := memcache.Item{
		Key:   id.String(),
		Value: []byte("David"),
	}
	memcache.Set(ctx, &item1)

	fmt.Fprint(res, "EVERYTHING SET ID:"+id.String())
}

	mySetMemC(id, req)

	fmt.Fprint(res, "EVERYTHING SET ID: " + id)

}

func noConfusion(res http.ResponseWriter, req *http.Request) {

	html := `
	<form method="POST">
	    <input type="text" name="koala">
	    <input type="submit" value="submit">
	</form>
	`

	if req.Method == "POST" {
		id := req.FormValue("koala")

		// get cookie value
		cookie, _ := req.Cookie("session-id")
		if cookie != nil {
			html += `
			<br>
			<p>Value from cookie: ` + cookie.Value + `</p>
			`
		}

		// get memcache value
		ctx := appengine.NewContext(req)
		item, _ := memcache.Get(ctx, id)
		if item != nil {
			html += `
			<br>
			<p>
			Value from memcache: ` + string(item.Value) + `
			</p>
		`
		}
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(res, html)
}
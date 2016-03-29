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
}

func index(res http.ResponseWriter, req *http.Request) {
	//test
	//fmt.Fprint(res, "Test")
	id, _ := uuid.NewV4()

	mySetMemC(id, req)

	fmt.Fprint(res, "EVERYTHING SET ID: " + id)

}

func Retrieve

func mySetCookie(res http.ResponseWriter) string {
	cookie := &http.Cookie{
		Name:  "session-id",
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,

		item, _ := memcache.Get(ctx, "foo")
		if item != nil {
		fmt.Fprintln(res, string(item.Value))
		}
	}
}

func mySetMemC (id string, req *http.Request){
	ctx := appengine.NewContext(req)

	item1 := memcache.Item{
		Key:   id,
		Value: []byte("David"),
	}

	memcache.Set(ctx, &item1)
}
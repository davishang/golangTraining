// Create a webpage which writes a cookie to the client's machine.
// This cookie should be designed to create a session and should use a UUID,
// HttpOnly, and Secure (though you'll need to comment secure out).

package main

import  (
	"fmt"
	"log"
	// generate random (version 4) UUIDs with:
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session-id")
	if err != nil{
		// Implementation of Universally Unique Identifier (UUID)
		// The returned UUID type is a 16 byte array...can retrieve
		// the binary value easily...also provides the standard
		// hex string representation via its String() method.
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-id",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
}

func main(){
	http.HandleFunc("/", handler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

// go get uuid
// https://github.com/nu7hatch/gouuid
// NewV4

/****Some other type UUID

func NewV1() UUID
NewV1 returns UUID based on current timestamp and MAC address.

func NewV2
func NewV2(domain byte) UUID
NewV2 returns DCE Security UUID based on POSIX UID/GID.

func NewV3
func NewV3(ns UUID, name string) UUID
NewV3 returns UUID based on MD5 hash of namespace UUID and name.

func NewV4
func NewV4() UUID
NewV4 returns random generated UUID.

func NewV5
func NewV5(ns UUID, name string) UUID
NewV5 returns UUID based on SHA-1 hash of namespace UUID and name.
 */
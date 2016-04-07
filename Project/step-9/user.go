package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func newVisitor(name string) *http.Cookie {
	mm := initialModel()
	id, _ := uuid.NewV4()
	return makeCookie(mm, id.String(), name)
}

func currentVisitor(m model, id string, name string) *http.Cookie {
	mm := marshalModel(m)
	return makeCookie(mm, id, name)
}

func makeCookie(mm []byte, id string, name string) *http.Cookie {
	b64 := base64.URLEncoding.EncodeToString(mm)
	code := getCode(b64)
	cookie := &http.Cookie{
		Name:  name,
		Value: id + "|" + b64 + "|" + code,
		// Secure: true,
		HttpOnly: true,
	}
	return cookie
}

func marshalModel(m model) []byte {
	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return bs
}

func initialModel() []byte {
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"myPic.jpg",
		},
	}
	return marshalModel(m)
}
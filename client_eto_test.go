package finfree_http_client

import (
	"log"
	"testing"
)

func TestEto(t *testing.T) {
	cl := New("https://finfree.app/profile")

	var resp HandleResponse
	_, err := cl.Get("/handle", &resp, NewQuery("username", "ErtugrulAcar"))

	if err != nil {
		log.Println("Error on /handle request. Err ->", err)
		t.FailNow()
	}

	log.Println(resp)

}

type HandleResponse struct {
	Username       string `json:"username"`
	HandleName     string `json:"handle_name"`
	Bio            string `json:"bio"`
	DisplayName    string `json:"display_name"`
	ContentCreator bool   `json:"content_creator"`
	Private        bool   `json:"private"`
}

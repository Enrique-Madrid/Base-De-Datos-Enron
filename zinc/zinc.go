package zinc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Mail struct {
	Name     string `json:"-"`
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Category string `json:"category"`
	Body     string `json:"body"`
}

type Auth struct {
	User     string
	Password string
}

var auth Auth

func AuthValues(username string, pass string) {
	auth = Auth{
		User:     username,
		Password: pass,
	}
}

func CreateJSON(MailString []Mail) {
	for _, mail := range MailString {
		// Create a JSON string for each mail excluding the "Name" field
		jsonMail, err := json.Marshal(struct {
			From     string `json:"from"`
			To       string `json:"to"`
			Subject  string `json:"subject"`
			Category string `json:"category"`
			Body     string `json:"body"`
		}{
			From:     mail.From,
			To:       mail.To,
			Subject:  mail.Subject,
			Category: mail.Category,
			Body:     mail.Body,
		})
		if err != nil {
			fmt.Println("Error marshalling mail:", err)
			continue
		}
		Indexer(jsonMail, mail.Name)
	}
}

func Indexer(json []byte, name string) {
	req, err := http.NewRequest("POST", "http://localhost:4080/api/"+name+"/_doc", strings.NewReader(string(json)))
	if err != nil {
		log.Println(err)
	}
	req.SetBasicAuth(auth.User, auth.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}

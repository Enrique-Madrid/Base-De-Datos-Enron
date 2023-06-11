package zinc

import (
	"encoding/json"
	"io"
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
type MailWithHeader struct {
	Name     string `json:"name"`
	From     string `json:"mail.from"`
	To       string `json:"mail.to"`
	Subject  string `json:"mail.subject"`
	Category string `json:"mail.category"`
	Body     string `json:"mail.body"`
}

func AuthValues(username string, pass string) {
	auth = Auth{
		User:     username,
		Password: pass,
	}
}

var auth Auth
var jsonToIndex []byte = nil
var IndexTemplate []byte = []byte("\n" + `{"index" : { "_index" : "enron_database"}}` + "\n")
var jsonMail []byte

var cant int = 0

func CreateJSON(mail Mail) {
	cant++
	jsonMail, err := json.Marshal(
		MailWithHeader{
			Name:     mail.Name,
			From:     mail.From,
			To:       mail.To,
			Subject:  mail.Subject,
			Category: mail.Category,
			Body:     mail.Body,
		})
	temp := append(IndexTemplate, jsonMail...)
	jsonToIndex = append(jsonToIndex, temp...)

	if err != nil {
		log.Println("\033[31mError Marshaling:\033[0m", err)
	}

	if cant%1000 == 0 {
		log.Println("\033[32mMail #:\033[0m", cant, "\033[32m indexed successfully\033[0m")
		go Indexer(jsonToIndex)
		jsonToIndex = nil
	}

}

func SendJSON() {
	Indexer(jsonToIndex)
}

func Indexer(json []byte) {
	req, err := http.NewRequest("POST", "http://localhost:4080/api/es/_bulk", strings.NewReader(string(json)))
	if err != nil {
		log.Println("\033[31mError:\033[0m", err)
	}
	req.SetBasicAuth(auth.User, auth.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	// Check the credentials
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("\033[31mError:\033[0m", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("\033[31mError indexing mail:\033[0m", err)
	}

	if strings.Contains(string(body), "auth") {
		log.Fatal("\033[31mError indexing mail (Invalid Credentials):\033[0m", string(body))
	}
}

func Searcher(search_term string, from string) []byte {
	query := `{
		"search_type": "matchphrase",
		"query": {
			"term": "` + search_term + `",
			"field": "mail.body"
		},
		"sort_fields": [
			"-@timestamp"
		],
		"from": ` + from + `,
		"max_results": 20,
				
		"_source": []
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_database/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(auth.User, auth.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("\033[31mError:\033[0m", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("\033[31mError finding mail:\033[0m", err)
	} else {
		log.Println("\033[32mMail finded successfully\033[0m")
	}

	// Jsonify the response
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	body, err = json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Println("\033[31mError marshalling mail:\033[0m", err)
	}

	return body
}

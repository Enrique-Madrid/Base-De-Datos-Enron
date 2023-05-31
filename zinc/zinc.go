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
	Name string `json:"name"`
	Mail Mail   `json:"mail"`
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
		// Create a JSON string for each mail
		jsonMail, err := json.Marshal(MailWithHeader{
			Name: mail.Name,
			Mail: Mail{
				From:     mail.From,
				To:       mail.To,
				Subject:  mail.Subject,
				Category: mail.Category,
				Body:     mail.Body,
			},
		})
		if err != nil {
			log.Println("\033[31mError Marshaling:\033[0m", err)
			continue
		}
		Indexer(jsonMail)
	}
}

func Indexer(json []byte) {
	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_database/_doc", strings.NewReader(string(json)))
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
	} else {
		log.Println("\033[32mMail indexed successfully\033[0m")
	}
}

func Searcher(search_term string, index string, from string) []byte {
	query := `{
		"search_type": "matchphrase",
		"query": {
			"term": "` + index + `",
			"field": "name"
		},
		"sort_fields": [
			"-@timestamp"
		],
		"from": ` + from + `,
		"max_results": 20,
				
		"_source": []
    }`
	log.Println(query)
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
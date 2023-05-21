package main

import (
	"basededatos/zinc"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// Grab the first argument as the directory to walk (e.g. "../../enron_mail_20110402/maildir")
	directoryToWalk := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]

	zinc.AuthValues(username, password)

	mailArray := make([]zinc.Mail, 0)

	err := filepath.Walk(directoryToWalk, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error walking directory: ", err)
			return nil
		}
		if !info.IsDir() {
			mail, err := os.Open(path)
			if err != nil {
				log.Println("Error opening file: ", err)
				return nil
			}
			defer mail.Close()

			scanner := bufio.NewScanner(mail)
			var from, to, subject, body string
			isBody := false

			// Read the content of the file line by line
			for scanner.Scan() {
				line := scanner.Text()

				if !isBody {
					if strings.HasPrefix(line, "From: ") {
						from = strings.TrimPrefix(line, "From: ")
					} else if strings.HasPrefix(line, "To: ") {
						to = strings.TrimPrefix(line, "To: ")
					} else if strings.HasPrefix(line, "Subject: ") {
						subject = strings.TrimPrefix(line, "Subject: ")
					} else if line == "" {
						isBody = true
					}
				} else {
					body += line + "\n"
				}
			}

			// Split the path to get the category and the name of the file
			sub_path := strings.Split(path, string(os.PathSeparator))
			category := sub_path[len(sub_path)-2]
			name := sub_path[len(sub_path)-3]

			mailArray = append(mailArray, zinc.Mail{name, from, to, subject, category, body})
		}
		return nil
	})

	if err != nil {
		log.Println("Error walking directory: ", err)
		return
	}

	zinc.CreateJSON(mailArray)
}

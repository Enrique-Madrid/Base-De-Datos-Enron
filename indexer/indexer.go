package main

import (
	"basededatos/zinc"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"runtime/pprof"
)

func main() {

	if len(os.Args) != 2 {
		log.Println("\x1b[31;1mIncorrect number of arguments: indexer <directory>\x1b[0m")
		return
	}

	// Grab the first argument as the directory to walk (e.g. "../../enron_mail_20110402/maildir")
	directoryToWalk := os.Args[1]

	env := os.Environ()

	for _, value := range env {
		if strings.HasPrefix(value, "ZINC_") {
			split := strings.Split(value, "=")
			if len(split) != 2 {
				log.Println("Error parsing environment variable: ", value)
				return
			}
			os.Setenv(split[0], split[1])
		}
	}

	// Start the profiler
	f, err := os.Create("profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)

	// Flamegraph
	// go tool pprof -http=":8080" profile.prof

	// Memory and CPU profiling
	defer pprof.Lookup("heap").WriteTo(os.Stdout, 10)
	defer pprof.Lookup("goroutine").WriteTo(os.Stdout, 10)
	defer pprof.Lookup("threadcreate").WriteTo(os.Stdout, 10)
	defer pprof.Lookup("block").WriteTo(os.Stdout, 10)

	defer pprof.StopCPUProfile()

	username := os.Getenv("ZINC_FIRST_ADMIN_USER")
	password := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")

	zinc.AuthValues(username, password)

	mailArray := make([]zinc.Mail, 0)

	err = filepath.Walk(directoryToWalk, func(path string, info os.FileInfo, err error) error {
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

			mailArray = append(mailArray, zinc.Mail{
				Name:     name,
				From:     from,
				To:       to,
				Subject:  subject,
				Category: category,
				Body:     body,
			})

		}
		return nil
	})

	if err != nil {
		log.Println("Error walking directory: ", err)
		return
	}

	zinc.CreateJSON(mailArray)

	log.Println("\033[32mIndex finished\033[0m")
}

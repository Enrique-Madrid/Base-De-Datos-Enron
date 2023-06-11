## Enron Email Indexer Project

![mamuro-email](https://i.ibb.co/3fnz2Bz/banner.jpg)

This project is a web application created using Vue.js, Vuex, Golang, ZincSearch, and Chi router for Golang as backend, which performs indexing and search of emails from the Enron corpus. The Enron email corpus is a collection of over 500,000 emails that were sent and received by Enron Corporation.

### Dependencies

* Golang
* Vue.js
* Vuex
* ZincSearch (similar to Elasticsearch)
* Chi router for Golang

### Installation

1. Clone this repository to your local machine.
2. Install Golang if not already installed.
3. Install ZincSearch by going to the [ZincSearch repository](https://github.com/zincsearch/zincsearch) and following the installation instructions.
4. Install Vue.js and Vuex if not installed by running `npm install vue` and `npm install vuex` respectively.
5. Install Chi router for Golang by running `go get -u github.com/go-chi/chi/v5`.
6. Run `go mod download` to download Go module dependencies.

### Usage

1. To start the application, go to the run directory of the project using `cd run` in your terminal and then run `./mamuro -p <PORT>` to start the server. ('<PORT>' is the desired port to run the server). You should be able to see the application running by accessing `localhost:<PORT>` in your browser.
2. To index the Enron email corpus, run first `cd indexer` and then `./indexer <PATH>` where `<PATH>` is the path to the directory containing the Enron email corpus.
3. You can now start searching for emails using the search bar at the top of the application. You can search for emails by keyword or phrase, sender or recipient, and date range.
4. The results of your search will be displayed on the page with email metadata, including sender, recipient and subject. You can read the content of each email by clicking on it.

![Enron Email Data](https://i.ibb.co/7t4J7fK/image.png)

### Contributing

Pull requests and contributions are always welcome. If you want to contribute to this project, please follow these steps:

1. Fork this repository.
2. Create your feature branch (`git checkout -b feature/your-feature`).
3. Make necessary changes and improvements.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature/your-feature`).
6. Create a new Pull Request.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
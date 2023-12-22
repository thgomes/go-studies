package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type Addr struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeScreenHandle)
	mux.HandleFunc("/contact", contactScreenHandle)
	mux.HandleFunc("/about", aboutScreenHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", mux)
}

func homeScreenHandle(w http.ResponseWriter, r *http.Request) {
	var (
		addr    *Addr
		err     error
		payload struct {
			Found bool
			Cep   string
			Addr  Addr
		}
	)

	payload.Found = false
	payload.Cep = r.URL.Query().Get("cep")
	if payload.Cep != "" {
		addr, err = findCep(payload.Cep)
		if err != nil {
			fmt.Println(err)
		} else {
			payload.Found = true
			payload.Addr = *addr
		}
	}

	templates := []string{"views/header.html", "views/footer.html", "views/home.html"}
	homeTemplate := template.Must(template.New("home.html").ParseFiles(templates...))
	err = homeTemplate.Execute(w, payload)
	if err != nil {
		fmt.Println(err)
	}
}

func contactScreenHandle(w http.ResponseWriter, r *http.Request) {
	templates := []string{"views/header.html", "views/footer.html", "views/contact.html"}
	contactTemplate := template.Must(template.New("contact.html").ParseFiles(templates...))
	err := contactTemplate.Execute(w, "contact.html")
	if err != nil {
		fmt.Println(err)
	}
}

func aboutScreenHandler(w http.ResponseWriter, r *http.Request) {
	templates := []string{"views/header.html", "views/footer.html", "views/about.html"}
	aboutTemplate := template.Must(template.New("about.html").ParseFiles(templates...))
	err := aboutTemplate.Execute(w, "about.html")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func findCep(cep string) (*Addr, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	var addr Addr
	err = json.Unmarshal(data, &addr)
	if err != nil {
		return nil, err
	}

	return &addr, nil
}

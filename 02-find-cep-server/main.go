package main

import (
	"encoding/json"
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
	http.HandleFunc("/", findCepHandler)
	http.ListenAndServe(":8080", nil)
}

func findCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepQueryParam := r.URL.Query().Get("cep")
	if cepQueryParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addr, err := findCep(cepQueryParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Context-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addr)

}

func findCep(cep string) (*Addr, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var addr Addr
	err = json.Unmarshal(res, &addr)
	if err != nil {
		return nil, err
	}

	return &addr, nil
}

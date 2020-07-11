package main

import "net/http"

type Coaster struct {
	Name         string `json:"name,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	ID           string `json:"id,omitempty"`
	InPark       string `json:"in_park,omitempty"`
	Height       int    `json:"height,omitempty"`

}


type coastersHandlers struct {
	store map[string]Coaster
}

func coastersHandlers(w http.ResponseWriter, r *http.Request){

}

func main() {
	//Default handler will be nil
	http.HandleFunc("/coasters", coastersHandlers)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
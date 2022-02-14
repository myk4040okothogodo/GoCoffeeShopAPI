package handlers

import (
    "log"
    "net/http"
    "github.com/myk4040okothogodo/GoMicroserve/data"
    )

type Products struct {
   l *log.Logger

}

func NewProducts(l *log.Logger) *Products {
    return &Products{l}
}


// ServeHTTP is the main entry point for the handler and satifies the http.Handler interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
    if r.Method == http.MethodGet{
        p.getProducts(rw, r)
        return
    }
    // catch all
    rw.WriteHeader(http.StatusMethodNotAllowed)
}
//getProducts returns the products form the data store

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request){
    
    // fetch the products from the datastore
    lp := data.GetProducts()
    
    //serialize the list to JSON
    err := lp.ToJSON(rw)
    if err != nil{
        http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
    }
}

package main
import (
  "time"
  "context"
  "os"
  "os/signal"
  "net/http"
  "log"
  "github.com/myk4040okothogodo/GoMicroserve/handlers"
)

func main() {

    
    l := log.New(os.Stdout, "product-api", log.LstdFlags)
    
    //create the handlers
    ph := handlers.NewProducts(l)

    // create a new serve mux and register the handlers
    sm := http.NewServeMux()
    sm.Handle("/", ph)
   
    // create a new server 
    s := &http.Server{
        Addr: ":9090",
        Handler: sm,
        ErrorLog: l,
        IdleTimeout:  120 * time.Second,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    
    }
    go func(){
      l.Println("\n Starting server on port 9090: \n")
        
        err := s.ListenAndServe()
        if err != nil {
            l.Printf("Error starting server: %s\n", err)
            os.Exit(1)
        }
    }()
    
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, os.Kill)
    
    sig := <- c
    l.Println("Received signal", sig)
    
    //gracefully shutdown the server, wiating max 30 seconds for current operations to finish
    tc, _  := context.WithTimeout(context.Background(), 30*time.Second)
    s.Shutdown(tc)
    

}

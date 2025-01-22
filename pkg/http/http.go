package http

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Init() {
	// Fancy ASCII banner
	banner := `
  ______ _____  ___  _  _     _____  _____   ____ _______ ____  
 |  ____|  __ \|__ \| || |   |  __ \|  __ \ / __ \__   __/ __ \ 
 | |__  | |__) |  ) | || |_  | |__) | |__) | |  | | | | | |  | |
 |  __| |  _  /  / /|__   _| |  ___/|  _  /| |  | | | | | |  | |
 | |    | | \ \ / /_   | |   | |    | | \ \| |__| | | | | |__| |
 |_|    |_|  \_\____|  |_|   |_|    |_|  \_\\____/  |_|  \____/ 
                                                                
Starting server on :8080...
`
	fmt.Print(banner)

	http.HandleFunc("/", hello)

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server closed")
		} else if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, fmt.Sprintf("Hello, %s!", req.URL.Path[1:]))
	if err != nil {
		log.Fatal(err)
	}
}

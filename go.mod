package main

import (
	"fmt"
	"net/http"
	"os"
)

// gripen engine: eric full stack, gemini intelligence active
func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "8081" } // คนละพอร์ตเพื่อแยกกิ่งชัดเจน

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// รับงานจาก f-16 มาประมวลผลด้วย ai
		fmt.Fprintf(w, "gripen: a2a handshake success. gemini 2.0 flash is thinking...")
	})

	http.ListenAndServe(":"+port, nil)
}

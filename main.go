package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "8081" }

	// 🔑 เสียบกุญแจที่เจ้านายให้มา (AIzaSyD3...)
	apiKey := "AIzaSyD3A1N2ZZAZEmu-LCUUYTwon7Um5r1wjJY" 

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// 🛡️ 1. Security Check (Snake Nudge Protocol)
		auth := r.Header.Get("X-ThitNuea-Auth")
		if auth != "DragonScale2026" {
			http.Error(w, "🚫 Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.Background()
		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		// 🧠 2. เรียกใช้โมเดล Gemini 2.0 Flash (หรือ 1.5 ตามสิทธิ์การใช้งาน)
		model := client.GenerativeModel("gemini-1.5-flash") 
		
		// รับงานจาก F-16 มาให้ Gemini วิเคราะห์
		prompt := "วิเคราะห์สถานการณ์จาก F-16 และสรุปแนวทางป้องกันความปลอดภัยให้ ThitNueaHub"
		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		
		if err != nil {
			http.Error(w, "Gemini Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// 3. ส่งคำตอบกลับไป
		fmt.Printf("🐉 Gripen: Gemini has processed the threat.\n")
		fmt.Fprintf(w, "Gripen Insight: %v", resp.Candidates[0].Content.Parts[0])
	})

	fmt.Println("🚀 Gripen Engine is standby on port :"+port)
	http.ListenAndServe(":"+port, nil)
}

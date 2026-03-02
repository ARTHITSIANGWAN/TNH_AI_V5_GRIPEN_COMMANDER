package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/genai" // 🏹 สมองกริพเพน 2.0 & 3.1
)

/* * 📜 LICENSE: MIT (Copyright 2026 ARTHIT SIANGWAN)
 * 🛡️ Project: F-16 Defender v.2 - Triple Yum Edition
 * 🐍 Agents: Nam-Ing (Voice/Recall) & Phrai Thong (Security)
 */

func main() {
	// ⛽ ยำที่ 1: ระบบตั้งเวลา (Gas Station Scheduler)
	go runDailyScheduler()

	port := os.Getenv("PORT")
	if port == "" { port = "8081" }

	// 🛰️ ยำที่ 2: ระบบ Handshake & Remote API
	http.HandleFunc("/process", handleA2AHandshake)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🛡️ THITNUEA GRIPEN MASTER FUSION\nStatus: Online ✅\nEngine: Gemini 2.0/3.1 (Nano Banana)\nMode: Zero-Garbage 2000%%")
	})

	log.Printf("🚀 ฐานทัพ Gripen บินขึ้นที่พอร์ต %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// --- [🧠 กริพเพนสมอง 2.0: ระบบเรียกใช้ AI ] ---
func askGripen(ctx context.Context, prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGoogleAI,
	})
	if err != nil { return "", err }

	// ปักหมุดใช้ Gemini 2.0 Flash (หรือเปลี่ยนเป็น 3.1-flash-image-preview สำหรับงานรูปภาพ)
	modelID := "gemini-2.0-flash" 
	resp, err := client.Models.GenerateContent(ctx, modelID, genai.Text(prompt), nil)
	if err != nil { return "", err }

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
	}
	return "No response", nil
}

// --- [🐍 น้ำอิงสายดุ: ระบบบันไดงู & Retry ] ---
func runDailyScheduler() {
	targetHours := []int{8, 12, 20} // 08:00, 12:00, 20:00
	loc := time.FixedZone("Asia/Bangkok", 7*60*60)

	for {
		now := time.Now().In(loc)
		nextRun := calculateNext(now, targetHours, loc)
		fmt.Printf("😴 น้ำอิงพักเครื่อง.. เจอกันรอบถัดไป: %s\n", nextRun.Format("15:04:05"))
		time.Sleep(time.Until(nextRun))

		// บันไดงู Retry 3 ครั้ง [Zero-Garbage Policy]
		for i := 0; i < 3; i++ {
			res, err := askGripen(context.Background(), "Task: SME Insight Summary. Style: Aggressive Thai.")
			if err == nil {
				fmt.Printf("✅ Mission Accomplished: %s\n", res)
				break
			}
			fmt.Printf("⚠️ Snake Nudge Recall! (Attempt %d)\n", i+1)
			time.Sleep(10 * time.Second)
		}
	}
}

// --- [🛠️ ระบบคำนวณเวลา & Security ] ---
func calculateNext(now time.Time, hours []int, loc *time.Location) time.Time {
	for _, h := range hours {
		t := time.Date(now.Year(), now.Month(), now.Day(), h, 0, 0, 0, loc)
		if t.After(now) { return t }
	}
	return time.Date(now.Year(), now.Month(), now.Day()+1, hours[0], 0, 0, 0, loc)
}

func handleA2AHandshake(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-ThitNuea-Auth") != os.Getenv("A2A_SECRET_KEY") {
		http.Error(w, "🚫 Unauthorized", http.StatusUnauthorized)
		return
	}
	// ... Logic รับ JSON และเรียก askGripen เหมือนตัวก่อนหน้า
}

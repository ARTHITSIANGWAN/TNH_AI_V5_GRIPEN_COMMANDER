package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"google.golang.org/genai" // SDK ตัวจี๊ด Gemini 2.0
)

// --- [🛡️ IDENTITY & AGENTS] ---
// 1. Gripen Engine (Gemini 2.0 Flash) - Intelligence
// 2. Nam-Ing (Supervisor) - Snake Nudge & Recall
// 3. Dark-Relay Scheduler - Gas Station (08:00, 12:00, 20:00)

func main() {
	// 1. รันระบบตั้งเวลา (Scheduler) แยกเป็น Background Task
	go runGasStationScheduler()

	port := os.Getenv("PORT")
	if port == "" { port = "8081" }

	// 2. Endpoint สำหรับ A2A Handshake (Gripen API)
	http.HandleFunc("/process", handleGripenProcess)
	
	// 3. หน้า Dashboard หลัก
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🛡️ THITNUEA GRIPEN FUSION 2000%% ONLINE\nStatus: Nam-Ing Active 🐍 | Intelligence: Gemini 2.0 Flash 🏹")
	})

	fmt.Printf("🚀 Gripen Engine Active on Port %s | 💰 Mode: Zero-Garbage\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// --- [⏰ ส่วนที่ 1: Gas Station Scheduler (มุดดิน)] ---
func runGasStationScheduler() {
	targetHours := []int{8, 12, 20}
	loc := time.FixedZone("Asia/Bangkok", 7*60*60)

	for {
		now := time.Now().In(loc)
		nextRun := calculateNextRun(now, targetHours, loc)

		fmt.Printf("😴 [Gas Station]: พักเครื่อง.. รอบถัดไปคือ %s\n", nextRun.Format("15:04:05"))
		time.Sleep(time.Until(nextRun))

		fmt.Println("🏁 [Nam-Ing]: ปล่อยตัว Gripen 4x100 Relay Sequence...")
		executeIntelligenceTask()
	}
}

// --- [🧠 ส่วนที่ 2: Intelligence Engine (บนฟ้า - Gripen)] ---
func askGripen(ctx context.Context, prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGoogleAI,
	})
	if err != nil { return "", err }

	// ใช้รุ่น 2.0 Flash ตาม Gripen SDK
	resp, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(prompt), nil)
	if err != nil { return "", err }

	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		// ดึงข้อความออกมาแบบสะอาดๆ
		return fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0]), nil
	}
	return "No response from AI", nil
}

// --- [🐍 ส่วนที่ 3: น้ำอิง Supervisor (การจัดการงาน)] ---
func executeIntelligenceTask() {
	ctx := context.Background()
	// บังคับ Policy: Zero-Garbage และดุดัน
	prompt := "Task: Generate Elite Insight. Policy: Zero-Garbage. Style: Aggressive Thai. Role: ThitNuea Finisher."
	
	// ระบบบันไดงู (Snake Nudge) Retry 3 ครั้ง
	for i := 0; i < 3; i++ {
		result, err := askGripen(ctx, prompt)
		if err != nil {
			log.Printf("⚠️ [Nam-Ing]: Recall Triggered (Attempt %d): %v", i+1, err)
			time.Sleep(10 * time.Second)
			continue
		}
		
		// 🔊 พ่นผลลัพธ์ (Finisher)
		fmt.Printf("\n--- 🛡️ GRIPEN FINAL REPORT ---\n%s\n------------------------------\n", result)
		break
	}
}

// --- [🛠️ Helper Functions] ---

func handleGripenProcess(w http.ResponseWriter, r *http.Request) {
	// ตรวจสอบ Authorization
	secret := os.Getenv("A2A_SECRET_KEY")
	if r.Header.Get("X-ThitNuea-Auth") != secret {
		http.Error(w, "🚫 Unauthorized", http.StatusUnauthorized)
		return
	}

	var task struct {
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	res, err := askGripen(r.Context(), task.Content)
	if err != nil {
		fmt.Fprintf(w, `{"status": "error", "message": "%v"}`, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "completed", "result": res})
}

func calculateNextRun(now time.Time, hours []int, loc *time.Location) time.Time {
	for _, h := range hours {
		t := time.Date(now.Year(), now.Month(), now.Day(), h, 0, 0, 0, loc)
		if t.After(now) { return t }
	}
	return time.Date(now.Year(), now.Month(), now.Day()+1, hours[0], 0, 0, 0, loc)
}

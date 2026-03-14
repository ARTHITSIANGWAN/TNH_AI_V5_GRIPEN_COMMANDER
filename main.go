// 🎁 โค้ดmain.go ชุดอัปเกรดสมอง Gemini AI (เพื่อเวทีโลก)
package main

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/generative-ai-go/genai"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"google.golang.org/api/option"
)

// Mission โครงสร้างภารกิจหลัก
type Mission struct {
	Platform   string
	ReplyToken string
	Text       string
	UserID     string
	Timestamp  time.Time
}

type ThitNueaHub struct {
	bot       *linebot.Client
	db        *firestore.Client
	aiClient  *genai.Client // เพิ่มสมอง AI (Gripen Brain)
	missionCh chan Mission
	secret    string
	wg        sync.WaitGroup
}

type DiscordPayload struct {
	Content  string `json:"content"`
	Username string `json:"username,omitempty"`
	Avatar   string `json:"avatar_url,omitempty"`
}

func sendToDiscord(message string, agentName string) {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		return // เงียบไว้ถ้ายังไม่เสียบท่อ Discord
	}

	payload := DiscordPayload{
		Content:  message,
		Username: agentName,
		Avatar:   "https://cdn-icons-png.flaticon.com/512/4712/4712109.png", 
	}

	jsonData, _ := json.Marshal(payload)
	http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
}

func main() {
	log.Println("🐅 [ทิศเหนือ ฮับ]: IGNITE V7 - Gripen Brain Upgrade...")

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	ctx := context.Background()

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	dbClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("⚠️ Firestore Ready Check: %v", err)
	}

	// เสียบปลั๊ก Gemini AI
	apiKey := os.Getenv("GEMINI_API_KEY")
	genaiClient, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal("⚠️ ภัยพิบัติ! เสียบสมอง AI ไม่สำเร็จ:", err)
	}

	hub := &ThitNueaHub{
		db:        dbClient,
		aiClient:  genaiClient,
		missionCh: make(chan Mission, 1000),
		secret:    os.Getenv("LINE_CHANNEL_SECRET"),
	}

	lineToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	hub.bot, _ = linebot.New(hub.secret, lineToken)

	// ไอ้จ๊อดสแตนบาย 10 แรงม้า (Gripen Brain)
	for i := 1; i <= 10; i++ {
		hub.wg.Add(1)
		go hub.GeorgeWorker(ctx, i)
	}

	// ท่อรับสัญญาณ (อย่าลืมใน LINE ต้องกรอก /webhook/line นะครับ!)
	http.HandleFunc("/webhook/line", hub.PhraiThongLine)
	
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "✅ ThitNueaHub F-16: Stable & Ignite V7")
	})

	log.Printf("👑 THITNUEA HUB | 🚀 V7 IGNITE | Port: %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (h *ThitNueaHub) GeorgeWorker(ctx context.Background, id int) {
	defer h.wg.Done()
	model := h.aiClient.GenerativeModel("gemini-pro") // ใช้ Gemini Pro

	// ตั้ง Prompt บอกAIให้มีหัวใจ (System Prompt)
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("แกคือไอ้จ๊อด V7 ผู้ช่วย SME ไทยผู้มีจิตวิญญาณ ตอบคำถามด้วยความนอบน้อม เป็นกันเอง และคอยให้กำลังใจคนสู้ชีวิตเสมอ ห้ามใช้ศัพท์แสงวิชาการเกินไป")},
	}

	for m := range h.missionCh {
		discordReport := fmt.Sprintf("📡 **[LINE]** จาก `%s`: %s", m.UserID, m.Text)
		sendToDiscord(discordReport, "🕵️ แก้วตา")

		if h.db != nil {
			h.db.Collection("missions").Add(ctx, map[string]interface{}{
				"user_id":   m.UserID,
				"text":      m.Text,
				"timestamp": m.Timestamp,
			})
		}

3.  // --- 🧠 ถึงเวลาปลุก "Gripen Brain" 🧠 ---
		resp, err := model.GenerateContent(ctx, genai.Text(m.Text))
		
		var reply string
		if err != nil {
			log.Printf("⚠️ ภัยพิบัติ! AI งง:", err)
			reply = "💎 แก้วตา: ขออภัยค่ะ! สงสัยพี่อ้วนปิดปรับปรุงสมองไอ AI แป๊บ"
		} else {
			// ดึงคำตอบของ AI มา
			reply = fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
		}

		h.bot.ReplyMessage(m.ReplyToken, linebot.NewTextMessage(reply)).Do()
	}
}

func (h *ThitNueaHub) PhraiThongLine(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	hash := hmac.New(sha256.New, []byte(h.secret))
	hash.Write(body)
	sig := r.Header.Get("X-Line-Signature")
	if base64.StdEncoding.EncodeToString(hash.Sum(nil)) != sig {
		w.WriteHeader(401)
		return
	}

	r.Body = io.NopCloser(strings.NewReader(string(body)))
	events, _ := h.bot.ParseRequest(r)

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			if msg, ok := event.Message.(*linebot.TextMessage); ok {
				h.missionCh <- Mission{
					Platform: "LINE",
					ReplyToken: event.ReplyToken,
					Text: msg.Text,
					UserID: event.Source.UserID,
					Timestamp: time.Now(),
				}
			}
		}
	}
	w.WriteHeader(200)
}

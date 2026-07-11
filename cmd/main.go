package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Config พอร์ตและ Secret
const LISTEN_PORT = ":2026"
var SECURITY_SECRET = []byte("tnh-gripen-sovereign-secret-2026")

// Payload โครงสร้างข้อมูลหลัก
type GripenCommandPayload struct {
	CommandID string `json:"command_id"`
	Action    string `json:"action"`
	Squadron  string `json:"squadron"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

func main() {
	mux := http.NewServeMux()

	// 🏰 หน้าจอสถานะระบบ
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>🏰 TNH V84.9.2 GRIPEN SQUADRON ENGINE ONLINE</h1>"))
	})

	// 📡 ด่านตรวจคำสั่งยุทธศาสตร์ (Launch Endpoint)
	mux.HandleFunc("/api/v84/squadron/launch", handleLaunch)

	fmt.Printf("⚡ [Go Engine Sovereign]: ล็อกตำแหน่งที่พอร์ต %s\n", LISTEN_PORT)
	log.Fatal(http.ListenAndServe(LISTEN_PORT, mux))
}

func handleLaunch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var p GripenCommandPayload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}

	// 🛡️ ด่านตรวจสอบลายเซ็น
	if !validateHMAC(p, p.Signature) {
		http.Error(w, "Security Breach: Invalid Signature", http.StatusUnauthorized)
		return
	}

	// 🧹 รันงานในฐานะลูกน้อง L5 ไอ้จ๊อด
	go executeTask(p)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"status":"Dispatched to Local Engine"}`))
}

func validateHMAC(p GripenCommandPayload, sig string) bool {
	h := hmac.New(sha256.New, SECURITY_SECRET)
	h.Write([]byte(p.CommandID + p.Action + p.Squadron))
	return hex.EncodeToString(h.Sum(nil)) == sig
}

func executeTask(p GripenCommandPayload) {
	log.Printf("🧹 [L5 ไอ้จ๊อด]: ทำการศัลยกรรมล้างเศษขยะให้ Job %s เรียบร้อย", p.CommandID)
}

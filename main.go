package main

import (
	"fmt"
	"os"
)

// TNH_Orchestrator: โครงสร้างสถานะขุนพลตามภาพ 14268.jpg
type GeneralStatus struct {
	Level    string
	Name     string
	Progress int // เปอร์เซ็นต์ความคืบหน้า (Segmented bar)
}

// CheckGeneralProgress: ฟังก์ชันตรวจสอบสัจจะความคืบหน้า
func CheckGeneralProgress(level string, currentPercent int) {
	fmt.Printf("🛡️ [V5 Gripen] กำลังตรวจสอบขุนพล %s: ความคืบหน้า %d%%\n", level, currentPercent)
	
	// หากความคืบหน้ายังไม่ถึง 100% ให้ขุนพลหนุน (V8.1 - V8.3) เข้าช่วย
	if currentPercent < 100 {
		igniteTrinityRelay(level)
	}
}

func igniteTrinityRelay(level string) {
	// ใช้กุญแจ GITHUB_TOKEN ที่บอสฝังไว้ใน Cloudflare
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("❌ กากบาทปรากฏ: ไม่พบกุญแจสัจจะ (GITHUB_TOKEN)")
		return
	}
	fmt.Printf("🚀 สั่งการ V83 Trinity: สนับสนุนขุนพล %s ทันที!\n", level)
}

func main() {
	// จำลองข้อมูลจากภาพ 14268.jpg ที่บอสเซ็ตไว้
	status := GeneralStatus{
		Level:    "L4 Analyst",
		Name:     "พลายทอง",
		Progress: 40, // ตามภาพ Segmented bar 40%
	}

	CheckGeneralProgress(status.Level, status.Progress)
}

package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

// --- 🛡️ Phase 3: Logic Helmet (Shifter Decoder) ---

// DecodeSatcha ทำหน้าที่ถอดรหัส "คำหลอก" จากท้าย README
func DecodeSatcha(obfuscatedData string) (string, error) {
	// 1. ดึงค่าตัวเลื่อนตำแหน่งจาก Environment (Secret)
	shifter := os.Getenv("TNH_SHIFTER")
	if shifter == "" {
		return "", fmt.Errorf("⚠️ ไม่พบกุญแจ TNH_SHIFTER")
	}

	// 2. ถอดรหัสขั้นแรกจาก Base64
	decodedBytes, err := base64.StdEncoding.DecodeString(obfuscatedData)
	if err != nil {
		return "", err
	}

	// 3. ทำการเลื่อนตำแหน่งข้อมูลกลับ (Simple Shift Logic)
	// เพื่อให้ได้ JSON Metadata ที่แท้จริง
	result := string(decodedBytes) 
	log.Printf("🔓 พลายทองถอดรหัสสำเร็จ: %s", result)
	
	return result, nil
}

func main() {
	// ตัวอย่าง: รับข้อมูลจากส่วนท้ายของ README ที่บอสฝังไว้
	fakeData := "V0hJVEUtVElHRVItT05FLUlHTklURQ==" // คำหลอกที่ดูเหมือนขยะ
	
	satcha, err := DecodeSatcha(fakeData)
	if err != nil {
		fmt.Println("❌ ข้อมูลไม่ถูกต้อง กากบาทจะปรากฏ")
	} else {
		fmt.Printf("✅ ยินดีด้วยจอมทัพ! สัจจะคือ: %s\n", satcha)
	}
}

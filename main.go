package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// --- 🛡️ 11 ขุนพล: ระบบรับคำสั่งข้ามมิติ ---

// ReadCommandFromImage ทำหน้าที่สกัดคำสั่งที่ซ่อนอยู่ท้ายรูปภาพ
func ReadCommandFromImage(imgData []byte) string {
	// ค้นหาตำแหน่งของ "CMD:" ที่บอสฝังไว้ท้ายไฟล์
	marker := []byte("CMD:")
	idx := bytes.Index(imgData, marker)
	if idx == -1 {
		return "❌ ไม่พบสัจจะในรูปภาพ"
	}

	// สกัดคำสั่งออกมา (Zero-Garbage Logic)
	command := string(imgData[idx+4:])
	return strings.TrimSpace(command)
}

// FetchCommandFromIssue (ตัวอย่าง) สำหรับดึงสัจจะจาก GitHub Issue
func FetchCommandFromIssue(issueBody string) {
	// ดึง manifesto_id และตรวจสอบสิทธิด้วย HMAC
	fmt.Printf("🕵️ พลายทอง ตรวจพบภารกิจจาก Issue: %s\n", issueBody)
}

func main() {
	// ตัวอย่าง: จำลองข้อมูลรูปภาพที่มีคำสั่งซ่อนอยู่
	mockImage := append([]byte{0xFF, 0xD8, 0xFF}, []byte("\nCMD:IGNITE_V83_EMPIRE")...)
	
	cmd := ReadCommandFromImage(mockImage)
	fmt.Printf("🚀 สัญญาณที่ได้รับ: %s\n", cmd)
}

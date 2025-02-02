package main

import (
	"fmt"
	"log"

	gosseract "github.com/otiai10/gosseract/v2"
)

func main() {
	// 创建一个新的 Tesseract 客户端
	client := gosseract.NewClient()
	defer client.Close()

	if err := client.SetLanguage("eng+chi_sim+chi_tra"); err != nil {
		panic(err)
	}
	// client.SetLanguage("chi_sim")

	// 设置图片文件路径
	err := client.SetImage("/root/iShot_2025-02-02_21.54.44.png")
	if err != nil {
		log.Fatalf("设置图片失败: %v", err)
	}

	// 获取识别结果
	text, err := client.Text()
	if err != nil {
		log.Fatalf("OCR 识别失败: %v", err)
	}

	fmt.Println("识别结果：")
	fmt.Println(text)
}

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
)

func convertANSIToUTF8(filePath string) error {
	// ファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ファイルを開けません: %v", err)
	}
	defer file.Close()

	// ANSI (Windows-1252) から UTF-8 への変換
	reader := transform.NewReader(bufio.NewReader(file), charmap.Windows1252.NewDecoder())

	// 変換された内容を読み取る
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("ファイル内容の読み取り中にエラーが発生しました: %v", err)
	}

	// ファイルを再度開いて書き込みモードにする
	file, err = os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ファイルを再作成できません: %v", err)
	}
	defer file.Close()

	// 変換された内容を書き込む
	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("ファイルへの書き込み中にエラーが発生しました: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("使用方法: convert <ファイル>")
		return
	}

	filePath := os.Args[1]

	err := convertANSIToUTF8(filePath)
	if err != nil {
		log.Fatalf("変換中にエラーが発生しました: %v", err)
	}

	fmt.Println("変換が完了しました。")
}

package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	// 引数のlogファイルに書き込むか、存在しなければ作成する、追記するかを設定する
	// パーミッションを0666に設定
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	// 書き込み先を標準出力とログファイルに設定している
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// ログのフォーマット設定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//
	log.SetOutput(multiLogFile)
}

package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/mori5600/gotodo/common"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// lumberjack の設定
	logRotator := &lumberjack.Logger{
		Filename:   common.LOG_PATH, // ログファイルのパス
		MaxSize:    10,              // ログファイルの最大サイズ（MB）
		MaxBackups: 5,               // 保持するバックアップファイルの最大数
		MaxAge:     0,               // ログファイルの最大保存日数（0は無制限）
		Compress:   false,           // 古いログファイルの圧縮有無
	}

	// 複数の出力先（標準出力とファイル）を設定
	multiWriter := io.MultiWriter(os.Stdout, logRotator)

	// slog のハンドラーを設定（JSON形式で出力）
	handler := slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		Level: slog.LevelInfo, // INFOレベル以上を出力
	})

	// カスタムロガーを作成
	logger := slog.New(handler)

	// デフォルトのロガーを設定
	slog.SetDefault(logger)

	// ログ出力の例
	slog.Info("アプリケーションが起動しました", "version", "1.0.0", "port", 8080)
	slog.Warn("高いメモリ使用率", "memory_usage", 90.5)
	slog.Error("データベース接続に失敗しました", "error", "timeout")
}

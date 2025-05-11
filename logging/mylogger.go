package logging

import (
	"io"
	"log/slog"
	"os"
	"sync"

	"github.com/mori5600/gotodo/common"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *slog.Logger
	once   sync.Once
)

// GetLogger はシングルトンの slog.Logger インスタンスを返します。
func GetLogger() *slog.Logger {
	once.Do(func() {
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

		// slog のハンドラーを設定（テキスト形式で出力）
		handler := slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
			Level: slog.LevelDebug, // DEBUGレベル以上を出力
		})

		// ロガーを作成
		logger = slog.New(handler)
	})

	return logger
}

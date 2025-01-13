package utils

import (
	"bytes"
	"sync"
)

// LogWriter 實現 io.Writer 接口
type LogWriter struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

// Write 方法將日誌寫入緩衝區
func (lw *LogWriter) Write(p []byte) (n int, err error) {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	return lw.buf.Write(p)
}

// GetLogs 方法返回緩衝區中的日誌內容
func (lw *LogWriter) GetLogs() string {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	return lw.buf.String()
}

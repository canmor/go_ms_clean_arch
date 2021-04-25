package logger

import (
	"github.com/canmor/go_ms_clean_arch/pkg/domain"
	"log"
)

type Logger struct {
}

func (l Logger) Log(_ domain.LogLevel, content string) {
	log.Printf(content)
}

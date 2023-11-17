package server

import (
	"fmt"
	"github.com/goccy/go-json"
	"go.uber.org/zap"
	"strings"
)

// Info info级别的日志
func (s *InnerStdLogger) Info(msg string) *StdLogger {
	return &StdLogger{
		Logger:  zap.NewNop(),
		message: msg,
		level:   Info,
	}
}
func (s *InnerStdLogger) Debug(msg string) *StdLogger {
	return &StdLogger{
		Logger:  zap.NewNop(),
		message: msg,
		level:   Debug,
	}
}
func (s *InnerStdLogger) Panic(msg string) *StdLogger {
	return &StdLogger{
		Logger:  zap.NewNop(),
		message: msg,
		level:   Panic,
	}
}
func (s *InnerStdLogger) Warning(msg string) *StdLogger {
	return &StdLogger{
		Logger:  zap.NewNop(),
		message: msg,
		level:   Warning,
	}
}
func (s *InnerStdLogger) Error(msg string) *StdLogger {
	return &StdLogger{
		Logger:  zap.NewNop(),
		message: msg,
		level:   Error,
	}
}

func (s *InnerStdLogger) Fatal(msg string) *StdLogger {
	return &StdLogger{
		Logger:  zap.NewNop(),
		message: msg,
		level:   Fatal,
	}
}

func (s *StdLogger) Set(k string, v any) *StdLogger {
	msg := fmt.Sprintf("[%s: %v]", k, v)
	s.printVal = append(s.printVal, msg)
	return s
}

func (s *StdLogger) SetError(err error) *StdLogger {
	s.err = err
	return s
}

func (s *StdLogger) SetStruct(k string, v any) {
	val, err := json.MarshalNoEscape(v)
	if err != nil {
		return
	}
	s.Set(k, string(val))
}

func (s *StdLogger) SetMap(msg map[string]any) {
	if len(msg) <= 0 {
		return
	}
	for k, v := range msg {
		s.Set(k, v)
	}

}

func (s *StdLogger) Print() {
	msg := fmt.Sprintf("[msg: %s][%s: %v]%s", s.message, errKey, s.err, strings.Join(s.printVal, ""))
	switch s.level {
	case Debug:
		s.Debug(msg)
		s.Sync()
		return
	case Error:
		s.Error(msg)
		return
	case Info:
		s.Info(msg)
		s.Sync()
		return
	case Warning:
		s.Warn(msg)
		return
	case Panic:
		s.Panic(msg)
		return
	case Fatal:
		s.Fatal(msg)
		return
	}
}

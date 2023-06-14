package test

import (
	"bytes"
	"net/url"

	"go.uber.org/zap"
)

type MemorySink struct {
	*bytes.Buffer
}

func (m *MemorySink) Close() error { return nil }
func (m *MemorySink) Sync() error  { return nil }

func SinkTestLogger() (*zap.Logger, *MemorySink, error) {
	buff := new(bytes.Buffer)

	sink := &MemorySink{Buffer: buff}
	errR := zap.RegisterSink("memory", func(url *url.URL) (zap.Sink, error) {
		return sink, nil
	})
	if errR != nil {
		return nil, nil, errR
	}

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"memory://"}
	logger, errL := cfg.Build()
	if errL != nil {
		return nil, nil, errL
	}
	return logger, sink, nil
}

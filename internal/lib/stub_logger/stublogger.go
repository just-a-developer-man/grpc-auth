package stub_logger

import (
	"context"
	"log/slog"
)

// type Handler interface {
// 	Enabled(context.Context, Level) bool
// 	Handle(context.Context, Record) error
// 	WithAttrs(attrs []Attr) Handler
// 	WithGroup(name string) Handler
// }

type StubHandler struct{}

func NewStubLogger() *slog.Logger {
	return slog.New(NewStubHandler())
}

func NewStubHandler() *StubHandler {
	return &StubHandler{}
}

func (h *StubHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *StubHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *StubHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *StubHandler) WithGroup(_ string) slog.Handler {
	return h
}

package interfaces

// Logger is
type Logger interface {
	Log(message string)
}

// 1. Kullanım Şekli

// Server is
type Server struct {
	logger Logger
}

// Console is
type Console struct {
	logger Logger
}

// 2. Kullanım Şekli

// Process is
func Process(logger Logger) {
	logger.Log("Hello")
}

package logx

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name        string
		level       string
		development bool
	}{
		{
			name:        "debug level in development mode",
			level:       "DEBUG",
			development: true,
		},
		{
			name:        "info level in production mode",
			level:       "INFO",
			development: false,
		},
		{
			name:        "warning level in production mode",
			level:       "WARNING",
			development: false,
		},
		{
			name:        "error level in production mode",
			level:       "ERROR",
			development: false,
		},
		{
			name:        "invalid level defaults to warn",
			level:       "INVALID",
			development: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewLogger(tt.level, tt.development)
			require.NotNil(t, logger)
		})
	}
}

func TestDefaultLogger(t *testing.T) {
	// First call should initialize the logger
	logger1 := DefaultLogger()
	require.NotNil(t, logger1)

	// Second call should return the same logger
	logger2 := DefaultLogger()
	require.NotNil(t, logger2)
	assert.Same(t, logger1, logger2)
}

func TestLevelToZapLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected zapcore.Level
	}{
		{levelDebug, zapcore.DebugLevel},
		{levelInfo, zapcore.InfoLevel},
		{levelWarning, zapcore.WarnLevel},
		{levelError, zapcore.ErrorLevel},
		{levelCritical, zapcore.DPanicLevel},
		{levelAlert, zapcore.PanicLevel},
		{levelEmergency, zapcore.FatalLevel},
		{"invalid", zapcore.WarnLevel},    // Default to warn for invalid levels
		{"  DEBUG  ", zapcore.DebugLevel}, // Test trimming and case insensitivity
		{"info", zapcore.InfoLevel},       // Test case insensitivity
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			level := levelToZapLevel(tt.input)
			assert.Equal(t, tt.expected, level)
		})
	}
}

func TestLoggerOutput(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer

	// Create a custom core that writes to our buffer
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(productionEncoderConfig),
		zapcore.AddSync(&buf),
		zapcore.InfoLevel,
	)

	// Create a logger with our custom core
	logger := zap.New(core).Sugar()

	// Log a message
	logger.Info("test message")

	// Parse the JSON output
	var logMap map[string]interface{}
	err := json.Unmarshal(buf.Bytes(), &logMap)
	require.NoError(t, err)

	// Verify the log message
	assert.Equal(t, "test message", logMap[message])
	assert.Equal(t, levelInfo, logMap[severity])
	assert.NotEmpty(t, logMap[timestamp])
}

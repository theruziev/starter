# AI Assistant Conventions for Starter Project

## Overview
This document provides guidelines for AI assistants working with the Starter project. Following these conventions ensures consistency with the project's structure and patterns.

## Project Structure
```
├── cmd/                  # Application entry points
├── internal/             # Internal packages
│   ├── app/              # Application runners
│   │   └── server/       # Server runner
│   └── pkg/              # Internal packages
│       ├── closer/       # Resource management
│       ├── logx/         # Logging functionality
│       ├── info/         # Version information
│       └── healthcheck/  # Health check endpoints
├── pkg/                  # Public packages
```

## Key Libraries
- **CLI**: [alecthomas/kong](https://github.com/alecthomas/kong)
- **Environment Variables**: [joho/godotenv](https://github.com/joho/godotenv)
- **HTTP Router**: [go-chi/chi](https://github.com/go-chi/chi/v5)
- **Logging**: [uber-go/zap](https://go.uber.org/zap) (via logx package)
- **Testing**: [stretchr/testify](https://github.com/stretchr/testify)

## Code Conventions

### Error Handling
- Always wrap errors with context: `fmt.Errorf("failed to do something: %w", err)`
- Use appropriate error types and check with `errors.Is()` or `errors.As()`

### Context Management
- Pass context through function calls for cancellation
- Use context for carrying request-scoped values like loggers
- Example: `ctx, cancel := context.WithTimeout(parentCtx, timeout)`

### Logging
Logging example:
```
// Creating a logger
logger := logx.NewLogger(level, isDebug)

// Adding logger to context
ctx = logx.WithLogger(ctx, logger)

// Getting logger from context
logger := logx.FromContext(ctx)

// Using the logger
logger.Info("message", zap.String("key", "value"))
logger.Error("error occurred", zap.Error(err))
```

### Resource Management
Resource management example:
```
// Create a closer
c := closer.New()

// Add resources to be closed
c.Add(func(ctx context.Context) error {
    return someResource.Close()
})

// Close all resources
if err := c.Close(ctx); err != nil {
    // Handle error
}
```

## Testing Guidelines

### Using testify
Testify examples:
```
// Fatal assertions (stops test on failure)
require.NoError(t, err)
require.Equal(t, expected, actual)
require.NotNil(t, object)

// Non-fatal assertions (continues test after failure)
assert.Equal(t, expected, actual)
assert.ErrorContains(t, err, "expected message")
assert.True(t, condition)
```

### Table-Driven Tests
Table-driven test example:
```
func TestFunction(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected string
        wantErr  bool
    }{
        {
            name:     "valid case",
            input:    "valid",
            expected: "result",
            wantErr:  false,
        },
        {
            name:     "error case",
            input:    "invalid",
            expected: "",
            wantErr:  true,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result, err := functionUnderTest(tc.input)
            
            if tc.wantErr {
                require.Error(t, err)
            } else {
                require.NoError(t, err)
                assert.Equal(t, tc.expected, result)
            }
        })
    }
}
```

### HTTP Testing
HTTP testing example:
```
func TestHTTPHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/path", nil)
    w := httptest.NewRecorder()
    
    handler.ServeHTTP(w, req)
    
    resp := w.Result()
    defer resp.Body.Close()
    
    require.Equal(t, http.StatusOK, resp.StatusCode)
    
    body, err := io.ReadAll(resp.Body)
    require.NoError(t, err)
    assert.JSONEq(t, `{"key":"value"}`, string(body))
}
```

## AI Assistant Guidelines

### When Generating Code
1. **Follow Project Structure**: Place new code in the appropriate package
2. **Use Existing Patterns**: Check similar functionality in the codebase
3. **Include Error Handling**: Wrap errors with context
4. **Add Tests**: Write tests for all new functionality
5. **Use Context**: Pass context through function calls
6. **Document Code**: Add comments for complex logic

### When Suggesting Changes
1. **Minimal Changes**: Make the smallest change needed to solve the problem
2. **Maintain Compatibility**: Avoid breaking existing functionality
3. **Consider Performance**: Be mindful of performance implications
4. **Follow Style**: Match the existing code style and patterns
5. **Explain Reasoning**: Provide clear explanations for your changes

### Common Patterns to Follow
1. **HTTP Endpoints**: Include health checks and version endpoints
2. **Graceful Shutdown**: Implement proper shutdown for all resources
3. **Configuration**: Use environment variables with sensible defaults
4. **Dependency Injection**: Use interfaces for better testability
5. **Error Responses**: Standardize error responses in HTTP handlers

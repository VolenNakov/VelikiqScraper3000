package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

// Colors for different status code ranges
const (
	green   = "\033[32m"
	yellow  = "\033[33m"
	red     = "\033[31m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	reset   = "\033[0m"
)

// StatusEmoji maps HTTP status codes to relevant emojis
var StatusEmoji = map[int]string{
	200: "✅",
	201: "🆕",
	204: "⭕",
	304: "💾",
	400: "⚠️",
	401: "🔒",
	403: "🚫",
	404: "❓",
	500: "💥",
	502: "🌋",
	503: "🔧",
}

type bodyDumpResponseWriter struct {
	io.Writer
	echo.Response
}

// EnhancedLogger is a custom middleware that provides detailed request/response logging
func EnhancedLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			// Create a buffer to store the response body
			resBody := new(bytes.Buffer)
			if req.Header.Get(echo.HeaderContentType) == echo.MIMEApplicationJSON {
				res.Writer = &bodyDumpResponseWriter{Writer: io.MultiWriter(res.Writer, resBody), Response: *res}
			}

			// Get request body if it's JSON
			var reqBody string
			if req.Header.Get(echo.HeaderContentType) == echo.MIMEApplicationJSON {
				body, _ := io.ReadAll(req.Body)
				req.Body = io.NopCloser(bytes.NewBuffer(body))
				if len(body) > 0 {
					var prettyJSON bytes.Buffer
					if err := json.Indent(&prettyJSON, body, "", "  "); err == nil {
						reqBody = prettyJSON.String()
					}
				}
			}

			// Process the request
			err := next(c)

			// Calculate duration
			stop := time.Now()
			duration := stop.Sub(start)

			// Get status code and corresponding color
			status := res.Status
			var statusColor string
			switch {
			case status >= 500:
				statusColor = red
			case status >= 400:
				statusColor = yellow
			case status >= 300:
				statusColor = blue
			case status >= 200:
				statusColor = green
			default:
				statusColor = magenta
			}

			// Get emoji for status code
			emoji := StatusEmoji[status]
			if emoji == "" {
				emoji = "📋"
			}

			// Build the log message
			logMsg := fmt.Sprintf("\n%s╭─── Request %s %s\n", blue, time.Now().Format("2006-01-02 15:04:05.000"), reset)
			logMsg += fmt.Sprintf("%s├ %s %s %s\n", blue, req.Method, req.URL.Path, reset)
			if len(req.URL.RawQuery) > 0 {
				logMsg += fmt.Sprintf("%s├ Query: %s\n", blue, req.URL.RawQuery)
			}
			if reqBody != "" {
				logMsg += fmt.Sprintf("%s├ Request Body:\n%s\n", blue, reqBody)
			}
			logMsg += fmt.Sprintf("%s├ Remote IP: %s\n", blue, c.RealIP())
			logMsg += fmt.Sprintf("%s├ User Agent: %s\n", blue, req.UserAgent())

			// Response information
			logMsg += fmt.Sprintf("%s├─── Response\n", blue)
			logMsg += fmt.Sprintf("%s├ Status: %s%d %s %s\n", blue, statusColor, status, emoji, reset)
			logMsg += fmt.Sprintf("%s├ Latency: %v\n", blue, duration)

			if resBody.Len() > 0 {
				var prettyJSON bytes.Buffer
				if err := json.Indent(&prettyJSON, resBody.Bytes(), "", "  "); err == nil {
					logMsg += fmt.Sprintf("%s├ Response Body:\n%s\n", blue, prettyJSON.String())
				}
			}

			// Error if any
			if err != nil {
				logMsg += fmt.Sprintf("%s├ Error: %v\n", blue, err)
			}

			logMsg += fmt.Sprintf("%s╰───\n", blue)

			// Print the final log message
			fmt.Print(logMsg)

			return err
		}
	}
}

// Custom writer for response body capture
func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

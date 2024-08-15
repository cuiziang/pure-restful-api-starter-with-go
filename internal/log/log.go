package log

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status       int
	responseSize int64
	body         *bytes.Buffer
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.responseSize += int64(size)
	rw.body.Write(b)
	return size, err
}

type Entry struct {
	Timestamp    string      `json:"timestamp"`
	RemoteAddr   string      `json:"remote_addr"`
	Method       string      `json:"method"`
	Path         string      `json:"path"`
	Protocol     string      `json:"protocol"`
	Status       int         `json:"status"`
	Duration     string      `json:"duration"`
	ResponseSize int64       `json:"response_size"`
	UserAgent    string      `json:"user_agent"`
	RequestBody  interface{} `json:"request_body,omitempty"`
	ResponseBody interface{} `json:"response_body,omitempty"`
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var requestBody any
		if r.Body != nil {
			bodyBytes, _ := io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			json.Unmarshal(bodyBytes, &requestBody)
		}

		buf := &bytes.Buffer{}
		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
			body:           buf,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		var responseBody any
		err := json.Unmarshal(rw.body.Bytes(), &responseBody)
		if err != nil {
			return
		}

		logEntry := Entry{
			Timestamp:    time.Now().Format(time.RFC3339),
			RemoteAddr:   r.RemoteAddr,
			Method:       r.Method,
			Path:         r.URL.Path,
			Protocol:     r.Proto,
			Status:       rw.status,
			Duration:     duration.String(),
			ResponseSize: rw.responseSize,
			UserAgent:    r.UserAgent(),
			RequestBody:  requestBody,
			ResponseBody: responseBody,
		}

		logJSON, err := json.Marshal(logEntry)
		if err != nil {
			log.Printf("Error marshaling log entry: %v", err)
			return
		}

		log.Println(string(logJSON))
	})
}

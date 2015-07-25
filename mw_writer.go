package Frodo

import (
	"net/http"
	"time"
)

// MiddlewareResponseWriter is used to hijack/embed http.ResponseWriter
// thus making it satisfy the ResponseWriter interface, we then add a written boolean property
// to trace when a write made and exit
type MiddlewareResponseWriter struct {
	http.ResponseWriter
	written   bool
	timeStart time.Time
	method    string
	route     string
}

// Write writes data back the client/creates the body
func (w *MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	w.written = true
	Log.Info("An application response was written back: %v\n", w.written)
	Log.Success("|%s| %v | %v - %s", w.method, http.StatusOK, time.Now().Sub(w.timeStart), w.route)
	return w.ResponseWriter.Write(bytes)
}

// WriteHeader is in charge of building the Header file and writing it back to the client
func (w *MiddlewareResponseWriter) WriteHeader(code int) {
	w.written = true
	Log.Info("Header has been written: %v", w.written)
	w.ResponseWriter.WriteHeader(code)
}

/*
   FUTURE:

   If a write happens, then we can track it here and we can place our
   AFTER middleware just before a write happens

   func (b *basicWriter) Write(buf []byte) (int, error) {
		b.WriteHeader(http.StatusOK)
		n, err := b.ResponseWriter.Write(buf)
		if b.tee != nil {
			_, err2 := b.tee.Write(buf[:n])
			// Prefer errors generated by the proxied writer.
			if err == nil {
				err = err2
			}
		}
		b.bytes += n
		return n, err
   }
*/

package main

import (
    "net/http"
)


// Chaining Model of http.handlers
type Middleware []http.Handler


func (m *Middleware)Add(handler http.Handler) ( *Middleware){
    *m = append(*m, handler)
    return m
}

func (m Middleware)ServeHTTP(w http.ResponseWriter, r *http.Request){
    mw := NewMiddlewareResponseWriter(w)
    for _,handler := range m {
        if mw.written {
            return
        }
        // write to mw
        handler.ServeHTTP(mw,r)
    }
    http.NotFound(w,r)
}


// MiddlewareResponseWriter is a http.ResonseWriter implementation,
// Create a NewMiddlewareResponseWriter from a http.ResponseWriter.
type MiddlewareResponseWriter struct {
    http.ResponseWriter
    written bool
}



func NewMiddlewareResponseWriter(w http.ResponseWriter) (*MiddlewareResponseWriter){
    return &MiddlewareResponseWriter{
        ResponseWriter: w,
        // written: false, as default
    }
}

// implement io.Writer.
// whenever Write() is called, written = true.
// Handle bytes to be written by passing toward http.ResponseWriter.Write()
func (w *MiddlewareResponseWriter)Write( bytes []byte) (int,error){
    w.written = true
    return w.ResponseWriter.Write(bytes)
}

func (w *MiddlewareResponseWriter)WriteHeader(code int){
    w.written = true
    w.ResponseWriter.WriteHeader(code)
}

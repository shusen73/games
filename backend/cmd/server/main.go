package main

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	uiembed "shusen.tech/game"
)

func main() {
	// --- Mode & router -------------------------------------------------------
	r := gin.New()
	r.Use(gin.Recovery()) // keeps the server alive on panics

	// --- Filesystem setup ----------------------------------------------------
	ui, err := fs.Sub(uiembed.StaticFS, "static")
	if err != nil {
		log.Fatalf("embed FS missing 'static': %v", err)
	}
	fileServer := http.FileServer(http.FS(ui))

	// --- Health endpoint -----------------------------------------------------
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// SSE endpoint
	r.GET("/status", func(c *gin.Context) {
		// Set SSE headers
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		flusher, ok := c.Writer.(http.Flusher)
		if !ok {
			c.String(http.StatusInternalServerError, "streaming unsupported")
			return
		}

		// Send status every 5 seconds
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-c.Request.Context().Done():
				// Client closed connection
				log.Println("SSE client disconnected")
				return
			case t := <-ticker.C:
				status := "online"
				fmt.Fprintf(c.Writer, "data: %s\n\n", status)
				flusher.Flush()
				log.Printf("pushed status at %s", t.Format(time.RFC3339))
			}
		}
	})

	// (Optional) cache headers for immutable build assets
	r.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/assets/") {
			// Vite fingerprints assets; safe to cache "forever"
			c.Header("Cache-Control", "public, max-age=31536000, immutable")
		}
	})

	// --- SPA static serving + fallback --------------------------------------
	r.NoRoute(func(c *gin.Context) {
		p := strings.TrimPrefix(c.Request.URL.Path, "/")
		if p == "" {
			p = "index.html"
		}

		// If the path maps to a real file, delegate to FileServer.
		if f, err := ui.Open(p); err == nil {
			_ = f.Close()
			fileServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		// Otherwise, serve index.html as the SPA fallback.
		f, err := ui.Open("index.html")
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		defer f.Close()
		c.Header("Content-Type", "text/html; charset=utf-8")
		_, _ = io.Copy(c.Writer, f)
	})

	// --- HTTP server with timeouts & graceful shutdown ----------------------
	srv := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      0, // allow streaming
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Println("Game server listening on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Listen for termination signals and shut down gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
}

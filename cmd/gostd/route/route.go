package route

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// StartServer はHTTPサーバーを起動します。
func StartServer() {
	router := newRouter()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  0,
		WriteTimeout: 0,
		IdleTimeout:  0,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server is running on port %s", srv.Addr)
	gracefulShutdown(srv)
}

// newRouter はGinのルーターを初期化します。
func newRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Healthy"})
	})

	return r
}

// gracefulShutdown はサーバーの優雅なシャットダウンを実行します。
func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}

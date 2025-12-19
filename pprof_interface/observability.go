// observability/observability.go
package observability

import (
	"context"
	"log"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/google/gops/agent"
)

var (
	pprofShutdown func(context.Context) error
)

// StartGops 启动gops诊断工具
func StartGops() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Printf("[observability] 启动gops失败: %v", err)
		return
	}
	log.Printf("[observability] gops已启动")
}

// StartPProf 启动pprof性能分析
// addr为空时使用默认地址 127.0.0.1:1234
func StartPProf(addr string) {
	if addr == "" {
		addr = "127.0.0.1:1234"
	}

	shutdown, err := startPProfServer(addr)
	if err != nil {
		log.Printf("[observability] 启动pprof失败: %v", err)
		return
	}

	pprofShutdown = shutdown
	log.Printf("[observability] pprof监听在 http://%s", addr)
}

// CleanAllObserv 清理所有观测性资源
func CleanAllObserv() {
	// 关闭pprof
	if pprofShutdown != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := pprofShutdown(ctx); err != nil {
			log.Printf("[observability] 关闭pprof失败: %v", err)
		} else {
			log.Printf("[observability] pprof已关闭")
		}
	}
}

// 内部函数：启动pprof服务器
func startPProfServer(addr string) (func(context.Context) error, error) {
	mux := http.NewServeMux()

	// 注册pprof路由
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// 启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[observability] pprof服务错误: %v", err)
		}
	}()

	return server.Shutdown, nil
}

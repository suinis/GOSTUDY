package observability

import (
	"context"
	"expvar"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

// Config 控制观测端口与开关，便于在其他项目直接复用。
type Config struct {
	Addr string // 监听地址，留空则使用默认 127.0.0.1:6060
}

// Start 启动观测端口，注册 pprof、trace、expvar 等处理器。
// 返回优雅关闭函数，便于主程序退出时清理资源。
func Start(cfg Config) (func(ctx context.Context) error, error) {
	addr := cfg.Addr
	if addr == "" {
		addr = "127.0.0.1:6060"
	}

	mux := http.NewServeMux()

	// pprof 相关路由
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile) // CPU
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	// expvar 暴露 Go 运行时统计（便于 Prometheus 抓取或临时查看）
	mux.Handle("/debug/vars", expvar.Handler())

	// 基础健康检查
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[observability] listen %s failed: %v", addr, err)
		}
	}()

	return srv.Shutdown, nil
}

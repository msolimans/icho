package servers

import (
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/msolimans/icho"
	"github.com/prometheus/common/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

var (
	echoEndpoint = flag.String("echo_endpoint", ":9999", "endpoint of YourService")
)

func StartRest(ctx context.Context) error {
	//add caller to context
	md := metadata.Pairs("caller", "1")
	ctx = metadata.NewOutgoingContext(ctx, md)

	opts := []grpc.DialOption{grpc.WithInsecure()}

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true}),
	)

	if err := icho.RegisterIchoServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts); err != nil {
		return err
	}

	s := &http.Server{
		Addr:    ":8888",
		Handler: allowCORS(mux),
	}

	go func() {
		<-ctx.Done()
		log.Info("Shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			log.Error("Failed to shutdown http server: ", err)
		}
	}()

	return s.ListenAndServe()
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

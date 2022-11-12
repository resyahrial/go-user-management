package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var StandardSignals []os.Signal = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

type Action struct {
	Start    func(context.Context)
	Shutdown func(context.Context)
}

// RunHttpServer runs grpc server using graceful shutdown mechanism. RunGgrpcServer is a blocking action.
func RunHttpServer(ctx context.Context, server *http.Server, gracefulPeriod time.Duration) {
	RunHttpServerWithCustomSignal(ctx, server, gracefulPeriod, StandardSignals...)
}

func RunHttpServerWithCustomSignal(ctx context.Context, server *http.Server, gracefulPeriod time.Duration, signals ...os.Signal) {

	action := Action{
		Start: func(c context.Context) {
			if errStart := server.ListenAndServe(); errStart != nil {
				log.Fatal(errStart.Error())
			}
		},
		Shutdown: func(c context.Context) {
			if errShutdown := server.Shutdown(c); errShutdown != nil {
				log.Print(errShutdown.Error())
			}
		},
	}

	Run(context.Background(),
		gracefulPeriod,
		map[string]Action{
			"HTTP-SERVER": action,
		},
		StandardSignals...)
}

// Run executes long-running jobs with automatic execution of shutdown actions. Once signal is caught, it will apply
// shutdown actions to all long-running jobs. Run is a blocking action.
// If you intend to use this. DO NOT do panic()/log.Fatal inside your Action func(). Otherwise, it will force the
// application to exit, ignoring the graceful shutdown
func Run(ctx context.Context, gracefulPeriod time.Duration, actionPairs map[string]Action, signals ...os.Signal) {
	var wg sync.WaitGroup

	for key, action := range actionPairs {
		log.Printf("Starting %s .....", key)
		go action.Start(ctx)
	}

	ctx, stop := signal.NotifyContext(ctx, signals...)
	defer stop()
	<-ctx.Done()

	for key, action := range actionPairs {
		log.Printf("Shutting down %s gracefully", key)

		keyItem := key
		actItem := action

		wg.Add(1)
		go func() {
			defer wg.Done()
			shutdownCtx, cancelCtx := context.WithTimeout(context.Background(), gracefulPeriod)
			defer cancelCtx()
			actItem.Shutdown(shutdownCtx)

			if err := shutdownCtx.Err(); err == context.DeadlineExceeded {
				log.Printf("%s shutdown: timeout %d ms has been elapsed, force exit", keyItem, gracefulPeriod.Milliseconds())
			}
		}()

	}

	wg.Wait()
}

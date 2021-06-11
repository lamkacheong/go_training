// brother server package provide BrotherServers data struct
// which can hold servers that serve and die together
// and can capture signals sent to this process

package brother_server

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type BrotherServers struct {
	serversMap *map[string]http.Server
	errorGroup *errgroup.Group
	cancel     context.CancelFunc
	context    context.Context
}

func Build(serversMap *map[string]http.Server) BrotherServers {
	b := BrotherServers{}
	b.serversMap = serversMap
	b.errorGroup, b.context = errgroup.WithContext(context.Background())
	b.context, b.cancel = context.WithCancel(b.context)
	return b
}

func (b BrotherServers) Run() error {
	for key, value := range *b.serversMap {
		serverName := key //avoid data race
		server := value   //avoid data race
		b.errorGroup.Go(
			func() error {
				go func() {
					select {
					case <-b.context.Done():
						fmt.Printf("gracefully stop server: %v\n", serverName)
						server.Shutdown(b.context)
					}
				}()
				err := server.ListenAndServe()
				b.cancel()
				return err
			})
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Printf("Capture system signal %v\n", s)
				b.cancel()
			}
		}
	}()

	return b.errorGroup.Wait()
}

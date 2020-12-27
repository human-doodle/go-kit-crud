package main

import (
	"context"
	"net/http"

	"flag"
	"fmt"

	"github.com/go-kit/kit/log"
	mgo "gopkg.in/mgo.v2"

	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log/level"

	"crud.com/account"
)

func main() {

	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// if err != nil {
	// 	level.Error(logger).Log("exit", err)
	// 	os.Exit(-1)
	// }

	var db *mgo.Database
	db = GetMongoDB()
	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	{
		// db := GetMongoDB()
		repository, _ := account.NewRepo(db, logger)

		srv = account.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := account.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}

package main

import (
	"context"
	
	"flag"
	"fmt"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"crud.com/account"
)

//change according to mongo
// const dbsource = ""
var db *mongo.Client
var blogdb *mongo.Collection
var mongoCtx context.Context

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

	//chnage mongo
	/*var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", dbsource) //chnge
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}*/

	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	{
		repository := account.NewRepo(db, logger)

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

	fmt.Println("Connecting to MongoDB...")
  
 
	mongoCtx = context.Background()
  
  
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
  
	if err != nil {
		log.Fatal(err)
	}
  
  
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}
}

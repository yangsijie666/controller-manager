package main

import (
	"context"
	"github.com/go-logr/zapr"
	"github.com/yangsijie666/controller-manager/pkg/manager"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent"
	"go.uber.org/zap"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

/**
* @Author: yangsijie666
* @Date: 2023/5/24 14:26
 */

var (
	identity  string
	entClient *ent.Client
)

func init() {
	var err error

	identity, err = os.Hostname()
	if err != nil {
		panic(err)
	}

	// init ent client
	entClient, err = ent.Open(
		"sqlite3",
		"file:ent.db?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		panic(err)
	}
	if err = entClient.Schema.Create(context.TODO()); err != nil {
		panic(err)
	}
}

func main() {
	logger, _ := zap.NewDevelopment()
	entrylog := zapr.NewLogger(logger)

	// setup manager
	entrylog.Info("setup manager...")
	mgr, err := manager.New(manager.Options{
		Logger: entrylog,
		LeaderElectionResourceLockInterface: sqllock.NewSqlLock(
			"example-controller",
			entClient,
			resourcelock.ResourceLockConfig{
				Identity:      identity,
				EventRecorder: entrylog,
			},
		),
	})
	if err != nil {
		entrylog.Error(err, "setup manager error")
		os.Exit(1)
	}

	// setup controller
	entrylog.Info("setup controller...")
	err = NewController(mgr)
	if err != nil {
		entrylog.Error(err, "setup controller error")
		os.Exit(1)
	}

	// starting
	entrylog.Info("starting manager...")
	if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
		return
	}
}

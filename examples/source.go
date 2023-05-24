package main

import (
	"context"
	"github.com/yangsijie666/controller-manager/pkg/event"
	"github.com/yangsijie666/controller-manager/pkg/handler"
	"github.com/yangsijie666/controller-manager/pkg/predicate"
	"github.com/yangsijie666/controller-manager/pkg/source"
	"k8s.io/client-go/util/workqueue"
	"math/rand"
	"time"
)

/**
* @Author: yangsijie666
* @Date: 2023/5/24 15:08
 */

var _ source.Source[int] = &Source{}

type Source struct {
}

func (s *Source) Start(ctx context.Context, evtHdr handler.EventHandler[int], q workqueue.RateLimitingInterface, _ ...predicate.Predicate[int]) error {
	go func(ctx context.Context) {
		t := time.NewTicker(time.Second * 5)
		defer t.Stop()

		for {
			s.fetchPendingObjects(ctx, evtHdr, q)

			select {
			case <-ctx.Done():
				return
			case <-t.C:
				continue
			}
		}
	}(ctx)
	return nil
}

func (s *Source) fetchPendingObjects(ctx context.Context, hdr handler.EventHandler[int], q workqueue.RateLimitingInterface) {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	rand.Seed(time.Now().UnixNano())
	hdr.Generic(ctx, event.GenericEvent[int]{
		Object: rand.Intn(100),
	}, q)
}

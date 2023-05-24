package sqllock

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	resourcelock2 "github.com/yangsijie666/controller-manager/pkg/resourcelock"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent"
	"github.com/yangsijie666/controller-manager/pkg/resourcelock/sqllock/ent/leaderelect"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"math/big"
	"time"
)

/**
* @Author: yangsijie666
* @Date: 2023/1/31 09:54
 */

type sqlLock struct {
	Name       string
	Client     *ent.Client
	LockConfig resourcelock2.ResourceLockConfig
	lock       *ent.LeaderElect
}

func NewSqlLock(name string, client *ent.Client, rlc resourcelock2.ResourceLockConfig) resourcelock.Interface {
	return &sqlLock{
		Name:       name,
		Client:     client,
		LockConfig: rlc,
	}
}

func (sl *sqlLock) Get(ctx context.Context) (*resourcelock.LeaderElectionRecord, []byte, error) {
	lock, err := sl.Client.LeaderElect.
		Query().
		Where(leaderelect.ID(sl.Name)).
		Only(ctx)
	if err != nil {
		// trans to k8s not found error
		if ent.IsNotFound(err) {
			return nil, nil, errors.NewNotFound(schema.GroupResource{
				Group:    "",
				Resource: "SqlLock",
			}, sl.Name)
		}
		return nil, nil, err
	}
	sl.lock = lock
	record := ToLeaderElectionRecord(sl.lock)
	recordByte, err := json.Marshal(*record)
	if err != nil {
		return nil, nil, err
	}
	return record, recordByte, nil
}

func (sl *sqlLock) Create(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	do := ToSqlLockDo(&ler)
	var err error
	sl.lock, err = sl.Client.LeaderElect.Create().
		SetID(sl.Name).
		SetResourceVersion("0").
		SetLeaseDurationSeconds(do.LeaseDurationSeconds).
		SetLeaderTransitions(do.LeaderTransitions).
		SetAcquireTime(do.AcquireTime).
		SetRenewTime(do.RenewTime).
		SetHolderIdentity(do.HolderIdentity).
		Save(ctx)
	return err
}

func (sl *sqlLock) Update(ctx context.Context, ler resourcelock.LeaderElectionRecord) error {
	if sl.lock == nil {
		return fmt.Errorf("sqllock not initialized, call get or create first")
	}
	do := ToSqlLockDo(&ler)
	do.ID = sl.Name
	originResourceVersion := sl.lock.ResourceVersion
	if rv, err := SqlResourceVersionPlusOne(sl.lock.ResourceVersion); err != nil {
		return err
	} else {
		do.ResourceVersion = rv
	}

	updateRows, err := sl.Client.LeaderElect.Update().
		Where(leaderelect.ID(do.ID), leaderelect.ResourceVersion(originResourceVersion)).
		SetResourceVersion(do.ResourceVersion).
		SetLeaseDurationSeconds(do.LeaseDurationSeconds).
		SetLeaderTransitions(do.LeaderTransitions).
		SetAcquireTime(do.AcquireTime).
		SetRenewTime(do.RenewTime).
		SetHolderIdentity(do.HolderIdentity).
		Save(ctx)
	if err != nil {
		return err
	}
	if updateRows != 1 {
		return errors.NewConflict(schema.GroupResource{
			Group:    "",
			Resource: "SqlLock",
		}, sl.Name, fmt.Errorf("ResourceVersion too old"))
	}

	sl.lock = do
	return nil
}

func (sl *sqlLock) RecordEvent(s string) {
	if sl.LockConfig.EventRecorder.GetSink() == nil {
		return
	}
	keyVals := []interface{}{
		"lock", sl.Name,
		"identity", sl.LockConfig.Identity,
		"module", "leaderelection",
	}
	sl.LockConfig.EventRecorder.Info(s, keyVals...)
}

func (sl *sqlLock) Identity() string {
	return sl.LockConfig.Identity
}

func (sl *sqlLock) Describe() string {
	return sl.Name
}

func ToLeaderElectionRecord(do *ent.LeaderElect) *resourcelock.LeaderElectionRecord {
	var r resourcelock.LeaderElectionRecord
	if do.HolderIdentity != "" {
		r.HolderIdentity = do.HolderIdentity
	}
	if do.LeaseDurationSeconds >= 0 {
		r.LeaseDurationSeconds = do.LeaseDurationSeconds
	}
	if do.LeaderTransitions >= 0 {
		r.LeaderTransitions = do.LeaderTransitions
	}
	if do.AcquireTime > 0 {
		r.AcquireTime = metav1.NewTime(time.UnixMicro(do.AcquireTime))
	}
	if do.RenewTime > 0 {
		r.RenewTime = metav1.NewTime(time.UnixMicro(do.RenewTime))
	}
	return &r
}

func ToSqlLockDo(ler *resourcelock.LeaderElectionRecord) *ent.LeaderElect {
	return &ent.LeaderElect{
		LeaseDurationSeconds: ler.LeaseDurationSeconds,
		AcquireTime:          ler.AcquireTime.UnixMicro(),
		RenewTime:            ler.RenewTime.UnixMicro(),
		LeaderTransitions:    ler.LeaderTransitions,
		HolderIdentity:       ler.HolderIdentity,
	}
}

func SqlResourceVersionPlusOne(rv string) (string, error) {
	n := new(big.Int)
	n, ok := n.SetString(rv, 10)
	if !ok {
		return "", fmt.Errorf("resourceVersion plus 1 error: setstring error")
	}
	n.Add(n, big.NewInt(1))
	return n.String(), nil
}

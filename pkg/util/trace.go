package util

import (
	"context"
	"k8s.io/apimachinery/pkg/types"
)

/**
* @Author: yangsijie666
* @Date: 2023/5/24 15:24
 */

// ReconcileIDKey is a context.Context Value key. Its associated value should
// be a types.UID.
type ReconcileIDKey struct{}

func AddReconcileID(ctx context.Context, reconcileID types.UID) context.Context {
	return context.WithValue(ctx, ReconcileIDKey{}, reconcileID)
}

func ExtractReconcileID(ctx context.Context) types.UID {
	if ctx == nil {
		return ""
	}
	if reconcileID, ok := ctx.Value(ReconcileIDKey{}).(types.UID); ok {
		return reconcileID
	}
	return ""
}

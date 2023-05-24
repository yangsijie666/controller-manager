package resourcelock

import (
	"github.com/go-logr/logr"
)

/**
* @Author: yangsijie666
* @Date: 2023/1/31 10:03
 */

type ResourceLockConfig struct {
	Identity      string
	EventRecorder logr.Logger
}

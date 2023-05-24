package predicate

import "github.com/yangsijie666/controller-manager/pkg/event"

/**
* @Author: yangsijie666
* @Date: 2023/2/2 14:02
 */

type Predicate[T interface{}] interface {
	Generic(event event.GenericEvent[T]) bool
}

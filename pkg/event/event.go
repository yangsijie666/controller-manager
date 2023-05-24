package event

/**
* @Author: yangsijie666
* @Date: 2023/2/2 14:00
 */

type GenericEvent[T interface{}] struct {
	Object T
}

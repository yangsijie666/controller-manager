package ratelimiter

/**
* @Author: yangsijie666
* @Date: 2023/2/2 14:31
 */

import "time"

// RateLimiter is an identical interface of client-go workqueue RateLimiter.
type RateLimiter interface {
	// When gets an item and gets to decide how long that item should wait
	When(item interface{}) time.Duration
	// Forget indicates that an item is finished being retried.  Doesn't matter whether its for perm failing
	// or for success, we'll stop tracking it
	Forget(item interface{})
	// NumRequeues returns back how many failures the item has had
	NumRequeues(item interface{}) int
}

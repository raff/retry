package retry

import "time"

const MaxRetries = 8

/*
 * RetrySleep can be used to implement retries with exponential backoff
 *
 * It returns `false` when the number of retries exceeds the maximum
 *
 * usage:
 *   for r := 0; RetrySleep(r, MaxRetries); r++ {
 *
 *      ok := doSomething()
 *      if ok {
 *        break
 *      }
 *   }
 */
func RetrySleep(r, max int) bool {
	if r >= max {
		return false
	}

	if r > 0 {
		p := 1 << uint(r)

		sleep_time := /* random.uniform(0.1, 0.13) * */ p * 100
		time.Sleep(time.Duration(sleep_time) * time.Millisecond)
	}

	return true
}

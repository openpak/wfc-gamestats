package gamestats

import (
	"sync/atomic"
	"time"

	"openpak/wfc-common/logging"
)

// refusedUnlinked counts requests refused because the console is not linked (link-first,
// prds repos/wfc-link-first-design.md); summarised once a minute.
var refusedUnlinked atomic.Int64

func init() {
	go func() {
		for range time.Tick(time.Minute) {
			if n := refusedUnlinked.Swap(0); n > 0 {
				logging.Notice("GSTATS", "link gate: refused", n, "stats logins of unlinked consoles in the last minute")
			}
		}
	}()
}

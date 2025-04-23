package syncprogram_test

import (
	"sync"
	"testing"

	syncprogram "github.com/anicse37/TDD_Go/Sync"
)

func TestSync(t *testing.T) {
	t.Run("Single Increments", func(t *testing.T) {

		counter := syncprogram.Counter{}
		counter.Increase()
		counter.Increase()
		counter.Increase()
		got := counter.Print()
		if got != 3 {
			t.Errorf("Got %v || Want 3 ", got)
		}
	})
	t.Run("Using Loops", func(t *testing.T) {
		counter := syncprogram.Counter{}
		counterRange := 100
		var wg sync.WaitGroup
		wg.Add(counterRange)
		for i := 0; i < counterRange; i++ {
			go func() {
				counter.Increase()
				wg.Done()
			}()
		}
		wg.Wait()
		if counter.Print() != counterRange {
			t.Errorf("Got %v || Want %v ", counter.Print(), counterRange)
		}

	})
}

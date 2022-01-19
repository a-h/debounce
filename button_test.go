package debounce

import (
	"testing"
	"time"
)

func TestButton(t *testing.T) {
	tests := []struct {
		name           string
		normallyClosed bool
		test           func(s *Switch)
		expectedCount  int
	}{
		{
			name:           "for a normally open switch, onClick is triggered when the button closes",
			normallyClosed: false,
			test: func(sw *Switch) {
				sw.SetState(true)
			},
			expectedCount: 1,
		},
		{
			name:           "for a normally closed switch, onClick is triggered when the button opens",
			normallyClosed: true,
			test: func(sw *Switch) {
				sw.SetState(false)
			},
			expectedCount: 1,
		},
		{
			name:           "multiple changes within the debounce period only trigger a single change",
			normallyClosed: false,
			test: func(sw *Switch) {
				sw.SetState(false)
				sw.SetState(true)
				sw.SetState(true)
				sw.SetState(false)
				sw.SetState(true)
			},
			expectedCount: 1,
		},
		{
			name:           "multiple changes after the debounce period each trigger a single change",
			normallyClosed: false,
			test: func(sw *Switch) {
				sw.SetState(true)
				sw.SetState(false)
				sw.SetState(true)
				time.Sleep(time.Millisecond * 2)
				sw.bounceTime = time.Millisecond
				sw.SetState(true)
				sw.SetState(false)
				sw.SetState(true)
				time.Sleep(time.Millisecond * 2)
				sw.SetState(true)
			},
			expectedCount: 3,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var count int
			counter := func() {
				count++
			}
			sw := Button(counter, test.normallyClosed)
			test.test(sw)
			if count != test.expectedCount {
				t.Errorf("expected count %d, got %d", test.expectedCount, count)
			}
		})
	}
}

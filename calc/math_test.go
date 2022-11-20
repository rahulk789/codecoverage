//unit tests for math.go
package math

import "testing"

func TestAdd( t *testing.T) {
    ans := Add(4,6)
    want := 10
    if ans != want {
        t.Errorf("got %q, wanted %q", ans , want)
    }
}

package humanized

import (
	"fmt"
	"testing"
	"time"
)

func TestSecond(t *testing.T) {
	tm := time.Unix(1670480999, 0)
	fmt.Printf("%s.\n", Time(tm))
}

func TestMinute(t *testing.T) {
	tm := time.Unix(1350480999, 0)
	fmt.Printf("%s.\n", Time(tm))
}

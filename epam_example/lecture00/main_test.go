package lecture00

import (
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := emoji.Sprint("Hello :world_map:!")

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

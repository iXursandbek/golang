package lecture00

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {

	fmt.Println(Hello())
}

func Hello() string {
	return emoji.Sprint("Hello :world_map:!")
}

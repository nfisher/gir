package main

import (
	"fmt"

	"github.com/nfisher/gir"
)

const selection = `{
  people {
    name,
    starships {
      name,
      length
    }
  }
}`

const query = `query HeroNameAndFriends($episode: Episode = "JEDI") {
  hero(episode: $episode) {
    name
    friends {
      name
    }
  }
}`

func main() {
	gir.ParseQuery([]byte(`hello 123 245 abc "ack"`))
	fmt.Println("= = = = = = = = = = = = = = = = = = = =")
	gir.ParseQuery([]byte(`[1.3 4.5]`))
}

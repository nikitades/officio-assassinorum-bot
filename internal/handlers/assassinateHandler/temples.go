package assassinateHandler

import (
	"fmt"
	"log"
	"math/rand"
)

type temple struct {
	name   string
	image  string
	chance float32
}

func (t *temple) doYouConfirmTheTarget() bool {
	r := rand.Float32()

	log.Println(fmt.Sprintf("Temple %v has the rand %v against the chance %v", t.name, r, 1-t.chance))

	return r > 1-t.chance
}

var temples = []temple{
	temple{
		name:   "Vindicare",
		image:  "https://static.wikia.nocookie.net/warhammer40k/images/1/14/Vindicare_Temple_Icon.jpg",
		chance: 0.004,
	},
	temple{
		name:   "Callidus",
		image:  "https://static.wikia.nocookie.net/warhammer40k/images/1/14/Callidus_Temple_Icon.jpg",
		chance: 0.003,
	},
	temple{
		name:   "Eversor",
		image:  "https://static.wikia.nocookie.net/warhammer40k/images/0/01/Eversor_Temple_Icon.jpg",
		chance: 0.011,
	},
	temple{
		name:   "Culexus",
		image:  "https://static.wikia.nocookie.net/warhammer40k/images/3/31/Cullexus_Temple_Icon.jpg",
		chance: 0.002,
	},
	temple{
		name:   "Venenum",
		image:  "https://static.wikia.nocookie.net/warhammerfb/images/8/87/Poison.jpg",
		chance: 0.002,
	},
	temple{
		name:   "Vanus",
		image:  "https://yt3.ggpht.com/ytc/AKedOLQyXRhiC3MHRO2q_o3TT151tQnfNGSESfGxf8w8_w=s900-c-k-c0x00ffffff-no-rj",
		chance: 0.003,
	},
	temple{
		name:   "Maerorus",
		image:  "https://i.pinimg.com/736x/cf/2f/c3/cf2fc30aec60c58441e0f45061e1c457.jpg",
		chance: 0.0001,
	},
}

func getTemple() *temple {
	currentTemples := temples
	rand.Shuffle(len(currentTemples), func(a, b int) { currentTemples[a], currentTemples[b] = currentTemples[b], currentTemples[a] })

	for _, temple := range currentTemples {
		if temple.doYouConfirmTheTarget() {
			return &temple
		}
	}

	return nil
}

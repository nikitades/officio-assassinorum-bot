package temples

import (
	"log"
	"math/rand"
)

//Temple is a building of the sacred purpose
type Temple struct {
	Name   string
	Image  string
	Chance float32
}

func (t *Temple) doYouConfirmTheTarget() bool {
	r := rand.Float32()

	log.Printf("Temple %v has the rand %v against the chance %v", t.Name, r, 1-t.Chance)

	return r > 1-t.Chance
}

func (t *Temple) malcadorAsksToConfirmTheFollowingTarget() bool {
	r := rand.Float32()

	chance := t.Chance * 10

	log.Printf("Temple %v has the rand %v against the special chance %v", t.Name, r, 1-chance)

	return r > 1-chance
}

var temples = []Temple{
	Temple{
		Name:   "Vindicare",
		Image:  "https://static.wikia.nocookie.net/warhammer40k/images/1/14/Vindicare_Temple_Icon.jpg",
		Chance: 0.004,
	},
	Temple{
		Name:   "Callidus",
		Image:  "https://static.wikia.nocookie.net/warhammer40k/images/1/14/Callidus_Temple_Icon.jpg",
		Chance: 0.003,
	},
	Temple{
		Name:   "Eversor",
		Image:  "https://static.wikia.nocookie.net/warhammer40k/images/0/01/Eversor_Temple_Icon.jpg",
		Chance: 0.006,
	},
	Temple{
		Name:   "Culexus",
		Image:  "https://static.wikia.nocookie.net/warhammer40k/images/3/31/Cullexus_Temple_Icon.jpg",
		Chance: 0.002,
	},
	Temple{
		Name:   "Venenum",
		Image:  "https://static.wikia.nocookie.net/warhammerfb/images/8/87/Poison.jpg",
		Chance: 0.003,
	},
	Temple{
		Name:   "Vanus",
		Image:  "https://yt3.ggpht.com/ytc/AKedOLQyXRhiC3MHRO2q_o3TT151tQnfNGSESfGxf8w8_w=s900-c-k-c0x00ffffff-no-rj",
		Chance: 0.004,
	},
	Temple{
		Name:   "Maerorus",
		Image:  "https://i.pinimg.com/736x/cf/2f/c3/cf2fc30aec60c58441e0f45061e1c457.jpg",
		Chance: 0.0001,
	},
}

//GetTemple is called upon every incoming non-reactible message and has the standard chance of the successful outcome
func GetTemple() *Temple {
	currentTemples := temples
	rand.Shuffle(len(currentTemples), func(a, b int) { currentTemples[a], currentTemples[b] = currentTemples[b], currentTemples[a] })

	for _, Temple := range currentTemples {
		if Temple.doYouConfirmTheTarget() {
			return &Temple
		}
	}

	return nil
}

//GetSpecialTemple is supposed to be called upon a command request and has an increased chance of the successful outcome
func GetSpecialTemple() *Temple {
	currentTemples := temples
	rand.Shuffle(len(currentTemples), func(a, b int) { currentTemples[a], currentTemples[b] = currentTemples[b], currentTemples[a] })

	for _, Temple := range currentTemples {
		if Temple.malcadorAsksToConfirmTheFollowingTarget() {
			return &Temple
		}
	}

	return nil
}

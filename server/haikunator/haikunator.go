package haikunator

//  source : https://github.com/yelinaung/go-haikunator

import (
	"fmt"
	"math/rand"
)

var (
	ADJECTIVES = []string{"philip-johnson", "luis-barragan", "james-stirling", "kevin-roche", "im-pei",
		"richard-meier", "hans-hollein", "gottfried-bohm", "kenzo-tange", "oscar-niemeyer", "gordon-bunshaft", "frank-gehry", "aldo-rossi",
		"robert-venturi", "alvaro-siza", "fumihiko-maki", "chritian-de-portzamparc", "tadao-ando", "rafael-moneo", "sverre-fehn",
		"renzo-piano", "norman-foster", "rem-koolhas", "jaques-herzog", "pierre-de-meuron", "glenn-murcutt", "john-utzon", "zaha-hadid", "thom-mayne",
		"paulo-mendes", "richard-rogers", "jean-nouvel", "peter-zumthor", "kazuyo-sejima", "ryue-nishizawa", "eduardo-souto", "wang-shu", "toyo-ito",
		"shigeru-ban", "frei-otto", "alejandro-aravena", "rafael-aranda", "carme-pigem", "ramon-vilalta", "bv-doshi",
		"arata-isozaki", "yvonne-farell", "shelly-mcnamara"}
	NOUNS = []string{"glass-house", "library", "institute", "gallery", "cathedral", "hall", "museum", "pavilion",
		"palace", "gymnasium", "embassy", "church", "airport", "bridge", "tower", "inn",
		"center", "chapel", "assembly", "stadium", "university", "hospital", "school", "art",
		"headquarters", "world-center", "park", "shelter", "landscape", "station", "concert-hall", "skyscraper",
		"house", "hotel", "bath", "island", "bank", "bowl", "courthouse",
		"arena", "tunnel", "temple", "district", "gardens", "community", "plantation", "avenue", "square", "street", "monument", "memorial", "statue", "installation"}
)

type Name interface {
	Haikunate() string
}

type RandomName struct {
	r *rand.Rand
}

func (r RandomName) Haikunate() string {
	return fmt.Sprintf("%v-%v", ADJECTIVES[r.r.Intn(len(ADJECTIVES))], NOUNS[r.r.Intn(len(NOUNS))])
}

func New(seed int64) Name {
	name := RandomName{rand.New(rand.New(rand.NewSource(99)))}
	name.r.Seed(seed)
	return name
}

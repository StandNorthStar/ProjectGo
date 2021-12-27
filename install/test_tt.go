package main

const (
	tye = iota
)

type cc struct {
	c1 map[string]string
}

type config struct {
	a1 cc
	a2 map[string]string
}

func MakeConfig() []config {
	return []config{
		{
			a1: cc{
				c1: map[string]string{
					"c11": "c11-haha",
				},
			},
			a2: map[string]string{
				"a22": "a22-heihie",
			},
		},
		{
			a1: cc{
				c1: map[string]string{
					"c12": "c12-haha",
				},
			},
			a2: map[string]string{
				"a23": "a23-heihie",
			},
		},
	}
}


func main() {

	aa2 := make(map[string]string)
	aa2["aa2-2"] = "aa2-test"

	ccc1 := []config


}

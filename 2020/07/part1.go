package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var in io.Reader = os.Stdin

	if len(os.Args) < 2 {
		log.Fatal("too few arguments")
	}
	//mine := os.Args[1]
	if len(os.Args) > 2 {
		f, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		in = f
	}

	g := makeGraph(in)
	j, _ := json.MarshalIndent(g, "", " ")

	fmt.Println(string(j))

	//containers := containingBags(mine)
	fmt.Printf("graph: %++v\n", g)
}

func containingBags(m map[string][]string, mine string) []string {
	cm, ok := m[mine]
	if !ok {
		return []string{}
	}

	var cc []string
	for _, c := range cm {
		cc = append(cc, c)
		cc = append(cc, containingBags(m, c)...)
	}
	mm := make(map[string]struct{})
	var containers []string
	for _, c := range cc {
		if _, seen := mm[c]; seen {
			continue
		}
		mm[c] = struct{}{}
		containers = append(containers, c)
	}

	return containers
}

type content struct {
	count int
	color string
}

func makeGraph(in io.Reader) (map[string][]content, map[string][]string) {
	scanner := bufio.NewScanner(in)
	phrase := " bags contain "
	contains := make(map[string][]content)
	for scanner.Scan() {
		s := scanner.Text()

		// xxx contains ....
		pos := strings.Index(s, phrase)
		if pos == -1 {
			log.Fatal(`"` + phrase + `" not found in ` + s)
		}

		container := strings.TrimSpace(s[:pos])
		if _, ok := contains[container]; ok {
			log.Fatal("duplicate container " + container)
		}

		// n yyy bag(s)...
		contains[container] = []content{}
		for _, c := range strings.Split(s[pos+len(phrase):], ",") {
			i := strings.Index(c, " bag")
			c = strings.TrimSpace(c[:i])
			i = strings.IndexByte(c, ' ')
			n, err := strconv.Atoi(c[:i])
			if err != nil {
				log.Fatal(err)
			}
			c = c[n+1:]
			contains[container] = append(contains[container], content{color: c, count: n})
		}
	}
	return contains
}

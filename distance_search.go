package main

import "fmt"

func main() {
	graph := buildGraph()
	var alreadyChecked []string
	var searchQueue []string

	searchQueue = graph["you"]
	var mangoDealer string
	for i := 0; i < len(searchQueue); i++ {
		if isChecked(searchQueue[i], alreadyChecked) {
			continue
		}
		if isSeller(searchQueue[i]) {
			mangoDealer = searchQueue[i]
			break
		} else {
			searchQueue = append(searchQueue, graph[searchQueue[i]]...)
			alreadyChecked = append(alreadyChecked, searchQueue[i])
		}

	}

	if mangoDealer == "" {
		fmt.Println("There's no mango dealer")
	} else {
		fmt.Printf("%s is the mango dealer", mangoDealer)
	}
}

func isChecked(name string, checked []string) bool {
	for _, val := range checked {
		if name == val {
			return true
		}
	}
	return false
}

func isSeller(name string) bool {
	return name == "thom"
}

func buildGraph() map[string][]string {
	graph := make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	return graph
}

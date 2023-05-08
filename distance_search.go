package main

import "fmt"

func main() {
	graph := buildGraph()
	var alreadyChecked []string
	var searchQueue []string

	searchQueue = graph["you"]
	for i := 0; i < len(searchQueue); i++ {
		if isChecked(searchQueue[i], alreadyChecked) {
			continue
		} else {
			if isSeller(searchQueue[i]) {
				fmt.Printf(" %s is a Mango Dealer \n", searchQueue[i])
				break
			} else {
				searchQueue = append(searchQueue, graph[searchQueue[i]]...)
				alreadyChecked = append(alreadyChecked, searchQueue[i])
			}
		}
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

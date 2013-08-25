/*

The following undirected network consists of seven vertices and twelve edges
with a total weight of 243.

The same network can be represented by the matrix below.

     A    B    C    D    E    F    G
A    -    16   12   21   -    -    -
B    16   -    -    17   20   -    -
C    12   -    -    28   -    31   -
D    21   17   28   -    18   19   23
E    -    20   -    18   -    -    11
F    -    -    31   19   -    -    27
G    -    -    -    23   11   27   -

However, it is possible to optimise the network by removing some edges and
still ensure that all points on the network remain connected. The network which
achieves the maximum saving is shown below. It has a weight of 93, representing
a saving of 243 - 93 = 150 from the original network.

Using network.txt (right click and 'Save Link/Target As...'), a 6K text file
containing a network with forty vertices, and given in matrix form, find the
maximum saving which can be achieved by removing redundant edges whilst
ensuring that the network remains connected.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type linkpair struct {
	from_node, to_node uint
}

type link struct {
	from_node, to_node uint
	cost               uint64
}

type network struct {
	links []link
}

func (n *network) addlink(from_node uint, to_node uint, cost uint64) {
	if from_node < to_node {
		from_node, to_node = to_node, from_node
	}
	n.links = append(n.links, link{from_node, to_node, cost})
}

func (n *network) nodecount() uint {
	nc := uint(0)
	for idx := 0; idx < len(n.links); idx += 1 {
		if n.links[idx].to_node >= nc {
			nc = n.links[idx].to_node + 1
		} else if n.links[idx].from_node >= nc {
			nc = n.links[idx].from_node + 1
		}
	}
	return nc
}

func (n *network) prettyprint() {
	link_map := make(map[linkpair]uint64)

	for idx := 0; idx < len(n.links); idx += 1 {
		link := n.links[idx]

		link_map[linkpair{link.from_node, link.to_node}] = link.cost
	}

	for row_idx := uint(0); row_idx < n.nodecount(); row_idx += 1 {
		for col_idx := uint(0); col_idx < n.nodecount(); col_idx += 1 {
			if col_idx > 0 {
				fmt.Printf(",")
			}

			cost, has_key := link_map[linkpair{row_idx, col_idx}]

			if has_key {
				fmt.Printf("%v", cost)
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println("")
	}
}

// Paperwork to implement sort interface
func (n *network) Len() int {
	return len(n.links)
}

func (n *network) Swap(i, j int) {
	n.links[i], n.links[j] = n.links[j], n.links[i]
}

func (n *network) Less(i, j int) bool {
	return n.links[i].cost > n.links[j].cost
}

func (n *network) sortbycost() {
	sort.Sort(n)
}

// Method to make actual fixed network
func (n *network) minimumspanningtree() network {
	var newnetwork network

	n.sortbycost()

	link_in_count := make([]int, n.nodecount())
	link_out_count := make([]int, n.nodecount())

	for i := 0; i < len(n.links); i += 1 {
		link := n.links[i]
		link_out_count[link.from_node] += 1
		link_in_count[link.to_node] += 1
	}

	for i := 0; i < len(n.links); i += 1 {
		link := n.links[i]
		if link_out_count[link.from_node] > 1 && link_in_count[link.to_node] > 1 {
			link_out_count[link.from_node] -= 1
			link_in_count[link.to_node] -= 1
		} else {
			newnetwork.addlink(link.from_node, link.to_node, link.cost)
		}
	}

	return newnetwork
}

func loadnetwork(filename string) (network, error) {
	var newnetwork network

	fmt.Println("Loading file", filename)

	file_handle, err := os.Open(filename)

	if err != nil {
		return newnetwork, err
	}

	file_reader := bufio.NewReader(file_handle)

	line_index := uint(0)

	for {
		line, _, err := file_reader.ReadLine()
		if err != nil {
			break
		}

		link_items := strings.Split(string(line), ",")

		for col_index := uint(0); col_index < uint(len(link_items)); col_index += 1 {
			cost, err := strconv.ParseUint(link_items[col_index], 10, 64)
			if err == nil {
				newnetwork.addlink(line_index, col_index, cost)
			}
		}

		line_index += 1
	}

	return newnetwork, err
}

func main() {
	loadednetwork, err := loadnetwork("network.txt")

	if err != nil {
		fmt.Printf("Error loading network file\n")
		return
	}

	newnetwork := loadednetwork.minimumspanningtree()

	fmt.Println("-- Old Network --")
	loadednetwork.prettyprint()
	fmt.Println("-- New Network --")
	newnetwork.prettyprint()
}

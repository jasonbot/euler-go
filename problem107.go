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
    "strconv"
	"strings"
)

type link struct {
	from_node, to_node uint
	cost               uint64
}

type network struct {
	links []link
}

func (n *network) minimumspanningtree() network {
	var newnetwork network
	return newnetwork
}

func (n *network) addlink(from_node uint, to_node uint, cost uint64) {
    n.links = append(n.links, link{from_node, to_node, cost})
}

func loadnetwork(filename string) (network, error) {
	var newnetwork network

	fmt.Printf("Loading file %s", filename)

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
		fmt.Printf("Line: %s\n", line)

		links := strings.Split(string(line), ",")

        for col_index := uint(0); col_index < uint(len(links)); col_index += 1 {
            cost, err := strconv.ParseUint(links[col_index], 10, 64);
            if (err != nil) {
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

	loadednetwork.minimumspanningtree()
}

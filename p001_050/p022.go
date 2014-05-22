package main

import (
	"bufio"
	"fmt"
	eu "github.com/wresch/euler_go"
	"os"
	"sort"
	"strings"
)

var desc = `

Using p022.data.txt, 46K text file containing over five-thousand first
names, begin by sorting it into alphabetical order. Then working out
the alphabetical value for each name, multiply this value by its
alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN,
which is worth $3 + 15 + 12 + 9 + 14 = 53$, is the 938th name in the
list. So, COLIN would obtain a score of $938 \cdot 53 = 49714$.

\emph{What is the total of all the name scores in the file?}

\subsection{Approach}

Not much to it - just figuring out how to read and sort in go.
Input data is comma separated and quoted. There are no newlines.
`

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inf, err := os.Open("p022.data.txt")
	check(err)
	defer inf.Close()
	scanner := bufio.NewScanner(inf)
	names := make([]string, 0, 2000)
	scanner.Split(eu.ScanComma)
	for scanner.Scan() {
		names = append(names, strings.Trim(scanner.Text(), "\""))
	}
	sort.Strings(names)
	var total int64 = 0
	for i, n := range names {
		ns := 0
		for _, r := range n {
			ns += int(r-'A') + 1
		}
		//fmt.Printf("%s -> %5d * %5d = %6d\n", n, i + 1, ns, ns * (i + 1))
		total += int64(ns) * int64(i+1)
	}
	eu.Output(22, desc,
		fmt.Sprintf("The sum of all name scoes is %d\n", total))
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type NameCount struct {
	Name  string
	Count int
}

func CountNames(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return CountNamesFromReader(file)
}

func CountNamesFromReader(reader io.Reader) (map[string]int, error) {
	counts := make(map[string]int)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())
		if name == "" {
			continue
		}

		counts[name]++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return counts, nil
}

func MakeNameList(counts map[string]int, sortByFrequency bool) []NameCount {
	list := make([]NameCount, 0, len(counts))

	for name, count := range counts {
		list = append(list, NameCount{
			Name:  name,
			Count: count,
		})
	}

	if sortByFrequency {
		sort.Slice(list, func(i, j int) bool {
			if list[i].Count == list[j].Count {
				return list[i].Name < list[j].Name
			}
			return list[i].Count > list[j].Count
		})
		return list
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})

	return list
}

func PrintResult(counts map[string]int, sortByFrequency bool, writer io.Writer) {
	list := MakeNameList(counts, sortByFrequency)

	for _, item := range list {
		fmt.Fprintf(writer, "%s:%d\n", item.Name, item.Count)
	}
}

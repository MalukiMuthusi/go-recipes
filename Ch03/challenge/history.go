package main

import (
	"bufio"
	"io"
	"regexp"
)

// cmdFreq returns the frequency of "go" subcommand usage in ZSH history
func cmdFreq(r io.Reader) (map[string]int, error) {
	// read the file

	commandRegex := `:\s\d+:0;go\s(\w+)\s`
	commandRe := regexp.MustCompile(commandRegex)
	matched := make(map[string]int)

	s := bufio.NewScanner(r)
	for s.Scan() {
		matches := commandRe.FindStringSubmatch(s.Text())
		if matches != nil {
			matched[matches[1]]++
		}
	}

	return matched, nil
}

func main() {
	// freqs, err := cmdFreq("zsh_history")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for cmd, count := range freqs {
	// 	fmt.Printf("%s -> %d\n", cmd, count)
	// }
}

package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/agnivade/levenshtein"
)

var (
	w = '⬜'
	y = '🟨'
	g = '🟩'
)

func main() {
	var answer, dark bool
	flag.BoolVar(&answer, "answer", false, "show me the damn word")
	flag.BoolVar(&dark, "dark", false, "use dark mode")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fail(errors.New("no command to not find"))
	}
	cmd := flag.Args()[0]
	if dark {
		w = '⬛'
	}

	binChan, err := getCommands(len(cmd))
	if err != nil {
		fail(fmt.Errorf("command %q not found and: %w", cmd, err))
	}

	type result struct {
		command  string
		distance int
	}

	levWg := sync.WaitGroup{}
	levChan := make(chan result)
	for bin := range binChan {
		levWg.Add(1)
		go func(bin string) {
			defer levWg.Done()
			levChan <- result{
				command:  bin,
				distance: levenshtein.ComputeDistance(cmd, bin),
			}
		}(bin)
	}
	go func() {
		levWg.Wait()
		close(levChan)
	}()

	var best *result
	for r := range levChan {
		if best != nil && best.distance <= r.distance {
			continue
		}
		best = &result{
			command:  r.command,
			distance: r.distance,
		}

	}

	if nil == best {
		fail(fmt.Errorf("command %q not found and: didn't find any similar commands in PATH", cmd))
	}

	if answer {
		fmt.Printf("it was %q\n", best.command)
		os.Exit(0)
	}

	fmt.Println("Did you mean:", wordle(cmd, best.command))
	os.Exit(127)
}

// getCommands returns the name of every file in PATH executable by other which
// is length characters long.
func getCommands(length int) (<-chan string, error) {
	paths := strings.Split(os.Getenv("PATH"), ":")
	n := len(paths)
	if n < 1 {
		return nil, fmt.Errorf("no search PATH")
	}

	binChan := make(chan string)
	binWg := sync.WaitGroup{}
	binWg.Add(n)
	go func() {
		binWg.Wait()
		close(binChan)
	}()
	for _, path := range paths {
		go func(path string) {
			defer binWg.Done()
			_ = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return nil
				}
				if d.IsDir() {
					return nil
				}
				i := 0
				for i = range d.Name() {
				}
				if i+1 != length {
					return nil
				}
				stat, err := d.Info()
				if err != nil {
					return nil
				}
				if stat.Mode()&0001 == 0 {
					return nil
				}
				binChan <- d.Name()
				return nil
			})
		}(path)
	}
	return binChan, nil
}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}

// wordle wordlises the guess and target, it probably panics if they aren't the
// same length
func wordle(guess, target string) string {
	rGuess := runify(guess)
	used := make([]bool, len(rGuess))
	rTatget := runify(target)
	out := make([]rune, len(rTatget))
	for i := range out {
		out[i] = w
	}

	for i := range rTatget {
		if rTatget[i] == rGuess[i] {
			out[i] = g
			if i < len(rGuess) {
				used[i] = true
			}
		}
	}

	for i := range rTatget {
		if out[i] == g {
			continue
		}
		for j := range rGuess {
			if used[j] {
				continue
			}
			if rTatget[i] == rGuess[j] {
				out[j] = y
				used[j] = true
				break
			}
		}
	}

	var s string
	for _, r := range out {
		s += string(r)
	}
	return s
}

func runify(s string) (runes []rune) {
	runes = make([]rune, 0, len(s))
	for _, r := range s {
		runes = append(runes, r)
	}
	return
}

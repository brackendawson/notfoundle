package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordle(t *testing.T) {
	type inputs struct {
		guess, target string
	}
	for in, out := range map[inputs]string{
		{"this", "this"}:       "🟩🟩🟩🟩",
		{"this", "that"}:       "🟩🟩⬜⬜",
		{"kins", "sink"}:       "🟨🟩🟩🟨",
		{"bracken", "barnyen"}: "🟩🟨🟨⬜⬜🟩🟩",
		{"yellow", "lagoon"}:   "⬜⬜🟨⬜🟩⬜",
		{"probet", "reboot"}:   "⬜🟨🟨🟨🟨🟩",
	} {
		t.Run(in.guess+":"+in.target, func(t *testing.T) {
			assert.Equal(t, out, wordle(in.guess, in.target))
		})
	}
}

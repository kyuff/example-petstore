package app

import "github.com/kyuff/anchor"

func Run(wire anchor.Wire) int {
	return anchor.New(wire, anchor.WithDefaultSlog()).
		Add(
			httpServer(),
		).
		Run()
}

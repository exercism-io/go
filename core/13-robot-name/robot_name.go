// Package robotname includes a solution for the "Robot Name" problem in the Go track on https://exercism.io.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var namespaceLimit int = 26 * 26 * 10 * 10 * 10
var seenNames = map[string]bool{}

// Robot is a robot with a name.
type Robot struct {
	name string
}

// Name generates a random name for a robot.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(seenNames) >= namespaceLimit {
		return "", errors.New("namespace is exhausted")
	}

	r.name = randName()
	for seenNames[r.name] {
		r.name = randName()
	}
	seenNames[r.name] = true
	return r.name, nil
}

// Reset resets the robot to its factory settings, which means that its name gets wiped.
func (r *Robot) Reset() {
	r.name = ""
}

func randName() string {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	num := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}

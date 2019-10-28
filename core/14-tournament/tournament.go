// Package tournament includes a solution for the "Tournament" problem in the Go track on https://exercism.io.
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const tableHeaderFormat = "%-30s | %2s | %2s | %2s | %2s | %2s\n"
const tableRowFormat = "%-30s | %2d | %2d | %2d | %2d | %2d\n"

type team struct {
	name   string
	win    int
	loss   int
	draw   int
	points int
	mp     int
}

// Tally tallies the results of a small football competition.
func Tally(reader io.Reader, writer io.Writer) error {
	teamMap, err := buildTeamMap(reader)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(writer, tableHeaderFormat, "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}
	for _, t := range getSortedTeams(teamMap) {
		_, err = fmt.Fprintf(writer, tableRowFormat, t.name, t.mp, t.win, t.draw, t.loss, t.points)
		if err != nil {
			break
		}
	}
	return err
}

func buildTeamMap(reader io.Reader) (map[string]team, error) {
	teamMap := make(map[string]team)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		parts := strings.Split(trimmed, ";")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		if len(parts) != 3 || parts[0] == "" || parts[1] == "" {
			return nil, fmt.Errorf("bad line: %q; expected format %q", line, "{team};{team};{win/loss/draw}")
		}

		team1 := teamMap[parts[0]]
		team2 := teamMap[parts[1]]
		team1.name = parts[0]
		team2.name = parts[1]
		team1.mp++
		team2.mp++
		switch parts[2] {
		case "win":
			team1.win++
			team1.points += 3
			team2.loss++
		case "loss":
			team1.loss++
			team2.win++
			team2.points += 3
		case "draw":
			team1.draw++
			team1.points++
			team2.draw++
			team2.points++
		default:
			return nil, fmt.Errorf("bad line: %q; expected result format %q", line, "{win/loss/draw}")
		}
		teamMap[parts[0]] = team1
		teamMap[parts[1]] = team2
	}

	return teamMap, nil
}

func getSortedTeams(teamMap map[string]team) []team {
	teams := make([]team, 0, len(teamMap))
	for _, v := range teamMap {
		teams = append(teams, v)
	}

	sort.Slice(teams, func(i, j int) bool {
		a, b := teams[i], teams[j]
		if a.points == b.points {
			return a.name < b.name
		}
		return a.points > b.points
	})
	return teams
}

package leaderboard

import (
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

// LeaderboardEntry represents a single leaderboard record.
type LeaderboardEntry struct {
	Name  string
	Score int
}

// LeaderboardStore defines the interface for leaderboard persistence.
type LeaderboardStore interface {
	Load() ([]LeaderboardEntry, error)
	Save(entries []LeaderboardEntry) error
	TopN(n int) ([]LeaderboardEntry, error)
}

// FileLeaderboardStore implements LeaderboardStore using a file.
type FileLeaderboardStore struct {
	Path string
}

// Load reads leaderboard entries from the file.
func (f *FileLeaderboardStore) Load() ([]LeaderboardEntry, error) {
	data, err := ioutil.ReadFile(f.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	var entries []LeaderboardEntry
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ",", 2)
		if len(parts) != 2 {
			continue
		}
		score, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		entries = append(entries, LeaderboardEntry{Name: parts[0], Score: score})
	}
	return entries, nil
}

// Save writes leaderboard entries to the file.
func (f *FileLeaderboardStore) Save(entries []LeaderboardEntry) error {
	var lines []string
	for _, entry := range entries {
		lines = append(lines, entry.Name+","+strconv.Itoa(entry.Score))
	}
	return ioutil.WriteFile(f.Path, []byte(strings.Join(lines, "\n")), 0644)
}

// TopN returns the top N leaderboard entries, sorted by score descending.
func (f *FileLeaderboardStore) TopN(n int) ([]LeaderboardEntry, error) {
	entries, err := f.Load()
	if err != nil {
		return nil, err
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Score > entries[j].Score
	})
	if len(entries) > n {
		entries = entries[:n]
	}
	return entries, nil
} 
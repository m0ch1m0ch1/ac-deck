package lib

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseTaskPage(t *testing.T) {
	file, err := os.Open("testdata/abc153_a.html")
	if err != nil {
		t.Fatal(err)
	}

	task, err := ParseTaskPage(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, &Task{
		Name: "A - Serval vs Monster",
		Examples: []*Example{
			{
				In:  "10 4",
				Exp: "3",
			},
			{
				In:  "1 10000",
				Exp: "1",
			},
			{
				In:  "10000 1",
				Exp: "10000",
			},
		},
	}, task)

}

func TestParseTasksPage(t *testing.T) {
	file, err := os.Open("testdata/abc153_tasks.html")
	if err != nil {
		t.Fatal(err)
	}

	tasksPaths, err := ParseTasksPage(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, []string{
		"/contests/abc153/tasks/abc153_a",
		"/contests/abc153/tasks/abc153_b",
		"/contests/abc153/tasks/abc153_c",
		"/contests/abc153/tasks/abc153_d",
		"/contests/abc153/tasks/abc153_e",
		"/contests/abc153/tasks/abc153_f",
	}, tasksPaths)
}

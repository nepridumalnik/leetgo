package main

import (
	"flag"

	"github.com/nepridumalnik/leetgo/internal/task"
)

func main() {
	level := flag.String("level", "", "input a level easy | medium | hard")
	leetcodeURL := flag.String("url", "", "input a LeetCode URL: https://leetcode.com<task-id>")

	flag.Parse()

	if level == nil || *level == "" {
		panic("level not provided")
	}
	if leetcodeURL == nil || *leetcodeURL == "" {
		panic("URL not provided")
	}

	if err := task.ValidateLevel(*level); err != nil {
		panic(err)
	}

	data := task.Data{
		Level: *level,
	}

	err := task.ParseURL(*leetcodeURL, &data)
	if err != nil {
		panic(err)
	}

	if err := task.CreateScaffold(&data); err != nil {
		panic(err)
	}
}

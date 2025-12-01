package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/milindmadhukar/Advent-Of-Code/2025/golang/solutions"
	"github.com/pkg/profile"
)

func main() {
	day := flag.String("day", "", "Specify the day to run")
	all := flag.Bool("all", false, "Run all solutions")
	help := flag.Bool("help", false, "Show help message")
	profileVal := flag.String("profile", "", "Enable Profiling")
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	var profiler interface{ Stop() }

	if *profileVal != "" {
		var profileType func(*profile.Profile)

		switch *profileVal {
		case "cpu":
			profileType = profile.CPUProfile
		case "mem":
			profileType = profile.MemProfile
		case "block":
			profileType = profile.BlockProfile
		case "mutex":
			profileType = profile.MutexProfile
		case "goroutine":
			profileType = profile.GoroutineProfile
		case "clock":
			profileType = profile.TraceProfile
		case "trace":
			profileType = profile.TraceProfile
		default:
			log.Fatalf("Invalid profile type: %s", *profileVal)
		}

		// Run the go pprof command
		profiler = profile.Start(profileType, profile.ProfilePath("./profiles"))
	}

	if *all {
		log.Println("Running all solutions")
		solutions.RunAllSolutions()
	} else if *day != "" {
		dayInt, err := strconv.Atoi(*day)
		if err != nil {
			log.Fatalf("Invalid day: %s", *day)
		}
		solutions.RunSolutionForDay(dayInt)
	} else {
		solutions.RunTodaysSolution()
	}

	if *profileVal != "" {
		profiler.Stop()
		args := []string{"tool", "pprof", "-http=:6060", fmt.Sprintf("profiles/%s.pprof", *profileVal)}
		if *profileVal == "trace" {
			args = []string{"tool", "trace", "profiles/trace.out"}
		}
		cmd := exec.Command("go", args...)
		err := cmd.Run()
		if err != nil {
			panic(err.Error())
		}
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  -day <number>: Run solution for a specific day")
	fmt.Println("  -all: Run all solutions")
	fmt.Println("  -help: Show this help message")
	fmt.Println("  -profile: Enable CPU profiling (pprof server on localhost:6060)")
	fmt.Println("If no flags are provided, today's solution will be run.")
}

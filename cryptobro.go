package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

// printBanner prints the ASCII art banner for the tool
func printBanner() {
	banner := `

  ██████╗██████╗ ██╗   ██╗██████╗ ████████╗ ██████╗ ██████╗ ██████╗  ██████╗ 
 ██╔════╝██╔══██╗╚██╗ ██╔╝██╔══██╗╚══██╔══╝██╔═══██╗██╔══██╗██╔══██╗██╔═══██╗
 ██║     ██████╔╝ ╚████╔╝ ██████╔╝   ██║   ██║   ██║██████╔╝██████╔╝██║   ██║
 ██║     ██╔══██╗  ╚██╔╝  ██╔═══╝    ██║   ██║   ██║██║  ██║██╔══██╗██║   ██║
 ╚██████╗██║  ██║   ██║   ██║        ██║   ╚██████╔╝██████╔╝██║  ██║╚██████╔╝
  ╚═════╝╚═╝  ╚═╝   ╚═╝   ╚═╝        ╚═╝    ╚═════╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝ 

>> Simulate crypto mining without the need for sus packages from the internet <<
`
	fmt.Println(banner)
}

// cpuIntensive performs a CPU-intensive calculation.
func cpuIntensive() {
	for i := 0; i < 100000; i++ {
		_ = 424242 * 424242
	}
}

func main() {
	// Print the banner
	printBanner()

	// Define command-line flags
	durationFlag := flag.Int("t", 5, "Duration in minutes for the CPU test to run")
	helpMessage := `CryptoBro - Crypto Mining Simulator

This tool simulates cryptocurrency mining by hammering your CPU with intensive calculations.
It uses all available CPU cores and provides real-time monitoring of CPU usage.
Perfect for stress testing your system without downloading suspicious packages from the internet.

Usage:
  cryptobro [options]

Options:
  -t int     Duration in minutes for the mining simulation to run (default: 5)
  -h         Display this help message

Examples:
  cryptobro -t 10    # Mine for 10 minutes
  cryptobro -h       # Show this help message`

	// Custom help flag handling
	flag.Usage = func() {
		printBanner()
		fmt.Println(helpMessage)
	}

	// Parse the flags
	flag.Parse()

	// Get the process ID
	pid := os.Getpid()

	// Get the number of CPU cores
	numCores := runtime.NumCPU()

	// Create a channel to signal when to stop
	done := make(chan int)

	// Get duration from flag
	duration := time.Duration(*durationFlag) * time.Minute

	// Start a goroutine for each CPU core
	for i := 0; i < numCores; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					// Perform the CPU-intensive function
					cpuIntensive()
				}
			}
		}()
	}

	fmt.Printf("Hammering CPU with PID %d. Press Ctrl+C to stop.\n", pid)
	fmt.Printf("Test will run for %s.\n", duration)

	// Handle Ctrl+C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Use a timer to stop the test after the specified duration
	timer := time.NewTimer(duration)

	// Start a goroutine to print CPU usage and time left dynamically
	startTime := time.Now()
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				// Get CPU usage
				cpuPercent, err := cpu.Percent(time.Second, false)
				if err != nil {
					fmt.Println("Error getting CPU usage:", err)
				} else {
					// Calculate time left
					elapsed := time.Since(startTime)
					timeLeft := duration - elapsed
					// Move the cursor up one line, clear the line, and then print the CPU usage
					fmt.Printf("\033[1A\033[K\033[31mCPU Usage: %.2f%%\033[0m\n", cpuPercent[0])
					fmt.Printf("Time Left: %s\r", timeLeft.Round(time.Second))
				}
				time.Sleep(time.Second)
			}
		}
	}()

	// Wait for either a signal or the timer to expire
	select {
	case <-c:
		fmt.Println("Test interrupted by user.")
	case <-timer.C:
		fmt.Println("Test finished after specified duration.")
	}

	// Signal the goroutines to stop
	close(done)

	fmt.Println("CPU usage test stopped.")
}

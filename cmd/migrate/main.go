package main

import (
	"fmt"
	"os"
	"time"
)

// A stand-in for real migrations: it prints its steps and exits 0, or 1 with MIGRATE_FAIL set.
func main() {
	for _, step := range []string{"001_users", "002_orders", "003_index_orders_created_at"} {
		fmt.Println("applying", step)
		time.Sleep(300 * time.Millisecond)
	}
	if os.Getenv("MIGRATE_FAIL") != "" {
		fmt.Fprintln(os.Stderr, "migration 004 failed: relation \"payments\" already exists")
		os.Exit(1)
	}
	fmt.Println("migrations done")
}

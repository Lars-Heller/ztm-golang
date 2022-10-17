//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

// * Create a function to print server status displaying:
//   - number of servers
//   - number of servers for each status (Online, Offline, Maintenance, Retired)
func printServers(servers map[string]int) {
	fmt.Println("We have ", len(servers), "servers")
	stats := make(map[int]int)
	for _, status := range servers {
		stats[status]++
	}
	fmt.Println("Online", stats[Online])
	fmt.Println("Offline", stats[Offline])
	fmt.Println("Maintenance", stats[Maintenance])
	fmt.Println("Retired", stats[Retired])
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatuses := make(map[string]int)
	//--Requirements:
	//* Create a map using the server names as the key and the server status
	//  as the value
	//* Set all of the server statuses to `Online` when creating the map
	for _, server := range servers {
		serverStatuses[server] = Online
	}
	//* After creating the map, perform the following actions:
	//  - call display server info function
	printServers(serverStatuses)
	//  - change server status of `darkstar` to `Retired`
	serverStatuses["darkstar"] = Retired
	// - change server status of `aiur` to `Offline`
	serverStatuses["aiur"] = Offline
	// - call display server info function
	printServers(serverStatuses)
	// - change server status of all servers to `Maintenance`
	for server := range serverStatuses {
		serverStatuses[server] = Maintenance
	}
	// - call display server info function
	printServers(serverStatuses)
}

package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/plain")
    containerID, _ := os.ReadFile("/proc/self/cgroup")
    containerID = []byte(strings.TrimSuffix(strings.Split(string(containerID), "\n")[len(strings.Split(string(containerID), "\n"))-2], "\n"))
    containerID = []byte(strings.Split(string(containerID), "/")[2])
    hostname, _ := os.Hostname()
    message := fmt.Sprintf("Container ID: %s\nHostname: %s\n", containerID, hostname)
    fmt.Fprintln(w, message)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}

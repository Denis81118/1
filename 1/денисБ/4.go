package main

import "fmt"

func countVotes(VOT []string) {

    candidat := []string{"Анна", "Борис", "Виктор"}
    
    counts := make(map[string]int)
    total := len(VOT)
   
    for _, vote := range VOT {
        counts[vote]++
    }
    for _, candidate := range candidat {
        votesCount := counts[candidate]
        per := float64(votesCount) * 100 / float64(total)
        fmt.Println(candidate, ":", votesCount, "голосов,", per, "%")
    }
}

func main() {
    VOT := []string{"Анна", "Борис", "Анна", "Виктор", "Борис", "Анна", "Анна"}
    countVotes(VOT)
}
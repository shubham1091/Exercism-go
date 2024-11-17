package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
    if initialVotes < 0 {
        initialVotes = 0  // Protect against negative initial votes
    }
    return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if counter == nil {
        return 0
    }
    return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
    if counter == nil || increment < 0 {
        return  // Protect against nil pointer and negative increments
    }
    *counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
    if candidateName == "" {
        candidateName = "Unknown"  // Provide default name
    }
    if votes < 0 {
        votes = 0  // Protect against negative votes
    }
    return &ElectionResult{
        Name:  candidateName,
        Votes: votes,
    }
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
    if result == nil {
        return "No result available"  // Handle nil result
    }
    return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
    if results == nil || results[candidate] <= 0 {
        return  // Protect against nil map and negative votes
    }
    if _, exists := results[candidate]; exists {
        results[candidate]--
    }
}

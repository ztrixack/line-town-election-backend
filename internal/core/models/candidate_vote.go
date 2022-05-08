package models

type CandidateVote struct {
	Model
	CandidateId int `json:"candidateId"`
	VotedCount  int `json:"votedCount"`
}

type CandidateVoteResponse struct {
	Id          int    `json:"id"`
	CandidateId int    `json:"candidateId"`
	VotedCount  string `json:"votedCount"`
}
package api

import "time"

// RoundState represents the round state at time of query
type RoundState struct {
	Height             int         `json:"height"`
	Round              int         `json:"round"`
	Step               int         `json:"step"`
	StartTime          time.Time   `json:"start_time"`
	CommitTime         time.Time   `json:"commit_time"`
	Validators         Validators  `json:"validators"`
	Proposal           interface{} `json:"proposal"`
	ProposalBlock      interface{} `json:"proposal_block"`
	ProposalBlockParts interface{} `json:"proposal_block_parts"`
	LockedRound        int         `json:"locked_round"`
	LockedBlock        interface{} `json:"locked_block"`
	LockedBlockParts   interface{} `json:"locked_block_parts"`
	ValidRound         int         `json:"valid_round"`
	ValidBlock         interface{} `json:"valid_block"`
	ValidBlockParts    interface{} `json:"valid_block_parts"`
	Votes              []Vote      `json:"votes"`
	CommitRound        int         `json:"commit_round"`
	LastCommit         interface{} `json:"last_commit"`
	LastValidators     Validators  `json:"last_validators"`
}

// Vote represents one validator vote
type Vote struct {
	Round              int      `json:"round"`
	Prevotes           []string `json:"prevotes"`
	PrevotesBitArray   string   `json:"prevotes_bit_array"`
	Precommits         []string `json:"precommits"`
	PrecommitsBitArray string   `json:"precommits_bit_array"`
}

// Validators represents the current round validators and proposer
type Validators struct {
	Validators []Validator `json:"validators"`
	Proposer   Validator   `json:"proposer"`
}

// Validator represents an individual validator node
type Validator struct {
	Address     string `json:"address"`
	PubKey      PubKey `json:"pub_key"`
	VotingPower int    `json:"voting_power"`
	Accum       int    `json:"accum"`
}

// PubKey represents a validator public key
type PubKey struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

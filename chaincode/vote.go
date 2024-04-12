package main

import (
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type VoteSmartContract struct {
	contractapi.Contract
}

type Vote struct {
	ID        string `json:"id"`
	Candidate string `json:"candidate"`
}

func (pc *VoteSmartContract) CountVotes(ctx contractapi.TransactionContextInterface) (int, error) {
	voteIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return 0, err
	}
	defer voteIterator.Close()
	count := 0
	for voteIterator.HasNext() {
		_, err := voteIterator.Next()
		if err != nil {
			return 0, err
		}
		count = count + 1
	}
	return count, nil
}

func (pc *VoteSmartContract) AddVote(ctx contractapi.TransactionContextInterface, candidate string) error {
	count, err := pc.CountVotes(ctx)
	id := strconv.Itoa(count + 1)
	if err != nil {
		return err
	}
	vote := Vote{
		ID:        id,
		Candidate: candidate,
	}
	voteJSON, err := json.Marshal(vote)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(vote.ID, voteJSON)
	return err
}

func (pc *VoteSmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	votes := []Vote{
		{ID: "1", Candidate: "Ice Cream"},
		{ID: "2", Candidate: "Pizza"},
		{ID: "3", Candidate: "Pizza"},
		{ID: "4", Candidate: "Pizza"},
		{ID: "5", Candidate: "Hot Dogs"},
		{ID: "6", Candidate: "Hot Dogs"},
		{ID: "7", Candidate: "Hot Dogs"},
		{ID: "8", Candidate: "Hot Dogs"},
		{ID: "9", Candidate: "Salad"},
		{ID: "10", Candidate: "Salad"},
	}

	for _, vote := range votes {
		voteJSON, err := json.Marshal(vote)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(vote.ID, voteJSON)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pc *VoteSmartContract) TallyVotes(ctx contractapi.TransactionContextInterface) (map[string]int, error) {
	voteIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer voteIterator.Close()
	tally := map[string]int{}
	for voteIterator.HasNext() {
		voteResponse, err := voteIterator.Next()
		if err != nil {
			return nil, err
		}
		var vote *Vote
		err = json.Unmarshal(voteResponse.Value, &vote)
		if err != nil {
			return nil, err
		}
		if count, ok := tally[vote.Candidate]; ok {
			tally[vote.Candidate] = count + 1
		} else {
			tally[vote.Candidate] = 1
		}
	}
	return tally, nil
}

func (pc *VoteSmartContract) QueryAllVotes(ctx contractapi.TransactionContextInterface) ([]*Vote, error) {
	voteIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer voteIterator.Close()
	var votes []*Vote
	for voteIterator.HasNext() {
		voteResponse, err := voteIterator.Next()
		if err != nil {
			return nil, err
		}

		var vote *Vote
		err = json.Unmarshal(voteResponse.Value, &vote)
		if err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func main() {
	voteSmartContract := new(VoteSmartContract)
	cc, err := contractapi.NewChaincode(voteSmartContract)
	if err != nil {
		panic(err.Error)
	}
	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}

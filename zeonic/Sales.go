package zeonic

import (
	"fmt"
)

const JOKER = "MSN-02"

type Rfp interface {
	GetCodes() []string
}

type Sales struct {
	Name     string
	Proposal []string
}

func (s *Sales) CheckRfp(rfp Rfp) {
	wants := rfp.GetCodes()
	s.Proposal = wants
	s.Proposal = append(s.Proposal, JOKER)
	fmt.Printf("%v\n", s.Proposal)
}

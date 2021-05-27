package main

import "fmt"

// 多态

type MemberRights interface {
	Information() string
}

type BronzeMember struct {
	Discount uint8
}

type SilverMember struct {
	Discount uint8
}

type GoldMember struct {
	Discount uint8
}

func (b *BronzeMember) Information() string {
	return fmt.Sprintf("Discount:%d", b.Discount)
}

func (s *SilverMember) Information() string {
	return fmt.Sprintf("Discount:%d", s.Discount)
}

func (g *GoldMember) Information() string {
	return fmt.Sprintf("Discount:%d", g.Discount)
}

func Price(m MemberRights) {
	fmt.Println(m.Information())
}

func main() {
	b := &BronzeMember{Discount: 9}
	Price(b)
	s := &SilverMember{8}
	Price(s)
	g := new(GoldMember)
	g.Discount = 7
	Price(g)
}

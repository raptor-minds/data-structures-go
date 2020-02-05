package main

import (
	"fmt"
)

type Polynomial struct {
}

type PolyNode struct {
	coef  int
	expon int
	link  *PolyNode
}

func main() {
	var P1, P2, PS *PolyNode
	var p *Polynomial
	// first param is the num of polynomials
	P1Input := []int{2, 12, 4, 5, 3}
	P2Input := []int{3, -8, 5, 2, 4, 6, 2}
	P1 = p.ReadPoly(P1Input)
	if P1 == nil {
		fmt.Println("Please input polynomial p1 correctly")
	}
	P2 = p.ReadPoly(P2Input)
	if P2 == nil {
		fmt.Println("Please input polynomial p2 correctly")
	}
	PS = p.Add(P1, P2)
	p.PrintPl(PS)
	//PP = p.Multiply(P1, P2)
	//p.PrintPl(PP)
}

func (p *Polynomial) ReadPoly(input []int) *PolyNode {
	var P, pRear *PolyNode
	P = &PolyNode{}
	pRear = P
	i := 1
	for N := input[0]; N >= 0; N = N - 2 {
		coef, expon := input[i], input[i+1]
		pRear = p.Attach(coef, expon, pRear)
		i += 2
	}
	return P
}

func (p *Polynomial) Attach(c int, e int, pRear *PolyNode) *PolyNode {
	pNew := PolyNode{coef: c, expon: e, link: nil}
	fmt.Println(pRear)
	pRear.link = &pNew
	pRear = &pNew
	return pRear
}

func (p *Polynomial) Add(p1 *PolyNode, p2 *PolyNode) *PolyNode {

	P := &PolyNode{}
	var pResult *PolyNode
	pResult = P
	for ; p1.link != nil && p2.link != nil; {
		switch p.Compare(p1, p2) {
		case 1:
			pResult = p.Attach(p1.coef, p1.expon, pResult)
			p1 = p1.link
		case -1:
			pResult = p.Attach(p2.coef, p2.expon, pResult)
			p2 = p2.link
		case 0:
			pResult = p.Attach(p1.coef+p2.coef, p2.expon, pResult)
			p1 = p1.link
			p2 = p2.link
		}
	}

	for ; p1.link != nil; p1 = p1.link {
		p.Attach(p1.coef, p1.expon, pResult)
	}

	for ; p2.link != nil; p2 = p2.link {
		p.Attach(p2.coef, p2.expon, pResult)
	}

	return P
}

func (p *Polynomial) Multiply(p1 *PolyNode, p2 *PolyNode) *PolyNode {
	return nil
}

func (p *Polynomial) PrintPl(polyNode *PolyNode) {
	for ; polyNode.link != nil; polyNode = polyNode.link {
		fmt.Println(fmt.Sprintf("coef is %d, expon is %d", polyNode.coef, polyNode.expon))
	}
}

func (p *Polynomial) Compare(p1 *PolyNode, p2 *PolyNode) int {
	if p1.expon < p2.expon {
		return -1
	} else if p1.expon > p2.expon {
		return 1
	}
	return 0
}

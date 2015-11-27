package qrtickets

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
	"fmt"
)

type MyKey struct {
	*ecdsa.PrivateKey
}

type MyCurveParams struct {
	*elliptic.CurveParams
}

func (p *MyKey) GetParams() {
	curveParams := p.Curve.Params()

	// Print out Curve Parameters
	fmt.Println("P ", curveParams.P)
	fmt.Println("N ", curveParams.N)
	fmt.Println("B ", curveParams.B)
	fmt.Printf("Gx, Gy : %v, %v\n", curveParams.Gx, curveParams.Gy)
	fmt.Println("BitSize : ", curveParams.BitSize)

	// Now Print out Key Parameters
	fmt.Println("X ", p.X)
	fmt.Println("Y ", p.Y)
	fmt.Println("D ", p.D)
	
}

func (c *MyCurveParams) SetParams (p, n, b, gx, gy *big.Int, bitsize int, name string) {
	c.P = p
	c.N = n
	c.B = b
	c.Gx = gx
	c.Gy = gy
	c.BitSize = bitsize
//	c.Name = name
}

func InitializeKey() {
	config := LoadConfig()
	fmt.Printf("%#v",config)
}

package day04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Puzzle struct{}

const secretkey string = "bgvyzdsv"

func (p *Puzzle) Part1() string {
	var num int64 = 0
	for {
		str := fmt.Sprintf("%s%d", secretkey, num)
		if hasFiveLeadingZeros(md5Hash(str)) {
			return fmt.Sprintf("%d", num)
		}
		num++
	}
}

func (p *Puzzle) Part2() string {
	var num int64 = 0
	for {
		str := fmt.Sprintf("%s%d", secretkey, num)
		if hasSixLeadingZeros(md5Hash(str)) {
			return fmt.Sprintf("%d", num)
		}
		num++
	}
}

func New() *Puzzle {
	return &Puzzle{}
}

func md5Hash(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func hasFiveLeadingZeros(s string) bool {
	if len(s) < 5 {
		return false
	}
	return s[:5] == "00000"
}

func hasSixLeadingZeros(s string) bool {
	if len(s) < 6 {
		return false
	}
	return s[:6] == "000000"
}

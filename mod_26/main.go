package main

import (
	"flag"
	"fmt"
	"strings"
)

const (
	first_lowercase  = 'a'
	last_lowercase   = 'z'
	first_uppdercase = 'A'
	last_uppercase   = 'Z'
	default_falg     = "cvpbPGS{arkg_gvzr_V'yy_gel_2_ebhaqf_bs_ebg13_uJdSftmh}"
)

func main() {
	flag_rot13 := flag.String("f", default_falg, "flag to pass to rot13")
	flag.Parse()

	fmt.Println("Flag:", rot13(*flag_rot13))
}

func rot13(s string) string {
	b := strings.Builder{}

	for _, char := range s {
		if char >= first_lowercase && char <= last_lowercase {
			b.WriteRune(moveCharchater13(char, last_lowercase))
		} else if char >= first_uppdercase && char <= last_uppercase {
			b.WriteRune(moveCharchater13(char, last_uppercase))
		} else {
			b.WriteRune(char)
		}
	}

	return b.String()
}

func moveCharchater13(char, bound rune) rune {
	c := char + 13
	alphabets_length := rune(26)

	if c > bound {
		c = char - alphabets_length + 13
	}

	return c
}

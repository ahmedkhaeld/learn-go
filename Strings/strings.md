# Strings
need to
think about strings in a logical way in
a physical way
and the reason for that is that strings
in GO are all unicode
and unicode is a particular technique
that allows us to represent
international characters,<br> in the old days programming
languages they all used something called
ascii
it represented characters with seven
bits and it basically only represented
the characters of american english so all the characters fit into one byte
<br> when we move to international
languages we get non-roman languages like
chinese or arabic
and we need very different techniques to
represent those
unicode is a way to represent them and
it uses numbers that are bigger than
what fits into a byte

**a rune** is a synonym for a 32-bit(int32): because
that four byte it is big enough to
represent any unicode could point
any character any logical character

in order to make programs
efficient we don't want to represent
every character all the time with four
bytes because the reality is
a lot of these programs are just going
to have ascii characters
and so there's a technique for encoding
unicode called utf-8

**byte** is just a synonym for an 8-bit integer

a string is physically a
sequence of the bytes
that are required to encode the unicode
characters that are there logically the
runes

## demonstrate this concept of bytes vs runes

```go
package main

import "fmt"

func main() {
	s := "élite"
	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	// cast the s to sequence of runes
	fmt.Printf("%8T %[1]v %d\n", []rune(s), len([]rune(s)))
	// cast s to a sequence of bytes
	fmt.Printf("%8T %[1]v %d\n", []byte(s), len([]byte(s)))

}

// output:
//string élite  6
//[]int32 [233 108 105 116 101]  6
//[]uint8 [195 169 108 105 116 101]  6


```

##### when we ask the length of string we are going to get the physical answer
the length of a string in the program is
the length of the byte string <br>

logically five
printable characters it's physically six
bytes in utf-8 encoding

we actually have to deal with the actual
memory and so it just makes more sense
to ask the length of the string how many
bytes because typically that's a bigger
number, so again the length of a string is
the number
of bytes required to represent the
unicode characters not the number of
unicode characters
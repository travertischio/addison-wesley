// Computes counts of characters
package main
import (
  "bufio"
  "fmt"
  "io"
  "os"
  "unicode"
  "unicode/utf8"
)

func main() {
  charcount()
  wordfreq()
}

func charcount() {
  counts := make(map[rune]int)
  types := make(map[string]int)
  var utflen [utf8.UTFMax + 1]int
  invalid := 0

  in := bufio.NewReader(os.Stdin)
  for {
    r, n, err := in.ReadRune()
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
      os.Exit(1)
    }
    if r == unicode.ReplacementChar && n == 1 {
      invalid++
      continue
    }
    if unicode.IsLetter(r) {
      types["letter"]++
    } else if unicode.IsDigit(r) {
      types["digit"]++
    } else if unicode.IsSymbol(r) {
      types["symbol"]++
    } else if unicode.IsControl(r) {
      types["control"]++
    }
    counts[r]++
    utflen[n]++
  }
  fmt.Printf("type\tcount\n")
  for t, n := range types {
    fmt.Printf("%s\t%d\n", t, n)
  }
  fmt.Printf("\nrune\tcount\n")
  for c, n := range counts {
    fmt.Printf("%q\t%d\n", c, n)
  }
  fmt.Printf("\nlen\tcount\n")
  for i, n := range utflen {
    if i > 0 {
      fmt.Printf("%d\t%d\n", i, n)
    }
  }
  if invalid > 0 {
    fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
  }
}

func wordfreq() {
  counts := make(map[string]int)

  in := bufio.NewScanner(os.Stdin)
  in.Split(bufio.ScanWords)
  for in.Scan() {
    word := in.Text()
    counts[word]++
  }
  fmt.Printf("word\tcount\n")
  for w, n := range counts {
    fmt.Printf("%s\t%d\n", w, n)
  }
}

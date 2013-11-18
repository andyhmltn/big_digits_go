package main

import (
  "fmt"
  "os"
  "path/filepath"
)

var big_digits = [][]string{
  {"  000  "," 0   0 ","0     0","0     0","0     0"," 0   0 ","  000  "},
  {"  111  ","11  1  ","    1  ","    1  ","    1  ","    1  ","1111111"},
  {"  222  "," 2   2 ","     2 ","    2  ","   2   ","  2    "," 22222 "},
  {" 33333 ","     3 ","    3  ","  333  ","    3  ","     3 "," 33333 "},
  {"   444 ","  4  4 "," 4   4 "," 444444","     4 ","     4 ","     4 "},
  {" 55555 "," 5     "," 5     ","  555  ","     5 ","     5 "," 55555 "},
  {"  666  "," 6     "," 6     "," 6666  "," 6   6 "," 6   6 "," 66666 "},
  {"7777777","      7","     7 ","    7  ","   7   ","  7    "," 7     "},
  {"  888  "," 8   8 "," 8   8 ","  888  "," 8   8 "," 8   8 ","  888  "},
  {"  9999 "," 9    9"," 9    9"," 999999","      9","      9"," 99999 "},
  {"       ","       ","       ","       ","       ","       ","       "},
}

func main() {
  if len(os.Args) == 1 {
    fmt.Printf("usage:%s <whole-number>\n", filepath.Base(os.Args[0]))
    os.Exit(1)
  }

  stringOfDigits := os.Args[1]

  for row := range big_digits[0] {
    line := ""

    for column := range stringOfDigits {
      digit := stringOfDigits[column] - '0'

      if 0 <= digit && digit <= 9 {
        line += big_digits[digit][row] + " "
      } else {
        if stringOfDigits[column] == '_' {
          line += big_digits[len(big_digits)-1][row]
        }
      }
    }

    fmt.Println(line)
  }
}
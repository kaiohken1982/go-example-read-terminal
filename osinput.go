package main

import (
  "fmt" 
  "os"
  "strings"
  "bufio"
)

func main() {
  // exampleDuplicatesInFilesAllFile()
  exampleDuplicatesInFiles()
  // exampleDuplicates()
  // exampleForLoop()
  // exampleRange() 
  // exampleStringJoin()
}

/**
 * go run . testo1.txt
 */
func exampleDuplicatesInFiles() {
  counts := make(map[string]int)
  files := os.Args[1:]

  if len(files) == 0 {
    countLines(os.Stdin, counts)
  } else {
    for _, arg := range files {
      f, err := os.Open(arg)
      if err != nil {
        fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
        continue
      }
      countLines(f, counts)
      f.Close()
    }
  }

  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}

func countLines(f *os.File, counts map[string]int) {
  input := bufio.NewScanner(f)
  for input.Scan() {
    counts[input.Text()]++
  }
  // NOTE: ignoring potential errors from input.Err()
}

/**
 * go run . 
 * Aggiungere linee nel terminale, premere invio
 * ctrl + d per uscire
 */
func exampleDuplicates() {
  /**
   * The bui lt-in function make createsanew emp ty
   * map; it has other uses too
   */
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)

  // CTRL + d per far restituire false ad input.Scan() e terminare il loop
  for input.Scan() {
    counts[input.Text()]++
  }

  // NOTE: ignoring potential errors from input.Err()
  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}

func exampleForLoop() {
  var s, sep string
  for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
}

func exampleStringJoin() {
  fmt.Println(strings.Join(os.Args[1:], " | "))
}

func exampleRange() {
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }
  fmt.Println(s)
}
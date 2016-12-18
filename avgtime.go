package main

import "fmt"
import "flag"
import "time"
import "os/exec"
import "strings"

func main() {

  itrPtr := flag.Int("i", 10, "Iteration count")

  flag.Parse()
  fmt.Println("executing:", strings.Join(flag.Args(), " "))

  cmdName := flag.Args()[0]
  cmdArgs := flag.Args()[1:]

  total := time.Duration(0)

  for i := 1; i <= *itrPtr; i++ {
    start := time.Now()
    exec.Command(cmdName, cmdArgs...).Run()
    total = total + time.Now().Sub(start)
    fmt.Printf("\r%d/%d", i, *itrPtr)
  }

  fmt.Println("\n", time.Duration(int(total) / *itrPtr))
}

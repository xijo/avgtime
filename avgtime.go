package main

import "fmt"
import "flag"
import "time"
import "os/exec"
import "strings"
import "math"

func main() {

  itrPtr := flag.Int("i", 10, "Iteration count")

  flag.Parse()
  fmt.Println("executing:", strings.Join(flag.Args(), " "))

  cmdName := flag.Args()[0]
  cmdArgs := flag.Args()[1:]

  var samples []int

  for i := 1; i <= *itrPtr; i++ {
    start := time.Now()
    exec.Command(cmdName, cmdArgs...).Run()
    samples = append(samples, int(time.Now().Sub(start)))
    fmt.Printf("\r%d/%d", i, *itrPtr)
  }
  fmt.Printf("\nAvg time: %15s", time.Duration(average(samples)))
  fmt.Printf("\nSt. dev.: %15s", time.Duration(standard_deviation(samples)))
}


func average(arr[]int) int {
  total := 0
  for _, v := range arr {
    total += v
  }
  return total/len(arr)
}

func standard_deviation(arr[]int) int {
  avg := average(arr)
  s := 0
  for _, v := range arr {
    s += int(math.Pow(float64(v - avg), 2))
  }
  return int(math.Sqrt(float64(s/len(arr))))
}

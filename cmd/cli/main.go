package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"
  "strconv"

	ligma "github.com/marcomaiermm/flextime/internal"
)

type stringFlag struct {
  value string
  set bool
}

func (sf *stringFlag) Set(x string) error {
    sf.value = x
    sf.set = true
    return nil
}

func (sf *stringFlag) String() string {
    return sf.value
}


func run(ctx context.Context, w io.Writer) error {
  _, cancel := signal.NotifyContext(ctx, os.Interrupt)
  defer cancel()
  var startTimeFlag, hoursFlag, breakTimeFlag stringFlag

  flag.Var(&startTimeFlag, "start", "The start time in the format of HH:MM")
  flag.Var(&hoursFlag, "hours", "The number of hours you want to work")
  flag.Var(&breakTimeFlag, "break", "The number of hours you want to have a break")
  flag.Parse()

  if !hoursFlag.set && !startTimeFlag.set {
    return fmt.Errorf("you need to provide either a start time or the number of hours you want to work")
  }

  if hoursFlag.set {
    hours, err := strconv.ParseFloat(hoursFlag.value, 64)
    if err != nil {
      return err
    }

    conf, err := ligma.NewConf()
    if err != nil {
      return err
    }
    conf.WorkHours = hours
    err = conf.WriteToConfig(*conf)
    if err != nil {
      return err
    }
  }

  if breakTimeFlag.set {
    breakTime, err := strconv.ParseFloat(breakTimeFlag.value, 64)
    if err != nil {
      return err
    }

    conf, err := ligma.NewConf()
    if err != nil {
      return err
    }
    conf.BreakTime = breakTime
    err = conf.WriteToConfig(*conf)
    if err != nil {
      return err
    }
  }

  startTime := time.Now().Format("15:04")
  if startTimeFlag.set {
    startTime = startTimeFlag.value
  }

  start, err := ligma.FormatToTime(startTime)
  if err != nil {
    return err
  }

  next, err := ligma.TimeToGo(*start)
  if err != nil {
    return err
  }

  fmt.Fprint(w, next.Format("15:04"))

  return nil
}

func main() {
  ctx := context.Background()
  if err := run(ctx, os.Stdout); err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err)
    os.Exit(1)
  }
}

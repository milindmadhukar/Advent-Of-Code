package models

import (
	"time"
)

type Solution interface {
  Part1() any
  Part2() any
  TimeTaken() time.Duration
}

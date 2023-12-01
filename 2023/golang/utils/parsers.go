package utils

import "strings"

func ParseFromString(data string) []string {
  data = strings.Trim(data, " ")
  data = strings.Trim(data, "\n")
  return strings.Split(data, "\n")
}

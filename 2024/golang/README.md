## 2024 Golang Solutions

### Instructions to run

1. Create a "sessionkey.txt" at the same level as `main.go` and paste the `session` cookie (without the `session=`)
1. To run all solutions, run `go run main.go --all`
1. To run solution for a given day, Run `go run main.go --day <day>`
1. If no flags are given, today's solution will be ran.


#### You will find your session key by:
1. Going to [Advent Of Code](https://adventofcode.com/)
1. On your browser, Inspect Element `(Ctrl + Shift + I)`
1. Click on Network Tab
1. `Ctrl + R` to open network activity
1. Click on any request that goes to adventofcode.com
1. Go to Request Headers, scroll down and copy the value of `session` from the **Cookies** part.

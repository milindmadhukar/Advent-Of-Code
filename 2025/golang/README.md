## 2025 Golang Solutions

### Instructions to run
1. Obtain your Advent Of Code session key
1. Create a `sessionkey.txt` file at the same level as `main.go` and paste the `session` cookie (without the `session=`)
1. To run all solutions, run `go run main.go --all`
1. To run solution for a given day, Run `go run main.go --day <day>`
1. If no flags are given, today's solution will be ran.
1. (Optional) Add the `--profile <type>` (eg. cpu) to profile the solution


#### Steps to find you session key:
1. Go to [Advent Of Code](https://adventofcode.com/)
1. On your browser, Inspect Element `(Ctrl + Shift + I)`
1. Click on Network Tab
1. `Ctrl + R` to open network activity
1. Click on any request that goes to adventofcode.com
1. Go to Request Headers, scroll down and copy the value of `session` from the **Cookies** part.

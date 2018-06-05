# Challenges

These are problems delivered through [https://dailycodingproblem.com](https://dailycodingproblem.com)

The solutions are my own, not official.

### Dev

1. Typically, create a subfolder for an issue.
2. Put the issue content into a README.md
3. Add a solution file (ex.  solution.go)
4. Add a test file (ex.  solution_test.go)

#### Recommendations

1. start with a brute force solution `func SolutionBruteForce()`
2. test the solution `func TestBruteForce()`, `go test`
3. Improve the solution `func Solution()`
4. test the solution `go test`
5. write benchmarks `func BenchmarkSolution()`
6. benchmark improvements `go test -bench *`

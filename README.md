## Solutions in Go for the Coding Problems from Cassidy Williams' Newsletter

This repository contains solutions written in Go, for the **interview question of the week** section problems published in the newsletter [Rendezvous with cassidoo](https://cassidoo.co/newsletter/)

### Solutions

- [Hyperfactorial](hyperfactorial): From newsletter Gloria Steinem (2025-09-29)
- [Fireworks](fireworks): From newsletter Thelonious Monk (2025-07-07)
- [No Repeats](norepeats): From newsletter E.E. Cummings (2025-06-30)

### Directory Structure
The solution to each problem is contained as a standalone package in its own directory. Each directory contains:
- `README.md`: Problem description and solution details.
- `<problemName>.go`: The Go implementation of the solution.
- `<problemName>_test.go`: Unit tests for the solution.

### How to Run the Tests

1. Clone the repository.
2. Change the directory to the cloned repository:
   ```bash
   cd /path/to/your/cloned/repo
   ```
3. Run the tests using the following command:
   ```bash
   go test -v ./...
   ```

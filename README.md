# Advent of Code solutions in Go
## How to run
You need `go version go1.21.x` to run the solutions.
Uncomment your wanted solution in `main.go` and run it with `go run main.go` in the root dir of this repository.
The input has to be put in a `input.txt` file in for example: `./2023/day01/input.txt`.
## Solved Puzzles
- Solutions for 2023
  - [Day 01](2023/day01/count.go)
  - [Day 02](2023/day02/games.go) 
  - [Day 03](2023/day03/engine.go) 
  - [Day 04](2023/day04/scratchcards.go)
    - Solved part 2 with recursion
  - [Day 05](2023/day05/garden.go)
    - Solved part 1 and 2 with an execution time of 2m38s and 20GB RAM usage
    - Optimized part 2 with an execution time of 2m02s and 2GB RAM usage
  - [Day 06](2023/day06/race.go)
  - [Day 07](2023/day07/camel.go)
  - [Day 08](2023/day08/nodes.go)
    - LCM Solution
    - Maybe other solutions: [chinese remainder theorem](https://en.wikipedia.org/wiki/Chinese_remainder_theorem), [cycle detection](https://en.wikipedia.org/wiki/Cycle_detection#Tortoise_and_hare)
  - [Day 09](2023/day09/history.go)
    - Solved with recursion for [binomial coefficient](https://en.wikipedia.org/wiki/Binomial_coefficient#Pascal's_triangle)
  - [Day 10](2023/day10/pipes.go)
    - Inclusive visual representation of pipe loop
  - [Day 11](2023/day11/galaxy.go)
    - First solution didn't work. 4100 difference to correct answer in part 1. Couldn't find the bug, test result was correct
    - Rebuild it with another solution
  - [Day 12](2023/day12/spring.go)
    - Solved with recursion
    - Without cache runtime 15min+ (didn't complete it)
    - Implemented cache, runtime < 1s
    - Thanks to [https://github.com/mebeim/aoc/](https://github.com/mebeim/aoc/tree/master/2023#part-2-9) for cache idea
  - [Day 17](2023/day17/crucible.go)
    - Used priority queue
    - Thanks to [teivah](https://github.com/teivah/advent-of-code/blob/lib/lib_ds_pq.go) for generic priority queue implementation
    - BFS (Breadth First Search) Solution
    - Resource on path finding [redblobgames](https://www.redblobgames.com/pathfinding/a-star/introduction.html)
  - [Day 18](2023/day18/dig.go)
    - The code for part 1 is purely from me and is suitable for that task. For part 2 it isn't going to cut it anymore.
    - Thanks to [teivah](https://github.com/teivah/) Part 2 is solved with [https://stackoverflow.com/a/717367](https://stackoverflow.com/a/717367) (EDIT section)
      - Another possible solution vector cross product

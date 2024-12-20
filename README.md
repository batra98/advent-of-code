# Advent of Code

Trying out Advent of Code with a twist.

I will randomly (weighted in reverse proportion of experience) select one language from the list for each problem.

1. C++ (weight: 1)
2. Rust (weight: 2)
3. Golang (weight: 2)
4. Python (weight: 1)
5. Haskell (weight: 4)
6. Julia (weight: 4)
7. Bash (weight: 1)

The script to choose language:

```Python
from numpy.random import choice
list_of_candidates = ["C++", "Rust", "Golang", "Python", "Haskell", "Julia", "Bash"]
p = [1/15,2/15,2/15,1/15,4/15,4/15,1/15]
draw = choice(list_of_candidates, 1, p=p)
```

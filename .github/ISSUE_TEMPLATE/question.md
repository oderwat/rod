---
name: Question
about: Title of your question.
title: ''
labels: question
assignees: ''
---

**Rod Version:** v0.0.0

## The code to demonstrate your question

1. Clone Rod to your local and cd to the repository:

   ```bash
   git clone https://github.com/go-rod/rod
   cd rod
   ```

1. Use your code to replace the content of function `TestLab` in file `lab_test.go`.

1. Test your code with: `go test -run TestLab`.

1. Replace this section with your entire `lab_test.go` content, like below:

```go
package rod_test

import (
    "testing"
)

func TestLab(t *testing.T) {
    g := setup(t)
    g.Eq(1, 1)
}
```

## What you expected to see

Such as what you want to do.

## What you instead got

Such as what error you see.

## What have you tried to solve the question

Such as after modifying some source code of Rod you are able to get rid of the problem.

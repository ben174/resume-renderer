package main

import (
    "fmt"
    "io"
    "gopkg.in/src-d/go-git.v3"
)

func main() {
    fmt.Printf("Pulling resume repository.\n")
    r, err := git.NewRepository("https://github.com/ben174/bugben", nil)
    if err != nil {
        panic(err)
    }

    if err := r.Pull("origin", "refs/heads/master"); err != nil {
        panic(err)
    }

    iter := r.Commits()
    defer iter.Close()

    for {
        commit, err := iter.Next()
        if err != nil {
            if err == io.EOF {
                break
            }

            panic(err)
        }

        fmt.Println(commit)
    }
}

package main

import (
    "fmt"
    "io"
    "gopkg.in/src-d/go-git.v3"
)

func main() {
    fmt.Printf("Pulling resume repository.\n")

    remote, err := git.NewRemote("https://github.com/ben174/bugben")

    if err != nil {
        panic(err)
    }

    if err := remote.Connect(); err != nil {
        panic(err)
    }

    fmt.Println(remote.Info().Head)


    repository, err := git.NewRepository("https://github.com/ben174/bugben", nil)
    if err != nil {
        panic(err)
    }

    if err := repository.Pull("origin", "refs/heads/master"); err != nil {
        panic(err)
    }

    head := remote.Info().Head

    fmt.Println("Fetching Head:", head)


    commit, err := repository.Commit(head)
    entries := commit.Tree().Entries


    for i, t := range entries {
      fmt.Println(i, t)
    }



    treewalker := git.NewTreeWalker(repository, commit.Tree())

    for {
      name, _, _, err := treewalker.Next()
      if err == io.EOF {
        break
      }
      fmt.Println(name)
    }



    /*
    if err := r.Pull("origin", "refs/heads/master"); err != nil {
        panic(err)
    }

    iter := r.Commits()

    defer iter.Close()

    commit, err := iter.Next()
    if err != nil {
        panic(err)
    }

    fmt.Println(commit.Message)
    */
}

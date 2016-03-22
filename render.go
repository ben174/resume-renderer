package main

import (
	"fmt"
	//"io"
	"encoding/json"

	"gopkg.in/src-d/go-git.v3"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type Resume struct {
	Profile    []string `json:"profile"`
	Experience []struct {
		StartDate       string   `json:"startDate"`
		EndDate         string   `json:"endDate"`
		Title           string   `json:"title"`
		URL             string   `json:"url"`
		Employer        string   `json:"employer"`
		Accomplishments []string `json:"accomplishments"`
		LogoURL         string   `json:"logoUrl"`
	} `json:"experience"`
	Expertise []struct {
		Category string   `json:"category"`
		Entries  []string `json:"entries"`
	} `json:"expertise"`
	Name  string `json:"name"`
	Links []struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"links"`
	URL       string `json:"url"`
	Email     string `json:"email"`
	Community []struct {
		ProjectURL string `json:"projectUrl"`
		Body       string `json:"body"`
		SrcURL     string `json:"srcUrl"`
		Summary    string `json:"summary"`
		Title      string `json:"title"`
	} `json:"community"`
}

func main() {
	resume := &Resume{}
	entry := &ExperienceEntry{}
	entry.Years = 10
	entry.Name = "balls"

	resume.Experience = []ExperienceEntry{entry}

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcD = append(slcD, "balls")
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

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

	/*

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

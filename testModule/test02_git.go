package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log"
	"os"
)

/*
参考：
https://www.codeleading.com/article/26252532481/
 */
func main() {

	cpath := "/mnt/d/ProjectGo/testModule/test02"
	username := "heliyun@ccbscf.com"
	password := "heliyun123"

	gitAuth := &http.BasicAuth{Username: username, Password: password}
	r := NewRepository(gitAuth, cpath)
	if r == nil {
		log.Fatal("Error, Return Resository Error!")
	}

	worktree, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	//if err := GitCheckout(worktree); err != nil {
	//	log.Fatal("git checkout preprd failed, ",err)
	//}

	if err := GitPull(gitAuth, worktree); err != nil {
		log.Fatal(err)
	}

}

func NewRepository(auth *http.BasicAuth, gitpath string) *git.Repository {

	if p := PathExists(gitpath); p == true {
		log.Printf("path: %s, exsit!", gitpath)
		return nil
	}

	branchname := plumbing.NewBranchReferenceName("preprd")
	rep, err := git.PlainClone(gitpath, false, &git.CloneOptions{
		RemoteName: "origin",
		URL: "http://gitlab.inner.com/heidan/nginx.git",
		Auth: auth,
		ReferenceName: branchname,
		SingleBranch: true,
		NoCheckout: false,
	})

	if err != nil {
		log.Println(err)
		return nil
	}
	return rep
}

func GitCheckout(w *git.Worktree) error {
	var branch plumbing.ReferenceName
	branch = "preprd"

	return w.Checkout(&git.CheckoutOptions{
		Branch: branch,
		Force: true,
	})
}

func GitPull(auth *http.BasicAuth, w *git.Worktree) error {
	log.Println("git pull ")

	return w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth: auth,
	})
}


func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
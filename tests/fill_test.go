package tests

import "testing"

type Repo struct {
	Domain     string `path:"domain"`
	Namespace  string `path:"namespace"`
	Repository string `path:"repository"`
}

var repoURL = "/swr/v2/domains/:domain/namespaces/:namespace/repositories/:repository"

func TestFillRestPath(t *testing.T) {
	repo := &Repo{
		Domain:     "aaaaa",
		Namespace:  "bbbbb",
		Repository: "cccccc",
	}
	t.Log(FillPath(repoURL, repo))
}

func BenchmarkFillRestPath(t *testing.B) {
	repo := &Repo{
		Domain:     "aaaaa",
		Namespace:  "bbbbb",
		Repository: "cccccc",
	}
	for i := 0; i < t.N; i++ {
		FillPath(repoURL, repo)
	}
}

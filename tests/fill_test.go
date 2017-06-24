package tests

import "testing"

type Repo struct {
	Domain string `restful:"domain"`
	Namespace string `restful:"namespace"`
	Repository string `restful:"repository"`
}

var repoURL = "/swr/v2/domains/:domain/namespaces/:namespace/repositories/:repository"

func TestFillRestPath(t *testing.T) {
	repo := Repo{
		Domain: "aaaaa",
		Namespace: "bbbbb",
		Repository: "cccccc",
	}
	t.Log(FillRestPath(repoURL, repo))
}

func BenchmarkFillRestPath(t *testing.B) {
	repo := Repo{
		Domain: "aaaaa",
		Namespace: "bbbbb",
		Repository: "cccccc",
	}
	t.Log(FillRestPath(repoURL, repo))
}
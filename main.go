package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Repository struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Description   string `json:"description"`
	HTMLURL       string `json:"html_url"`
	Language      string `json:"language"`
	DefaultBranch string `json:"default_branch"`
	Stars         int    `json:"stargazers_count"`
	Forks         int    `json:"forks_count"`
	OpenIssues    int    `json:"open_issues_count"`
	Watchers      int    `json:"watchers_count"`
	Private       bool   `json:"private"`
	Archived      bool   `json:"archived"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type GitHubClient struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

func NewGitHubClient() *GitHubClient {
	return &GitHubClient{
		BaseURL: "https://api.github.com",
		Token:   os.Getenv("GITHUB_TOKEN"),
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *GitHubClient) GetRepository(owner, repo string) (*Repository, error) {
	if owner == "" || repo == "" {
		return nil, errors.New("owner e repository são obrigatórios")
	}

	url := fmt.Sprintf("%s/repos/%s/%s", c.BaseURL, owner, repo)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("criando requisição: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("User-Agent", "github-repo-analyzer")

	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("consultando GitHub: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("lendo resposta: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("repositório não encontrado: %s/%s", owner, repo)
		}
		return nil, fmt.Errorf("GitHub API retornou HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var repository Repository
	if err := json.Unmarshal(body, &repository); err != nil {
		return nil, fmt.Errorf("interpretando JSON: %w", err)
	}

	return &repository, nil
}

func printRepository(repository *Repository) {
	description := repository.Description
	if description == "" {
		description = "Sem descrição."
	}

	language := repository.Language
	if language == "" {
		language = "Não informada"
	}

	fmt.Println()
	fmt.Println("GitHub Repo Analyzer")
	fmt.Println("====================")
	fmt.Printf("Repositório: %s\n", repository.FullName)
	fmt.Printf("Descrição:   %s\n", description)
	fmt.Printf("URL:         %s\n", repository.HTMLURL)
	fmt.Printf("Linguagem:   %s\n", language)
	fmt.Printf("Branch:      %s\n", repository.DefaultBranch)
	fmt.Printf("Stars:       %d\n", repository.Stars)
	fmt.Printf("Forks:       %d\n", repository.Forks)
	fmt.Printf("Issues:      %d\n", repository.OpenIssues)
	fmt.Printf("Watchers:    %d\n", repository.Watchers)
	fmt.Printf("Privado:     %t\n", repository.Private)
	fmt.Printf("Arquivado:   %t\n", repository.Archived)
	fmt.Printf("Criado em:   %s\n", repository.CreatedAt)
	fmt.Printf("Atualizado:  %s\n", repository.UpdatedAt)
	fmt.Println()
}

func printUsage() {
	fmt.Println("Uso:")
	fmt.Println("  go run . <owner>/<repository>")
	fmt.Println()
	fmt.Println("Exemplo:")
	fmt.Println("  go run . Coelho7z7/GoStock")
	fmt.Println()
	fmt.Println("Opcional:")
	fmt.Println("  GITHUB_TOKEN=<token> go run . <owner>/<repository>")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	parts := strings.Split(strings.Trim(os.Args[1], "/"), "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		fmt.Println("Erro: informe o repositório no formato owner/repository.")
		os.Exit(1)
	}

	client := NewGitHubClient()

	repository, err := client.GetRepository(parts[0], parts[1])
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		os.Exit(1)
	}

	printRepository(repository)
}

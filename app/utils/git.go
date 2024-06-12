package utils

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/goal-web/supports/utils"
	"os"
	"os/exec"
	"strings"
)

type RepoInfo struct {
	Commit   string
	Comment  string
	Branches []string
	Tags     []string
}

// CloneRepoBranchOrCommit 克隆指定分支或提交
func CloneRepoBranchOrCommit(repoURL, publicKey, branchOrCommit, destDir string) (RepoInfo, error) {
	var info RepoInfo
	auth, err := ssh.NewPublicKeys("git", []byte(publicKey), "")
	if err != nil {
		return info, fmt.Errorf("error creating ssh auth: %v", err)
	}

	_ = os.RemoveAll(destDir)

	_, err = git.PlainClone(destDir, false, &git.CloneOptions{
		URL:  repoURL,
		Auth: auth,
	})

	// Change to the specified directory
	if err = os.Chdir(destDir); err != nil {
		return info, fmt.Errorf("failed to change directory: %v", err)
	}

	// Run 'git checkout' command
	cmd := exec.Command("git", "checkout", branchOrCommit)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return info, fmt.Errorf("failed to execute 'git checkout': %v", err)
	}

	commit, comment, err := getCurrentCommitAndMessage(destDir)

	if err != nil {
		return info, fmt.Errorf("failed to get commit: %w", err)
	}

	info.Commit = commit
	info.Comment = comment

	branches, tags, err := getGitBranchesAndTags(destDir)

	if err != nil {
		return info, fmt.Errorf("failed to clone repository: %w", err)
	}

	info.Branches = branches
	info.Tags = tags

	return info, nil
}

// getGitBranchesAndTags returns the list of branches and tags for the given git directory
func getGitBranchesAndTags(gitDir string) ([]string, []string, error) {
	// Set the Git command
	cmd := exec.Command("git", "-C", gitDir, "branch")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get branches: %v", err)
	}
	branches := strings.Fields(strings.ReplaceAll(out.String(), "*", ""))

	// Set the Git command for tags
	cmd = exec.Command("git", "-C", gitDir, "tag")
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get tags: %v", err)
	}
	tags := strings.Fields(out.String())

	return branches, tags, nil
}

func CloneRepo(repoAddress, privateKey, version, targetDir string) (string, string, error) {
	// 创建临时文件保存私钥
	tempDir, err := os.MkdirTemp("", "pem"+utils.RandStr(10))

	if err != nil {
		return "", "", err
	}

	pem := fmt.Sprintf("%s/%s.pem", tempDir, utils.RandStr(10))

	// 将私钥字符串写入临时文件
	if err := os.WriteFile(pem, []byte(privateKey), 0600); err != nil {
		return "", "", fmt.Errorf("failed to write private key to temp file: %w", err)
	}
	defer os.Remove(pem) // 函数返回时删除临时文件

	// 设置 GIT_SSH_COMMAND 环境变量以使用临时文件中的私钥
	originalSSHCommand := os.Getenv("GIT_SSH_COMMAND")
	os.Setenv("GIT_SSH_COMMAND", fmt.Sprintf("ssh -i %s", pem))

	// 构建 git clone 命令
	cmd := exec.Command("git", "clone", repoAddress, targetDir)

	// 捕获命令的输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 运行命令
	err = cmd.Run()

	// 恢复原来的 GIT_SSH_COMMAND 环境变量
	os.Setenv("GIT_SSH_COMMAND", originalSSHCommand)

	if err != nil {
		return "", "", fmt.Errorf("failed to clone repository: %w", err)
	}

	// 如果 version 是一个特定的提交，切换到该提交
	cmd = exec.Command("git", "-C", targetDir, "checkout", version)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return "", "", fmt.Errorf("failed to checkout commit: %w", err)
	}

	commit, comment, err := getCurrentCommitAndMessage(targetDir)

	if err != nil {
		return "", "", fmt.Errorf("failed to clone repository: %w", err)
	}

	return commit, comment, nil
}

// getCurrentCommitAndMessage 获取当前的 commit hash 和 commit message
func getCurrentCommitAndMessage(repoDir string) (string, string, error) {
	// 获取当前 commit hash
	cmd := exec.Command("git", "-C", repoDir, "rev-parse", "HEAD")
	var out bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &out
	err := cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("failed to get current commit hash: %w", err)
	}
	commitHash := strings.TrimSpace(out.String())

	// 获取当前 commit message
	cmd = exec.Command("git", "-C", repoDir, "log", "-1", "--pretty=%B")
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", "", fmt.Errorf("failed to get current commit message: %w", err)
	}
	commitMessage := strings.TrimSpace(out.String())

	return commitHash, commitMessage, nil
}

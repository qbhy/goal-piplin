package utils

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/goal-web/supports/utils"
	"os"
	"os/exec"
	"strings"
)

// CloneRepoBranchOrCommit 克隆指定分支或提交
func CloneRepoBranchOrCommit(repoURL, publicKey, branchOrCommit, destDir string) (string, string, error) {
	auth, err := ssh.NewPublicKeys("git", []byte(publicKey), "")
	if err != nil {
		return "", "", fmt.Errorf("error creating ssh auth: %v", err)
	}

	_ = os.RemoveAll(destDir)

	_, err = git.PlainClone(destDir, false, &git.CloneOptions{
		URL:  repoURL,
		Auth: auth,
		//ReferenceName: reference,
	})

	// Change to the specified directory
	if err = os.Chdir(destDir); err != nil {
		return "", "", fmt.Errorf("failed to change directory: %v", err)
	}

	// Run 'git checkout' command
	cmd := exec.Command("git", "checkout", branchOrCommit)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return "", "", fmt.Errorf("failed to execute 'git checkout': %v", err)
	}

	commit, comment, err := getCurrentCommitAndMessage(destDir)

	if err != nil {
		return "", "", fmt.Errorf("failed to clone repository: %w", err)
	}

	return commit, comment, nil
}

// GetRepositoryBranchesAndTags 获取 Git 仓库的分支和 Tags
func GetRepositoryBranchesAndTags(repoURL string, publicKey string) ([]string, []string, error) {
	auth, err := ssh.NewPublicKeys("git", []byte(publicKey), "")
	if err != nil {
		return nil, nil, fmt.Errorf("error creating ssh auth: %v", err)
	}

	// 克隆仓库到内存中
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:  repoURL,
		Auth: auth,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("unable to clone repository: %w", err)
	}

	// 获取所有分支
	branches, err := repo.Branches()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to list branches: %w", err)
	}

	var branchList []string
	err = branches.ForEach(func(ref *plumbing.Reference) error {
		branchList = append(branchList, ref.Name().Short())
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	// 获取所有 Tags
	tags, err := repo.Tags()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to list tags: %w", err)
	}

	var tagList []string
	err = tags.ForEach(func(ref *plumbing.Reference) error {
		tagList = append(tagList, ref.Name().Short())
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return branchList, tagList, nil
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

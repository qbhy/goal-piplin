package utils

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path/filepath"
)

// ConnectSFTP 连接到SFTP服务器
func ConnectSFTP(serverAddress, username, privateKey string) (*sftp.Client, error) {
	// 解析私钥
	signer, err := ssh.ParsePrivateKey([]byte(privateKey))
	if err != nil {
		return nil, err
	}

	// 创建SSH客户端配置
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 注意：在实际应用中应该使用更安全的方式
	}

	// 建立SSH连接
	conn, err := ssh.Dial("tcp", serverAddress, config)
	if err != nil {
		return nil, err
	}

	// 创建SFTP客户端
	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// SyncDir 使用SFTP同步本地目录到远程目录
func SyncDir(client *sftp.Client, localDir, remoteDir string) error {
	// 递归遍历本地目录
	return filepath.Walk(localDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算远程路径
		relativePath, err := filepath.Rel(localDir, path)
		if err != nil {
			return err
		}
		remotePath := filepath.Join(remoteDir, relativePath)

		if info.IsDir() {
			// 创建远程目录
			err := client.MkdirAll(remotePath)
			if err != nil {
				return err
			}
		} else {
			// 同步文件
			localFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer localFile.Close()

			// 创建远程文件
			remoteFile, err := client.Create(remotePath)
			if err != nil {
				return err
			}
			defer remoteFile.Close()

			// 复制文件内容
			_, err = io.Copy(remoteFile, localFile)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

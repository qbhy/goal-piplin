package utils

import (
	"archive/zip"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path/filepath"
	"strings"
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

// SyncFile 使用SFTP同步本地文件到远程目录
func SyncFile(client *sftp.Client, localFilePath, remoteFilePath string) error {
	// Open the local file
	localFile, err := os.Open(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to open local file: %w", err)
	}
	defer localFile.Close()

	// Create the remote file
	remoteFile, err := client.Create(remoteFilePath)
	if err != nil {
		return fmt.Errorf("failed to create remote file: %w", err)
	}
	defer remoteFile.Close()

	// Read the content of the local file and write it to the remote file
	content, err := io.ReadAll(localFile)
	if err != nil {
		return fmt.Errorf("failed to read local file: %w", err)
	}

	if _, err := remoteFile.Write(content); err != nil {
		return fmt.Errorf("failed to write to remote file: %w", err)
	}

	return nil
}

// ZipFolder compresses the specified folder into a zip file, maintaining the folder's structure
func ZipFolder(source, target string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create the zip header based on the relative path to the source folder
		relPath := strings.TrimPrefix(path, filepath.Clean(source))
		relPath = strings.TrimPrefix(relPath, string(filepath.Separator)) // remove leading separator if exists

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = relPath
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

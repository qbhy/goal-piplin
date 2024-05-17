package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
)

// ConnectToSSHServer 使用SSH私钥连接到SSH服务器
func ConnectToSSHServer(serverAddress, publicKey, username string) (*ssh.Client, error) {
	// 创建SSH签名
	signer, err := ssh.ParsePrivateKey([]byte(publicKey))
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}

	// 创建SSH客户端配置
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		// 非生产环境可以使用这个选项来忽略公钥验证
		// 在生产环境中，应该提供HostKeyCallback来验证服务器的公钥
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接到SSH服务器
	client, err := ssh.Dial("tcp", serverAddress, config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to [%s]: %v", serverAddress, err)
	}

	return client, nil
}

// ExecuteSSHCommand 在SSH会话中执行命令并返回输出结果
func ExecuteSSHCommand(client *ssh.Client, commands ...string) (string, error) {
	if len(commands) == 0 || commands[0] == "" {
		return "", nil
	}
	// 创建一个新的SSH会话
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// 运行命令并获取其标准输出
	output, err := session.CombinedOutput(strings.Join(commands, "\n"))
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}

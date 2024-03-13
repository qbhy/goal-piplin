package tests

import (
	"fmt"
	"github.com/qbhy/goal-piplin/app/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	private = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAoZgbv2WY3QJe+v/oWQNtFSIPcevOiMZXEu3LZJmewUZfC0QS
oJbZ7gU83Lwx+v4mfFjwbDy7xbcZHtg8KFRrG52Gk9ds3w6fQEn7aPB0nuW77101
bxu1FrU8HHuKaQCQixVV0XvxdmukLjwYeqjBGeoujHxpgnmzUh7joX4gQCtl/rZW
pA7bcqDnTQ1Arxbdq3vVP12zbk7nZKVglCJOUAlLHHKzc7W9Z1PtTxfqBuDh9C2I
V/ofIOoLdzvbZi2xNKMuR7KxOIADlk3mX74u+uqAzMh5HkYQiJuLZaerVvTxDvAc
8snJ6tYhv3Q9VgwrugDiyD/UNHoFrnjtC3qd8wIDAQABAoIBAF4aa4ZBTwzddZxr
7M3xfdPBuwbXkLX78vc1/a0+/MGHDpBL+yED2DORX0kYW27UFGtzi8cscxkHsJxS
xm2iA6HYUWfFG0SmijzxGHSbGv7xwEj7mcNzejiYIEJh+098obAtI4XTzODufHTk
UEkV/yXtR2BOj08JOlFHZb5E6dkTv6j4XzMiiot1GHlXhxxc3YTyBpDYUVOLKAoH
Qelic8wyq+dkj+9mB6YOMXzrfzUdKsv4td95fH9hHXa5Z3FX0AootUUMaa5eDgRC
r3/aOCr8S9daA5ejjDkl3UUkPCvTzKcd3gIbz8EyYsrsz8No1jOc9KiQsC5G0s3E
HRH9AgECgYEAwXg5VFNufecIYOsZtEnSSP3NjlVQf8u4XJz+5TOAm747prNfnG1D
9HaLhT6IPIIpyyvFBSq+2b0pLxtk348qYaBEbeUVHUtyKoU5HWJpe9fnVYHE8Fkt
g01NpjWAZ90Hl3qcb9r4w82Lyw0INgXRGygBFBAAbYWfFbKN6QuxPvMCgYEA1dJ+
RNoi6AvWGiu55jZvTZI6qM0Dv9c6GxoUV2Eqck/dAJE8vGd4yZczzNPGqPySpz1C
jfKylNY1sKwm87fyogY4zZZvTZjY4ZcR4CXQVhRJLAiB/Q6lI7+GtPI3tyupZcSU
JKbK7W/Jo0Upoz8HMlINtD2pbMtDKwjrCyYe5QECgYBtPcI5QCbahnJvrzBDYY+Q
UWcY4Elk75X5DVjL+Gm9BwxNk2kAPZ4qUilzohxw9ho9M0i7IyjCb5HqnHA333HR
0BnzZ2+lq+0Z30GhuujO2dkwqeaWw/Pz+NlIaVtIykA4Iy5j5mOiw9QUYhZp0p0A
1XTObD6hmNp7+OcyWLzLSwKBgQCt66Avaus3qeEFuolkWuSfyRCTmuaw1WT7BHSF
OpCnGJTf0EMB0HwsJSPKOHv/minDhI2tHjrp228ifHTWisn9xmfPucxg5rGKlTHC
C5/xVGDMQ0NQTeg/MptkdRyijg4krAf/4/dtuB7gAfLDSRIWeS2SbRFxX8gLqh5d
HC3HAQKBgGErl8SHqXGnimDbTBHPIZNx9tMJsOs10AJbPnyikkXY8DmrLmufIJn4
jBdO95me98G0DdLw4H2RE+gkS30LuzncTnTyXBDhAyizybQkjt2bYeN/WdRJDtL3
lWKyORXyRSYsGAWgcAka+aG1xmhOF8ZiclNGV9cetJkOU+eYaeCK
-----END RSA PRIVATE KEY-----
`
	address  = "192.168.106.4:22"
	username = "root"
)

func TestSsh(t *testing.T) {

	client, err := utils.ConnectToSSHServer(address, private, username)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	output, err := utils.ExecuteSSHCommand(client, "echo 'goal piplin'")
	if err != nil {
		log.Fatalf("execute failed: %v", err)
	}

	fmt.Println(output)

	fmt.Println("Successfully connected to server")
	// 在这里执行你的SSH会话操作，例如运行命令、转发端口等
	// 确保在完成后关闭客户端连接
	defer client.Close()

}

func TestSyncFiles(t *testing.T) {
	localPath := "/Users/qbhy/project/go/goal-piplin/tests"
	remotePath := "/home/qbhy.linux/data"
	client, err := utils.ConnectSFTP(address, username, private)
	assert.NoError(t, err)
	assert.NoError(t, utils.SyncDir(client, localPath, remotePath))
}

package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func main() {
	// Параметры подключения
	server := "localhost:2022"
	username := "sftp_user"
	password := "spbec"

	// Настройка SSH-клиента
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // В реальных условиях это не рекомендуется
	}

	// Подключение к серверу
	conn, err := ssh.Dial("tcp", server, config)
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	defer conn.Close()

	// Создание SFTP-клиента
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Не удалось создать SFTP-клиент: %v", err)
	}
	defer client.Close()

	// Локальный и удалённый путь файла
	localFilePath := "local_file.json"
	remoteFilePath := "C:/Users/roxxxyie/World/GolangProjects/work/SPBECMINING/SFTP_Server/temp_sftp_dir/remote.json"

	// Открываем локальный файл для чтения
	localFile, err := os.Open(localFilePath)
	if err != nil {
		log.Fatalf("Не удалось открыть локальный файл: %v", err)
	}
	defer localFile.Close()

	// Открываем удалённый файл для записи
	remoteFile, err := client.Create(remoteFilePath)
	if err != nil {
		log.Fatalf("Не удалось создать удалённый файл: %v", err)
	}
	defer remoteFile.Close()

	// Копируем содержимое локального файла в удалённый файл
	bytesCopied, err := remoteFile.ReadFrom(localFile)
	if err != nil {
		log.Fatalf("Ошибка при копировании файла: %v", err)
	}

	fmt.Printf("Файл успешно отправлен. Скопировано %d байт.\n", bytesCopied)
}

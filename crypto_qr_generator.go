
package main

import (
	"fmt"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	// Адрес кошелька и название сети
	walletAddress := "0xYourWalletAddressHere" // Замените на ваш адрес
	network := "Ethereum"

	// Формирование данных для QR-кода
	data := fmt.Sprintf("ethereum:%s", walletAddress)

	// Генерация QR-кода
	fileName := "wallet_qr.png"
	err := qrcode.WriteFile(data, qrcode.Medium, 256, fileName)
	if err != nil {
		log.Fatalf("Ошибка создания QR-кода: %v", err)
	}

	fmt.Printf("QR-код успешно создан и сохранен в файл: %s
", fileName)
}

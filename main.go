package main

import (
	"bufio"
	"fmt"
	m "models"
	"net"
	"os"
)

func main() {
	// Écoute sur le port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erreur lors de l'écoute:", err)
		return
	}
	defer listener.Close()

	fmt.Println("En attente de connexions sur le port 8080...")

	// Accepte les connexions entrantes de manière asynchrone
	for {
		var client m.Client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation de la connexion:", err)
			continue
		}
		client.Conn = conn

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Entrez votre nom : ")
		name, err := reader.ReadString('\n')
		if err != nil || name == "" {
			fmt.Println("Erreur lors de l'entrée du nom !")
		}
		client.Name = name

		fmt.Println("Nouvelle connexion établie!")

		// Traitez la connexion dans une goroutine séparée
		go handleConnection(client)
	}
}

func handleConnection(client m.client) {
	// Code pour traiter la connexion
	defer client.Conn.Close()

	// Exemple: lire les données de la connexion
	buffer := make([]byte, 1024)
	for {
		_, err := client.Conn.Read(buffer)
		if err != nil {
			fmt.Println("Un client a quitté le chat !")
			return
		}

		fmt.Println("Données reçues:", string(buffer))
	}
}

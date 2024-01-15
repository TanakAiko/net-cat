package main

import (
	"fmt"
	"net"
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
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erreur lors de l'acceptation de la connexion:", err)
			continue
		}

		fmt.Println("Nouvelle connexion établie!")

		// Traitez la connexion dans une goroutine séparée
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Code pour traiter la connexion
	defer conn.Close()

	// Exemple: lire les données de la connexion
	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Un client a quitté le chat !")
			return
		}
	
		fmt.Println("Données reçues:", string(buffer))
	}
}

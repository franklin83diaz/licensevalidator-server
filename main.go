package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

func main() {
	configFile := "/etc/licvs/default.conf"
	if os.Getenv("ENV") == "debug" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		configFile = "default.conf"
	}

	// Add driver to load TOML files
	config.WithOptions(config.ParseEnv)
	config.AddDriver(toml.Driver)

	// Load config from files
	err := config.LoadFilesByFormat("toml", configFile)
	if err != nil {
		log.Fatal(err)
	}

	//Read
	httpsCertPath := config.String("https.cert_path")
	httpsKeyPath := config.String("https.key_path")
	fmt.Println("Ruta del certificado HTTPS:", httpsCertPath)
	fmt.Println("Ruta de la clave privada HTTPS:", httpsKeyPath)

	jwtPrivateKeyPath := config.String("jwt.private_key_path")
	jwtPublicKeyPath := config.String("jwt.public_key_path")
	jwtAlgorithm := config.String("jwt.algorithm")
	fmt.Println("Ruta de la clave privada JWT:", jwtPrivateKeyPath)
	fmt.Println("Ruta de la clave pública JWT:", jwtPublicKeyPath)
	fmt.Println("Algoritmo de firma JWT:", jwtAlgorithm)
}

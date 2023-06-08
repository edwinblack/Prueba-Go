package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Estructura del objecto esperado de la consulta
type JokerC struct {
	Categories string `json:"categories"`
	Created_at string `json:"created_at"`
	Icon_url   string `json:"icon_url"`
	Id         string `json:"id"`
	Updated_at string `json:"updated_at"`
	Url        string `json:"url"`
	Value      string `json:"value"`
}

func main() {
	// Array contendedor de los get consultados
	// Key sera un string para ir comparado si el ID existe con anterioridad o se agregara como uno nuevo y se almacenara el elemento objetido
	chuckJoker := make(map[string]JokerC)

	// Control de los jokes almacenados
	cantidadDeChucks := 0

	// Control de intentos fallidos para que no se cree un loop infinito en dado caso no se logre obtener la cantidad de elementos requeridos tenga un final
	intentosFallidos := 0
	// Loop para obtener los 25 jokers
	for cantidadDeChucks < 25 {
		// consulta a la restApi
		resp, err := http.Get("https://api.chucknorris.io/jokes/random")
		if err != nil {
			fmt.Println(err)
		}

		// Parceo del elemento obtenido
		jsonMap := JokerC{}
		erro := json.NewDecoder(resp.Body).Decode(&jsonMap)
		if err != nil {
			fmt.Println(erro)
		}

		// se comprueba si existe o no el elemento ya en el array
		// el if se encuentra en negación si este no se encuentra con el OK se agrega al elemento y se reinicia los intentos fallidos
		if _, ok := chuckJoker[jsonMap.Id]; !ok {
			chuckJoker[jsonMap.Id] = jsonMap
			cantidadDeChucks++
			intentosFallidos = 0
		} else {
			// de suponer que ya se fallo mas de 2 veces la consulta por un nuevo elemento este dará por cumplida la condicion de la obtencion de Jokes
			if intentosFallidos > 2 {
				cantidadDeChucks = 25
				fmt.Println("Se alcanzo el máximo de intentos fallidos")
			}
			intentosFallidos++
		}
	}

	// se pasa a imprimir la información obtenida con la cantidad total de Jokes y una peque;a descripción de cada uno de ellos
	fmt.Println("Chucks totales recolectados:", len(chuckJoker))
	noChuck := 1
	for _, chuck := range chuckJoker {
		fmt.Println("Chucks #", noChuck)
		fmt.Println("id:", chuck.Id)
		fmt.Println("icon_url:", chuck.Icon_url)
		fmt.Println("url:", chuck.Url)
		fmt.Println("value:", chuck.Value)
		fmt.Printf("\n")
		noChuck++
	}
}

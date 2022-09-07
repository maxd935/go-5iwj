package boardGame

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init_adress() string {
	godotenv.Load(".env")
	return os.Getenv("URL_SERVER") + ":" + os.Getenv("PORT_SERVER")
}

func addAdversaire() {
	fmt.Println(">>> addAdversaire")
	var URL_SERVER = init_adress()
	fmt.Println(URL_SERVER + "/add")

	// encodedJson, err := json.Marshal(car)

	// if err != nil {
	// 	panic(err)
	// }
	// resp, err := http.Post("http://localhost:8080/cars", "application/json", bytes.NewBuffer(encodedJson))

	// if err != nil {
	// 	panic(err)
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	// var returnedCar Car

	// err = json.Unmarshal(body, &returnedCar)

	// if err != nil {
	// 	panic(err)
	// }

	// return returnedCar
}

func etatPlateau() {
	fmt.Println(">>> etatPlateau")
	var URL_SERVER = init_adress()
	fmt.Println(URL_SERVER + "/board")
}

func nbrBateau() int {
	fmt.Println(">>> nbrBateau")
	var URL_SERVER = init_adress()
	fmt.Println(URL_SERVER + "/boats")
	resp, err := http.Get(URL_SERVER + "/boats")
	if err != nil {
		panic(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	type BoatResponse struct {
		NbrBateau int
	}
	var responseObject BoatResponse
	json.Unmarshal(responseData, &responseObject)
	fmt.Println("nbrBateau = ", responseObject.NbrBateau)
	return responseObject.NbrBateau
}

func launchAttack() {
	fmt.Println(">>> launchAttack")

}

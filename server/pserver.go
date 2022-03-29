package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

var jstr = `
{
    "status": 200,
    "list": [
        {
            "number": "ББ 1771706",
            "grz": "У348АВ48",
            "startdate": "10.01.2022 00:00",
            "validitydate": "14.01.2022 00:00",
            "allowedzona": "МКАД",
            "passstatus": "Выдан",
            "datecancellation": "",
            "typepassvalidityperiod": "Ночной"
        },
        {
            "number": "ББ 1774197",
            "grz": "У348АВ48",
            "startdate": "16.01.2022 00:00",
            "validitydate": "20.01.2022 00:00",
            "allowedzona": "МКАД",
            "passstatus": "Выдан",
            "datecancellation": "",
            "typepassvalidityperiod": "Ночной"
        },
        {
            "number": "ББ 1801078",
            "grz": "У348АВ48",
            "startdate": "14.03.2022 00:00",
            "validitydate": "18.03.2022 00:00",
            "allowedzona": "МКАД",
            "passstatus": "Выдан",
            "datecancellation": "",
            "typepassvalidityperiod": "Дневной"
        },
        {
            "number": "ББ 1805879",
            "grz": "У348АВ48",
            "startdate": "21.03.2022 00:00",
            "validitydate": "30.03.2022 00:00",
            "allowedzona": "МКАД",
            "passstatus": "Выдан",
            "datecancellation": "",
            "typepassvalidityperiod": "Дневной"
        }
    ],
    "inquiry": {
        "price": 0.6,
        "speed": 9,
        "attempts": 2
    }
}
`

type L struct {
	AllowedZona      string
	DatecanCellation string
	Grz              string
	Number           string
	PassStatus       string
	StartDate        string
	EndDate          string `json:"validitydate"`
	Type             string `json:"typepassvalidityperiod"`
}

type J struct {
	stat    int `json:"status"`
	List    []L
	Inquiry map[string]interface{}
}

func getData() {
	url := "https://api-cloud.ru/api/transportMos.php?type=pass&licenseSeries=%D0%91%D0%91&token=2159537639e5e8aa94fa83092785972b&regNumber=%D1%83348%D0%B0%D0%B248"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func Index(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(w, "Error!")
		}
	}()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var result J

	json.Unmarshal([]byte(jstr), &result)

	for _, v := range result.List {
		fmt.Println(v)
	}

	fmt.Println("Id :", result.stat)

	json.NewEncoder(w).Encode(result.List)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)

	router.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8088", router))

}

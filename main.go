package main

import (
	"net/http"
	//"log"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"encoding/json"
	"encoding/json"
	//"os"
)

type data struct {
	Id   	string
	Name 	string
	Symbol  string
	Rank    string
}


type Res struct {
	Success bool
	Data data
}

func handle(c echo.Context) error{
	resp, err := http.Get("https://api.coinmarketcap.com/v1/ticker/bitcoin/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jsonBlob = []byte(body)
	var d []data
	error := json.Unmarshal(jsonBlob, &d)
	if error != nil {
		fmt.Println("error:", error)
	}
	var r Res
	r.Success = true
	r.Data = d

	fmt.Println( r)

	return c.JSON(http.StatusOK, &r)
}

//func getEmail(c echo.Context) error {
//
//
//	res := emailRes{
//		Success: false,
//	}
//	b, err := json.Marshal(res)
//	if err != nil {
//		fmt.Println("error:", err)
//	}
//	fmt.Println(b)
//	os.Stdout.Write(b)
//
//
//	var jsonBlob = []byte(b)
//	var r emailRes
//	error := json.Unmarshal(jsonBlob, &r)
//	if error != nil {
//		fmt.Println("error:", error)
//	}
//	fmt.Println(r)
//
//	return c.JSON(http.StatusOK, &r)
//}

func main() {


	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.GET("/", handle)
	//e.POST("/sendEmail", getEmail)
	e.Logger.Fatal(e.Start(":8080"))
}

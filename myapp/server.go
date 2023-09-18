package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"myapp/calc"
	"net/http"
)

/*
func main() {
	//random_message := calc.Randmsg()
	//fmt.Println("random_message:", random_message)

	//address := calc.PublickeytoAddress([]byte("ad8frg5"))
	//fmt.Println("publickeytoaddress:", address)

	//hash := calc.HashMessage("f5rg5r7g")
	//fmt.Println("hashmessage:", hash)

	getaddress := calc.Getaddress("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	fmt.Println("get", getaddress)

	signedmsg := calc.SignMessage("f5rg5r7g", "8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	fmt.Println("signedmsg:", signedmsg)

	ve := calc.VerifySignature("f5rg5r7g", getaddress, signedmsg)
	fmt.Println("ve:", ve)
}
*/

func main() {

	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/get_message", func(c echo.Context) error {
		fmt.Println("Received random_message req...\n")

		//create a new session
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		//generate random message
		random_message := calc.Randmsg()
		fmt.Println("Server provide random_message:", random_message)

		//store message in a session
		sess.Values["message"] = random_message
		sess.Save(c.Request(), c.Response())

		//send message to client
		return c.JSON(http.StatusOK, map[string]interface{}{"message": random_message})
	})

	e.POST("/verify", func(c echo.Context) error {

		sess, _ := session.Get("session", c)

		//get wallet address and signed message from client
		address := c.FormValue("address")
		signedMessage := c.FormValue("signedMessage")

		//get known message from session
		message := sess.Values["message"].(string)

		//verify message
		result := calc.VerifySignature(message, address, signedMessage)

		//send verified result to client
		return c.JSON(http.StatusOK, map[string]interface{}{"verified": result})
	})

	//start sever
	e.Logger.Fatal(e.Start(":1323"))
}

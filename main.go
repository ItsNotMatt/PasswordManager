package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Login struct {
    Title string    `json:"title"`
    User string     `json:"username"`
    Pass string     `json:"password"`
}

var pass = "testpass"

func parseArgs() {
    if (os.Args[1] == "help") {
        fmt.Println("Commands: view {username} {password}")
    } else if (os.Args[1] == "view") {
        if len(os.Args) < 3 {
            log.Fatal("Expecting password to view")
        }
        pass := os.Args[2]
        if validatePass(pass) {
            fmt.Println("Password is correct")
            viewLogins()
        } else {
            fmt.Println("Password is incorrect")
        }
    } else if (os.Args[1] == "programpass") {
        pass := os.Args[2]
        writeHash(hash(pass))
    }
}

func hash(pass string) string {
    hash := sha256.New()
    hash.Write([]byte(pass))
    hashed := hash.Sum(nil)
    return hex.EncodeToString(hashed)
}

func writeHash(hash string) {
    err := os.WriteFile("hash.txt", []byte(hash), 0644)
    if err != nil {
        log.Fatal("Err: Unable to write to file", err)
    }
}

func validatePass(pass string) bool {
    hash := hash(pass)

    content, err := os.ReadFile("hash.txt")
    if err != nil {
        log.Fatal("Err: Unable to read file", err)
    }

    if string(content) != hash {
        return false
    } else {
        return true
    }
}

func viewLogins() {
    file, err := os.Open("logins.json")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	var logins []Login
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&logins); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	for _, login := range logins {
		fmt.Printf("%s, Username: %s, Password: %s\n", login.Title, login.User, login.Pass)
	}
}

func encrypt() {
}

func decrypt() {
}

func addLogin() {
}

func main() {

    logins := []Login{
		{Title: "Login for Amazon", User: "user1", Pass: "pass1"},
		{Title: "Login for Google", User: "user2", Pass: "pass2"},
	}

    content, err := json.MarshalIndent(logins, "", "    ")
    if err != nil {
        log.Fatal("Err trying to marshal login: ", err)
    }

    err = os.WriteFile("logins.json", content, 0644)
    if err != nil {
        log.Fatal("Unable to write to json file")
    }


    parseArgs()
}







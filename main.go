//usr/bin/go run $1 $@ ; exit
// That's a special She-bang for go

// This is a demo rocketbot in golang
// Its purpose is to showcase some features

// Specify we are the main package (the one that contains the main function)
package main

import (
    // Import from the current directory the folder rocket and call the package rocket
    "./rocket"
    "fmt"
    "strings"
    //"time"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "math/rand"
)

func main() {
    rock, err := rocket.NewConnectionConfig("rb.cfg")
    rock.UserTemporaryStatus(rocket.STATUS_ONLINE)

    // If there was an error connecting, panic
    if err != nil {
        panic(err)
    }

    for {
        // Wait for a new message to come in
        msg, err := rock.GetNewMessage()

        // If error, quit because that means the connection probably quit
        if err != nil {
            break
        }

        // Print the message structure in a user-legible format
        // yml is []byte type, _ means send the returned error to void
        yml, _ := yaml.Marshal(msg)
        fmt.Println(string(yml))

        // If begins with '@Username ' or is in private chat
        if msg.IsAddressedToMe || msg.RoomName == "" || msg.IsDirect {
            if (strings.Contains(strings.ToLower(msg.GetNotAddressedText()), "i am")) {
                name := strings.TrimPrefix(strings.ToLower(msg.GetNotAddressedText()), "i am")
                msg.Reply(fmt.Sprintf("Hello %s, I'm Dad.", name))
            }
            if(strings.Contains(strings.ToLower(msg.Text), "tell me a joke")) {
                reply := joke()
                msg.Reply(reply)
            }
        }
    }
}

func joke() string {
    jokeFile, err := ioutil.ReadFile("jokes.txt")

    if err != nil {
        return "No Joke found"
    }

    jokes := string(jokeFile)
    jokeArray := strings.Split(jokes, "\n")

    num := rand.Intn(len(jokeArray) - 1)

    return jokeArray[num]
}

package main

import "fmt"

type User struct {
	Id int
	Name, Location string
}

type Player struct {
	User // User will contain all the required attributes
	GameId int
}

func (player *Player) showPlayerInfo () string {
	fmt.Println("METHOD=showPlayerInfo")
	fmt.Println(player.Id)
	fmt.Println(player.Name)
	fmt.Println(player.Location)
	fmt.Println(player.GameId)
	return player.Name
}

type Message struct {
	text string
}



func main() {
	p := Player{
		User{Id: 42, Name: "Matt", Location: "LA"},
		90404,
	}
	fmt.Println(p)

	fmt.Println("Initializing var msg")
	var msg Message
	fmt.Println("msg variable created")

	msg = Message{text:"Hello World from Message struct"}
	fmt.Println(msg.text)
	p.showPlayerInfo()
}	

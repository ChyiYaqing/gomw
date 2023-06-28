package main

import "fmt"

/*
the Memento Pattern
	> The "Memento Pattern" is a powerful pattern that can be used to implement a load and save system for games.

*/

type Originator interface {
	Save() Memento
	Restore(Memento)
}

type GameWorld struct {
	level int
	score int
}

func (g *GameWorld) Save() Memento {
	return Memento{level: g.level, score: g.score}
}

func (g *GameWorld) Restore(m Memento) {
	g.level = m.level
	g.score = m.score
}

type Memento struct {
	level int
	score int
}

type SaveFileManager struct {
	mementos []Memento
}

func (s *SaveFileManager) Save(m Memento) {
	s.mementos = append(s.mementos, m)
}

func (s *SaveFileManager) Load(index int) Memento {
	return s.mementos[index]
}

func main() {
	gameWorld := &GameWorld{level: 1, score: 0}
	SaveFileManager := &SaveFileManager{}
	// Save the initial state of the game
	SaveFileManager.Save(gameWorld.Save())
	// Increase the level and score
	gameWorld.level++
	gameWorld.score += 1000
	// Save the new state of the game
	SaveFileManager.Save(gameWorld.Save())
	// Restore the previous state of the game
	gameWorld.Restore(SaveFileManager.Load(0))
	fmt.Println(gameWorld)
}

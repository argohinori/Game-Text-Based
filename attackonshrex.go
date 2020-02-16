package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

const N = 9999

type player struct {
	lvl    int
	name   string
	class  string
	health int
	damage int
	status bool
	xp     int
	maxHP  int
}
type bot struct {
	name   string
	lvl    int
	health int
	damage int
	status bool
}

type player_character [N]player
type bot_monster [10]bot

var player_character_hero player_character
var bot_monster_enemy bot_monster

/*
func npc(bot_monster_enemy bot_monster, i int) bot_monster {

	return bot_monster_enemy
}
*/

func character_lvlup(pch player_character) player_character {
	if player_character_hero[0].xp >= 100 {
		player_character_hero[0].lvl = player_character_hero[0].lvl + 1
		player_character_hero[0].maxHP = player_character_hero[0].maxHP + 20
		player_character_hero[0].damage = player_character_hero[0].damage + 4
		player_character_hero[0].xp = player_character_hero[0].xp - 100
		player_character_hero[0].health = player_character_hero[0].health + 15
	}
	return player_character_hero
}
func character_alive(pch player_character) player_character { //Status hidup player

	if player_character_hero[0].health <= 0 {
		player_character_hero[0].health = 0
		player_character_hero[0].status = false
		player_character_hero[0].name = "(Dead)"
	} else {
		player_character_hero[0].status = true
	}
	return player_character_hero
}
func monster_alive(bme bot_monster, i int) bot_monster { //Status hidup bot

	if bot_monster_enemy[i].health <= 0 {
		bot_monster_enemy[i].health = 0
		bot_monster_enemy[i].status = false
		bot_monster_enemy[i].name = bot_monster_enemy[i].name + "(Dead)"
	} else {
		bot_monster_enemy[i].status = true
	}
	return bot_monster_enemy
}
func insertName(pch player_character) player_character {
	var jawaban string

	fmt.Print("\nInsert the player's name: ")
	fmt.Scanln(&player_character_hero[0].name)
	fmt.Print("Player's name:", player_character_hero[0].name, " Are you sure? Y/N: ")
	fmt.Scanln(&jawaban)
	for jawaban == "N" || jawaban == "n" {
		fmt.Print("\nInsert the player's name: ")
		fmt.Scanln(&player_character_hero[0].name)
		fmt.Print("Player's name:", player_character_hero[0].name, " Are you sure? Y/N: ")
		fmt.Scanln(&jawaban)
	}
	return player_character_hero
}
func player_board(pch player_character) player_character { // Menampilkan informasi player
	fmt.Println("\nHero\n======")
	fmt.Println("Name: ", player_character_hero[0].name)
	fmt.Println("Level: ", player_character_hero[0].lvl)
	fmt.Println("Class: ", player_character_hero[0].class)
	fmt.Println("HP: ", player_character_hero[0].health, "/", player_character_hero[0].maxHP)
	fmt.Println("Damage: ", player_character_hero[0].damage)
	fmt.Println("Status: ", player_character_hero[0].status)
	fmt.Println("XP: ", player_character_hero[0].xp)

	return player_character_hero
}
func bot_board(bme bot_monster, i int) bot_monster { // Menampilkan informasi bot
	fmt.Println("\nEnemy\n======")
	fmt.Println("Name: ", bot_monster_enemy[i].name)
	fmt.Println("Level: ", bot_monster_enemy[i].lvl)
	fmt.Println("HP: ", bot_monster_enemy[i].health)
	fmt.Println("Damage: ", bot_monster_enemy[i].damage)
	fmt.Println("Status: ", bot_monster_enemy[i].status)

	return bot_monster_enemy
}

func gameplay(pch player_character, i int) player_character { //serang-menyerang
	var action int

	fmt.Println("==================================")
	fmt.Println("\nAttack[1]\nRest[2]")
	if player_character_hero[0].status != false {
		fmt.Print("\n", player_character_hero[0].name, ", What do you want to do?: ")
		fmt.Scanln(&action)
		if action == 1 {
			bot_monster_enemy[i].health = bot_monster_enemy[i].health - player_character_hero[0].damage
			player_character_hero[0].xp = player_character_hero[0].xp + 25
		} else if action == 2 {
			player_character_hero[0].health = player_character_hero[0].health + 20
			player_character_hero[0].xp = player_character_hero[0].xp + 15
			if player_character_hero[0].health >= player_character_hero[0].maxHP {
				player_character_hero[0].health = player_character_hero[0].maxHP
			}
		}
	}
	return player_character_hero
}

func bot_attack(pch player_character, bme bot_monster, i int) player_character {
	player_character_hero[0].health = player_character_hero[0].health - bot_monster_enemy[i].damage

	return player_character_hero
}

func bubbleSort_level(bot_monster_enemy *bot_monster) {
	for i := 10; i > 0; i-- {
		for j := 1; j < i; j++ {
			if bot_monster_enemy[j-1].lvl > bot_monster_enemy[j].lvl {
				temp := bot_monster_enemy[j]
				bot_monster_enemy[j] = bot_monster_enemy[j-1]
				bot_monster_enemy[j-1] = temp
			}
		}
	}
}

func bubbleSort_hp(bot_monster_enemy *bot_monster) {
	for i := 10; i > 0; i-- {
		for j := 1; j < i; j++ {
			if bot_monster_enemy[j-1].health > bot_monster_enemy[j].health {
				temp := bot_monster_enemy[j]
				bot_monster_enemy[j] = bot_monster_enemy[j-1]
				bot_monster_enemy[j-1] = temp
			}
		}
	}
}

func bubbleSort_damage(bot_monster_enemy *bot_monster) {
	for i := 10; i > 0; i-- {
		for j := 1; j < i; j++ {
			if bot_monster_enemy[j-1].damage > bot_monster_enemy[j].damage {
				temp := bot_monster_enemy[j]
				bot_monster_enemy[j] = bot_monster_enemy[j-1]
				bot_monster_enemy[j-1] = temp
			}
		}
	}
}

func main() { //main
	var i, j, enter, choice, choice_class, choice_list int
	var monster_name string

	i = 0
	player_character_hero[0].lvl = 1
	player_character_hero[0].maxHP = 100
	player_character_hero[0].status = true
	player_character_hero[0].health = player_character_hero[0].maxHP
	bot_monster_enemy = [10]bot{{"Wolf", 3, 55, 15, true}, {"Undead", 7, 75, 23, true},
		{"Spirit", 10, 55, 30, true}, {"Troll", 9, 100, 21, true}, {"Werewolf", 13, 70, 33, true},
		{"Slime", 15, 115, 30, true}, {"Wisp", 21, 70, 39, true}, {"Vampire", 25, 130, 37, true},
		{"Giant", 31, 120, 45, true}, {"Mr.Shrex", 99, 225, 40, true}}

	fmt.Println("\nWelcome to [Attack on Shrex] RPG Game")
	fmt.Println("\nEnter[1]\nList[2]")
	fmt.Print("\nPlease Insert your answer: ")
	fmt.Scanln(&choice)
	for choice == 1 || choice == 2  {
		if choice == 1 && player_character_hero[0].status != false && i < 10 {
			b, err := ioutil.ReadFile("Intro.txt")
	
			if err != nil {
			fmt.Println(err)
			}
			str := string(b)
			fmt.Println(str)
			
			player_character_hero = insertName(player_character_hero)
			fmt.Print("\n")

			fmt.Println("\nTipe Class:\nWarrior[1]\nArchery[2]\nAssasin[3]")

			fmt.Print("\nChoose your class: ")
			fmt.Scanln(&choice_class)
			switch choice_class {
			case 1:
				player_character_hero[0].damage = 35
				player_character_hero[0].class = "Warrior"
			case 2:
				player_character_hero[0].damage = 22
				player_character_hero[0].class = "Archery"
			case 3:
				player_character_hero[0].damage = 18
				player_character_hero[0].class = "Assasin"
			default:
				player_character_hero[0].damage = 31
				player_character_hero[0].class = "Goblin"
			}

			for player_character_hero[0].status != false && i < 10 {
				player_character_hero[0].health = player_character_hero[0].maxHP
				player_character_hero = character_alive(player_character_hero)
				bot_monster_enemy = monster_alive(bot_monster_enemy, i)

				player_character_hero = player_board(player_character_hero)
				bot_monster_enemy = bot_board(bot_monster_enemy, i)

				for bot_monster_enemy[i].status != false && player_character_hero[0].status != false {
					player_character_hero = gameplay(player_character_hero, i)
					player_character_hero = character_lvlup(player_character_hero)
					player_character_hero = bot_attack(player_character_hero, bot_monster_enemy, i)
					player_character_hero = character_alive(player_character_hero)
					bot_monster_enemy = monster_alive(bot_monster_enemy, i)
					player_character_hero = player_board(player_character_hero)
					bot_monster_enemy = bot_board(bot_monster_enemy, i)

				}
				i++
				fmt.Print("\nPress Enter to Continue\nChange name[9]\n")
				fmt.Scanln(&enter)
				if enter == 9 {
					player_character_hero = insertName(player_character_hero)
					fmt.Print("\n")
				} else if player_character_hero[0].status == false {
					fmt.Println("GGWP, Mr.Shrex's monster killed you!")
					fmt.Println("You succeed until stage", i)

					file, lose := os.Create("Tips_from_Devs.txt")
					if lose != nil {
						log.Fatal("Error", lose)
					}

					defer file.Close()
					fmt.Fprintf(file, "You lose! Try resting every now and then, you could accumulate XP and level up!")

				} else if i >= 10 {
					fmt.Print("\n==================================\n")
					fmt.Println("Congratulations, you have defeated Mr.Shrex.")
					fmt.Println("The world is peaceful again, for now....")
					fmt.Print("\n==================================")
				}

			}

		} else if choice == 2 {
			fmt.Print("\nWhat do you want to see? ")
			fmt.Println("\nMonster's level[1] \nMonster's Damage[2] \nMonster's HP[3] \nSpecific Monster[4]")
			fmt.Scanln(&choice_list)
			if choice_list == 1 {
				fmt.Println("\nMonster's list based on level:")
				fmt.Print("\n==================================\n")
				bubbleSort_level(&bot_monster_enemy)
				for j = 0; j < 10; j++ {
					fmt.Println(bot_monster_enemy[j].lvl, bot_monster_enemy[j].name)
				}
				fmt.Print("\nWhat do you want to do?")
				fmt.Print("\nEnter the game[1]\nLook up something else[2]\n")
				fmt.Scanln(&choice)
				fmt.Print("\n==================================")
				fmt.Print("\n")
				bot_monster_enemy = [10]bot{{"Wolf", 3, 55, 15, true}, {"Undead", 7, 75, 23, true},
					{"Spirit", 10, 55, 30, true}, {"Troll", 9, 100, 21, true}, {"Werewolf", 13, 70, 33, true},
					{"Slime", 15, 115, 30, true}, {"Wisp", 21, 70, 39, true}, {"Vampire", 25, 130, 37, true},
					{"Giant", 31, 120, 45, true}, {"Mr.Shrex", 99, 225, 40, true}}

			} else if choice_list == 2 {
				fmt.Println("\nMonster's list based on damage:")
				fmt.Print("\n==================================\n")
				bubbleSort_damage(&bot_monster_enemy)
				for j = 0; j < 10; j++ {
					fmt.Println(bot_monster_enemy[j].damage, bot_monster_enemy[j].name)
				}
				fmt.Print("\nWhat do you want to do?")
				fmt.Print("\nEnter the game[1]\nLook up something else[2]\n")
				fmt.Scanln(&choice)
				fmt.Print("\n==================================")
				fmt.Print("\n")
				bot_monster_enemy = [10]bot{{"Wolf", 3, 55, 15, true}, {"Undead", 7, 75, 23, true},
					{"Spirit", 10, 55, 30, true}, {"Troll", 9, 100, 21, true}, {"Werewolf", 13, 70, 33, true},
					{"Slime", 15, 115, 30, true}, {"Wisp", 21, 70, 39, true}, {"Vampire", 25, 130, 37, true},
					{"Giant", 31, 120, 45, true}, {"Mr.Shrex", 99, 225, 40, true}}

			} else if choice_list == 3 {
				
				fmt.Println("\nMonster's list based on Health:")
				fmt.Print("\n==================================\n")
				bubbleSort_hp(&bot_monster_enemy)
				for j = 0; j < 10; j++ {
					fmt.Println(bot_monster_enemy[j].health, bot_monster_enemy[j].name)
				}
				fmt.Print("\nWhat do you want to do?")
				fmt.Print("\nEnter the game[1]\nLook up something else[2]\n")
				fmt.Scanln(&choice)
				fmt.Print("\n==================================")
				fmt.Print("\n")
				bot_monster_enemy = [10]bot{{"Wolf", 3, 55, 15, true}, {"Undead", 7, 75, 23, true},
					{"Spirit", 10, 55, 30, true}, {"Troll", 9, 100, 21, true}, {"Werewolf", 13, 70, 33, true},
					{"Slime", 15, 115, 30, true}, {"Wisp", 21, 70, 39, true}, {"Vampire", 25, 130, 37, true},
					{"Giant", 31, 120, 45, true}, {"Mr.Shrex", 99, 225, 40, true}}

			} else if choice_list == 4 {
				ketemu := false
				j = 0
				fmt.Println("\nList of all the monsters: ")
				fmt.Print("\n==================================\n")
				for j < 10 {
					fmt.Println(bot_monster_enemy[j].name)
					j++
					}
				j = 0
				fmt.Println("\nEnter the monster's name:")
				fmt.Scanln(&monster_name)
				for j < 10 && !ketemu {
					if monster_name == bot_monster_enemy[j].name {
						ketemu = true
						j--
					}
					j++
				}
				fmt.Print("\nEnemy\n======\n", bot_monster_enemy[j].name, "\nStage ", j+1, "\nLevel ", bot_monster_enemy[j].lvl, "\nHP ", bot_monster_enemy[j].health, "\nDamage ", bot_monster_enemy[j].damage)
				fmt.Println("\nWhat do you want to do? ")
				fmt.Print("\nEnter the game[1]\nLook up something else[2]\n")
				fmt.Scanln(&choice)
				fmt.Print("\n==================================")
				fmt.Print("\n")
			}
		}
	}
}
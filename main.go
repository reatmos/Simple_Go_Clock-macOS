package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"os/exec"
	"time"
)

var inc int
var ah int
var am int
var as int
var stc int

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func key() {
	keysEvents, _ := keyboard.GetKeys(10)
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		if keysEvents != nil {
			event := <-keysEvents
			if event.Key == keyboard.KeyTab {
				inc = 1
				break
			}
		}
	}
}

func clock() {
	inc = 0
	fmt.Println("Clock\n--------------------\nPress Tab to Menu\n--------------------")
	for {
		tn := time.Now()
		if ah != 0 || am != 0 || as != 0 {
			fmt.Printf("\rTime : %02d:%02d:%02d - Alarm : %02d:%02d:%02d", tn.Hour(), tn.Minute(), tn.Second(), ah, am, as)
			time.Sleep(time.Millisecond * 100)
		} else {
			fmt.Printf("\rTime : %02d:%02d:%02d", tn.Hour(), tn.Minute(), tn.Second())
			time.Sleep(time.Millisecond * 100)
		}

		if inc == 1 {
			break
		}

		if tn.Hour() == ah && tn.Minute() == am && tn.Second() == as {
			clear()
			cmd := exec.Command("osascript", "-e", `tell application "Terminal" to activate do script "echo ALARMMMMMMMMMMMMMMM\nexit"`)
			cmd.Stdout = os.Stdout
			cmd.Run()

			time.Sleep(time.Second * 3)
			clear()
			ah = 0
			am = 0
			as = 0
			fmt.Println("Clock\n--------------------\nPress Tab to Menu\n--------------------")
		}
	}
}

func stopw() {
	inc = 0
	for {
		if stc == 1 {
			min := 0
			for {
				for sec := 0; sec < 60; sec++ {
					fmt.Printf("\rStopwatch : %02d:%02d", min, sec)
					time.Sleep(time.Second * 1)
					if inc == 1 {
						break
					}
				}
				min++
				if inc == 1 {
					break
				}
			}
		}
		break
	}
}

func main() {
	clear()
	go clock()
	key()
	for {
		var sel int
		clear()
		fmt.Print("Menu\n--------------------\n1. Clock 2. Alarm 3. Stopwatch 4. Timer 5. Exit\nEnter>")
		fmt.Scan(&sel)
		if sel == 1 {
			clear()
			go clock()
			key()
		} else if sel == 2 {
			clear()
			var alm int
			for {
				clear()
				fmt.Print("Alarm\n--------------------\n1. Set to Alarm 2. Clear Alarm 3. Exit\nEnter>")
				fmt.Scan(&alm)
				if alm == 1 {
					clear()
					for {
						clear()
						fmt.Print("Alarm\n--------------------\nSet to Alarm for 24-Hours Format(ex : 13:00:00)\nEnter>")
						fmt.Scanf("%02d:%02d:%02d", &ah, &am, &as)
						if ah < 25 && am < 60 && as < 60 {
							fmt.Printf("Set the Alarm : %02d:%02d:%02d", ah, am, as)
							time.Sleep(time.Second * 2)
							clear()
							go clock()
							key()
							break
						} else {
							fmt.Print("Try again")
							time.Sleep(time.Second * 2)
						}
					}
					break
				} else if alm == 2 {
					ah = 0
					am = 0
					as = 0
					fmt.Print("Clear Alarm")
					time.Sleep(time.Second * 2)
					clear()
					go clock()
					key()
					break
				} else if alm == 3 {
					break
				}
			}
		} else if sel == 3 {
			var sw int
			for {
				clear()
				fmt.Print("Stopwatch\n--------------------\n1. Start 2. Exit \nEnter>")
				fmt.Scan(&sw)
				if sw == 1 {
					stc = 1
					clear()
					fmt.Print("Stopwatch\n--------------------\nTab to Stop\n--------------------\n")
					go stopw()
					key()
					fmt.Print("\n")
					for i := 5; i > 0; i-- {
						fmt.Printf("\rExit.. %d", i)
						time.Sleep(time.Second * 1)
					}
					break
				} else if sw == 2 {
					break
				}
			}
		} else if sel == 4 {
			tim := 0
			for {
				clear()
				fmt.Print("Timer\n--------------------\nEnter second for timer (ex : 1min = 60sec)\nEnter>")
				fmt.Scan(&tim)
				for i := tim; i > 0; i-- {
					clear()
					fmt.Printf("Timer\n--------------------\n\rTimer : %d", i)
					time.Sleep(time.Second * 1)
				}
				clear()
				cmd := exec.Command("osascript", "-e", `tell application "Terminal" to activate do script "echo TIMERRRRRRRRRRRRRR\nexit"`)
				cmd.Stdout = os.Stdout
				cmd.Run()

				cmd = exec.Command("exit")
				cmd.Stdout = os.Stdout
				cmd.Run()

				time.Sleep(time.Second * 3)
				clear()
				go clock()
				key()
				break
			}
		} else if sel == 5 {
			clear()
			fmt.Print("Goodbye")
			break
		}
	}
}

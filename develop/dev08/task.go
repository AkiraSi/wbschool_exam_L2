package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

func cd(way string) {
	err := os.Chdir(way)
	if err != nil {
		fmt.Println(err)
	}
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func ps() {
	cmd := exec.Command("tasklist")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func kill(pid int) {
	err := syscall.TerminateProcess(syscall.PROCESS_TERMINATE, uint32(pid))
	if err != nil {
		fmt.Println(err)
	}
}

// func connect(ip string, port int) net.Conn {
// 	connectStr := fmt.Sprintf("%s:%d", ip, port)
// 	conn, err := net.Dial("tcp", connectStr)
// 	if err != nil {
// 		fmt.Printf("couldn't connect to %s...\n", connectStr)
// 	}
// 	return conn
// }

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Fprint(os.Stdout, "Выход из оболочки")
		os.Exit(0)
	}()
	reader := bufio.NewReader(os.Stdin) // ридер консоли
	for {
		fmt.Print("$ ") // для красоты
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		commands := strings.Split(input, "|") // Разделение на команды по "|"
		for i := 0; i < len(commands); i++ {
			cmd := strings.TrimSpace(commands[i])
			if cmd == "" {
				continue
			}
			parts := strings.Split(cmd, " ")
			switch parts[0] { // по первому куску из слайса
			case "pwd":
				pwd()
			case "ps":
				ps()
			case "kill":
				if len(parts) != 2 {
					fmt.Println("Аргумент должен быть 1")
					break
				}
				pid, _ := strconv.Atoi(parts[1])
				kill(pid)
			case "echo":
				if len(parts) < 2 {
					fmt.Println("Недостаточно аргументов")
					break
				}
				echo(parts[1:])
			case "cd":
				if len(parts) != 2 {
					fmt.Println("Аргумент должен быть 1")
					break
				}
				cd(parts[1])
			default:
				if i == 0 {
					RunCommand(parts, nil)
				} else {
					RunCommand(parts, os.Stdout)
				}
			}
		}
	}
}

func RunCommand(parts []string, out io.Writer) { // запускальщик команд
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

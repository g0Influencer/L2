package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Telnet struct{ // структура с полученными данными
	host string
	port string
	timeout time.Duration
}

func connect() *Telnet{
	var timeout time.Duration
	flag.DurationVar(&timeout,"timeout",time.Second*10,"timeout")

	flag.Parse() // парсим флаги
	args:=flag.Args()
	if len(args) < 2 {
		panic("invalid params")
	}
	a:= &Telnet{
		host: args[0],
		port: args[1],
		timeout: timeout,
	}
	return a

}


func send(conn net.Conn, sigChan chan os.Signal, errChan chan error ){
	for {
		reader := bufio.NewReader(os.Stdin) // читаем из консоли
		text,err:=reader.ReadString('\n') // читаем до разделителя
		if err!= nil || err != io.EOF{ // если файл для чтения кончился
			sigChan<-syscall.SIGQUIT
			return
		}
		errChan <-err

		fmt.Fprintf(conn, text+"\n") // отправляем в сокет
	}
}

func read(conn net.Conn, sigChan chan os.Signal, errChan chan error){

	reader := bufio.NewReader(conn) // читаем из канала
	text,err:=reader.ReadString('\n') // читаем до разделителя
		if err!= nil || err != io.EOF{ // если файл для чтения кончился
			sigChan<-syscall.SIGQUIT
			return
		}
		errChan<-err

		fmt.Print("Message Received:", text) // выводим полученое сообщение в консоль
}

func main(){
	a:= connect()
	sigChan:=make(chan os.Signal,1) // канал для отправки сигналов о завершении работы клиента
	errChan:=make(chan error, 1)
	signal.Notify(sigChan,syscall.SIGINT,syscall.SIGTERM) // отлавливаем сигналы о завершении работы
	address:=net.JoinHostPort(a.host,a.port) // объединяем хост и порт
	conn,err:= net.DialTimeout("tcp",address,a.timeout) // подключаемся к сети
	if err!=nil{
		log.Fatal(err)
	}

	go send(conn,sigChan,errChan)
	go read(conn,sigChan,errChan)

	select{
	case <-sigChan:
		fmt.Println("Telnet client is close")
		conn.Close()
	case err:= <-errChan:
		log.Println(err)
		return
	}

}
package main

import (
	"fmt"
	controllers "modulo/controllers"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("Microserviço iniciado corretamente")
	c := cron.New()
	//_, err := c.AddFunc("*/3 * * * *", controllers.TarefaDiaria) // Agenda para rodar a cada 3 minutos
	_, err := c.AddFunc("0 0 9 * * *", controllers.TarefaDiaria) // Agenda para rodar diariamente às 9:00 AM
	if err != nil {
		fmt.Println("Erro ao agendar tarefa:", err)
		return
	}
	c.Start()

	select {}
}

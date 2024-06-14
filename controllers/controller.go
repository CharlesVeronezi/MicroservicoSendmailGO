package controllers

import (
	"fmt"
	datasearch "modulo/services/DataSearch"
	sendemail "modulo/services/SendEmail"
	"sync"
)

type Acoes struct {
	Ativo         string  `json:"ativo"`
	PrecoLucro    float32 `json:"precoLucro"`
	ValorAtual    float32 `json:"valorAtual"`
	VariacaoDoDia float32 `json:"variacaoDiaria"`
	MinimoAnual   float32 `json:"minimoAno"`
	MaximoAnual   float32 `json:"maximoAno"`
}

func TarefaDiaria() {
	fmt.Println("Iniciando tarefa...")
	acoesList := []string{
		"CMIG4", "CPLE6", "PETR4", "VALE3",
		"EGIE3", "CPLE3", "TRPL4", "TAEE11",
		"CPFE3", "ALUP11", "AURE3", "BBAS3",
		"SANB11", "ITSA4", "BRSR6", "ABCB4",
		"BBDC3", "BBDC4", "BBSE3", "CXSE3",
		"PSSA3", "KLBN4", "KLBN11", "VIVT3",
		"TIMS3", "CSMG3", "SAPR11", "SAPR4",
		"VALE3", "VBBR3",
	} // Adicione as ações que você quer buscar aqui
	var conteudo string

	var wg sync.WaitGroup
	ch := make(chan datasearch.Acoes, len(acoesList))

	for _, acao := range acoesList {
		wg.Add(1)
		go datasearch.BuscarDadosApi(acao, &wg, ch)
	}

	wg.Wait()
	close(ch)

	var todasAcoes []datasearch.Acoes
	for acao := range ch {
		todasAcoes = append(todasAcoes, acao)
	}

	for _, acao := range todasAcoes {
		conteudo += fmt.Sprintf("Ativo: %s\nPreço/Lucro: %.2f\nValor Atual: %.2f\nVariação do dia: %.2f%%\nMínimo Anual: %.2f\nMáximo Anual: %.2f\n\n",
			acao.Ativo, acao.PrecoLucro, acao.ValorAtual, acao.VariacaoDoDia, acao.MinimoAnual, acao.MaximoAnual)
	}

	if err := sendemail.EnviarEmail(conteudo); err != nil {
		fmt.Println("Erro ao enviar email:", err)
	}
}

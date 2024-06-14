package datasearch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Objeto de retorno da api
type Retorno struct {
	Results     []Results `json:"results"`
	RequestedAt string    `json:"requestedAt"`
	Took        string    `json:"took"`
}

// Objeto que retorna dentro do retorno do endpoint
type Results struct {
	Currency                   string  `json:"currency"`
	ShortName                  string  `json:"shortName"`
	LongName                   string  `json:"longName"`
	RegularMarketChange        float32 `json:"regularMarketChange"`
	RegularMarketChangePercent float32 `json:"regularMarketChangePercent"`
	RegularMarketTime          string  `json:"regularMarketTime"`
	RegularMarketPrice         float32 `json:"regularMarketPrice"`
	RegularMarketDayHigh       float32 `json:"regularMarketDayHigh"`
	RegularMarketDayRange      string  `json:"regularMarketDayRange"`
	RegularMarketDayLow        float32 `json:"regularMarketDayLow"`
	RegularMarketVolume        int     `json:"regularMarketVolume"`
	RegularMarketPreviousClose float32 `json:"regularMarketPreviousClose"`
	RegularMarketOpen          float32 `json:"regularMarketOpen"`
	FiftyTwoWeekRange          string  `json:"fiftyTwoWeekRange"`
	FiftyTwoWeekLow            float32 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh           float32 `json:"fiftyTwoWeekHigh"`
	Symbol                     string  `json:"symbol"`
	PriceEarnings              float32 `json:"priceEarnings"`
	EarningsPerShare           float32 `json:"earningsPerShare"`
	Logourl                    string  `json:"logourl"`
}

type Acoes struct {
	Ativo         string  `json:"ativo"`
	PrecoLucro    float32 `json:"precoLucro"`
	ValorAtual    float32 `json:"valorAtual"`
	VariacaoDoDia float32 `json:"variacaoDiaria"`
	MinimoAnual   float32 `json:"minimoAno"`
	MaximoAnual   float32 `json:"maximoAno"`
}

const baseURL = "https://brapi.dev/api/quote/"
const tokenURL = "substituir-pelo-seu-token"

func BuscarDadosApi(acao string, wg *sync.WaitGroup, ch chan<- Acoes) {
	defer wg.Done()

	response, err := http.Get(baseURL + acao + "?token=" + tokenURL)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer response.Body.Close()

	var retorno Retorno
	if err := json.NewDecoder(response.Body).Decode(&retorno); err != nil {
		fmt.Println("Erro ao decodificar resposta:", err)
		return
	}

	//Adiciona o retorno da api dentro do objeto ação
	for _, result := range retorno.Results {
		acaoObj := Acoes{
			Ativo:         result.Symbol,
			PrecoLucro:    result.PriceEarnings,
			ValorAtual:    result.RegularMarketPrice,
			VariacaoDoDia: result.RegularMarketChangePercent,
			MinimoAnual:   result.FiftyTwoWeekLow,
			MaximoAnual:   result.FiftyTwoWeekHigh,
		}
		ch <- acaoObj
	}
}

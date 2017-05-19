package main

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"fmt"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"github.com/naipath/bwfclient"
)

type Intent interface {
	name() string
	handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse)
}

type ElevatorPitch struct {
}

func (r ElevatorPitch) name() string {
	return "ElevatorPitch"
}
func (r ElevatorPitch) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeechSSML(
		`
		<speak>
			<p>Starting the elevator pitch.</p> 
			<p>With our newly build application <emphasis level="moderate">farmer bank</emphasis>, 
			it is possible to interact with your bank account in various ways using a natural dialogue.</p>
			<p><emphasis level="moderate">Farmer bank</emphasis> will be your new personal financial assistant who can take care of all your financial needs and questions.</p>
			<p>Our app can give you house estimations, mortgage capabilities, account balance information and much more.
			Customer interaction has never been this easy.</p>
			<p>All you have to do is ask questions.</p>
		</speak>
	`).EndSession(false)
}

type MaximumMortgage struct {
}

func (r MaximumMortgage) name() string {
	return "MaximumMortgage"
}
func (r MaximumMortgage) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	income, err := echoReq.GetSlotValue("income")
	fmt.Print("Total year income", income)
	if err != nil {
		echoResp.Reprompt("what is your year income").EndSession(false)
	} else {
		fmt.Print(income, err)
		yearIcome, _ := strconv.Atoi(income)
		echoResp.OutputSpeechSSML(fmt.Sprintf(`
		<speak>
		You can get a mortgage of <say-as interpret-as="cardinal">%s</say-as> 
		based on 1.7% with a fixed rate of 10 years
		</speak>`,
			strconv.Itoa(retrieveKoopsomBedr(yearIcome)))).EndSession(false)
	}
}

func retrieveKoopsomBedr(yearIncome int) int {
	resp, err := bwfClient.Request(bwfclient.BwfRequest{
		AanvragerBrutoJaarinkomenBedr: yearIncome,
		PartnerBrutoJaarinkomenBedr:   0,
	})
	if err != nil {
		fmt.Print(err)
	}
	return resp.MaxTeLenenObvInkomen.Tienjaarsrente.KoopsomBedr
}

type CancelIntent struct {
}

func (r CancelIntent) name() string {
	return "AMAZON.CancelIntent"
}
func (r CancelIntent) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Thank you for using farmer bank.").EndSession(true)
}

type StopIntent struct {
}

func (r StopIntent) name() string {
	return "AMAZON.StopIntent"
}
func (r StopIntent) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Thank you for using farmer bank.").EndSession(true)
}

type HouseEstimation struct {
}

func (r HouseEstimation) name() string {
	return "HouseEstimation"
}

func (r HouseEstimation) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	fmt.Print(echoReq.AllSlots())

	number, _ := echoReq.GetSlotValue("number")
	numberConvert, _ := strconv.Atoi(number)

	secondNumber, _ := echoReq.GetSlotValue("secondNumber")
	secondNumberConvert, _ := strconv.Atoi(secondNumber)

	echoResp.OutputSpeech("The value is " + strconv.Itoa(numberConvert+secondNumberConvert)).EndSession(false)
}

type YesOrNoResponse struct {
	Answer string `json:"answer"`
	Forced bool   `json:"forced"`
	Image  string `json:"image"`
}

type YesOrNo struct {
}

func (r YesOrNo) name() string {
	return "YesOrNo"
}

func (r YesOrNo) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	resp, _ := http.Get("https://yesno.wtf/api")
	defer resp.Body.Close()

	var data YesOrNoResponse
	json.NewDecoder(resp.Body).Decode(&data)

	echoResp.OutputSpeech("The dice has been rolled .... the answer is: " + data.Answer).EndSession(false)
}

type MadeBy struct {
}

func (r MadeBy) name() string {
	return "MadeBy"
}

func (r MadeBy) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Farmer bank was created by e bird and the care bearers").
		EndSession(false)
}

type Transaction struct {
	Type         string    `json:"Type"`
	Amount       string    `json:"Amount"`
	CounterParty string    `json:"CounterParty"`
	Date         time.Time `json:"Date"`
}

type Transactions []Transaction

type ListTransactions struct {
}

func (r ListTransactions) name() string {
	return "ListTransactions"
}

func (r ListTransactions) handle(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	counterParty, _ := echoReq.GetSlotValue("counterParty")
	transactionType, _ := echoReq.GetSlotValue("type")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Get("https://farmerbank.nl/transactions/Transactions")

	defer resp.Body.Close()

	var data Transactions
	json.NewDecoder(resp.Body).Decode(&data)

	if counterParty != "" {
		totalTransactions := 0
		for _, v := range data {
			if strings.HasPrefix(v.CounterParty, counterParty) {
				totalTransactions++
			}
		}
		echoResp.OutputSpeech("For the counterparty " + counterParty + " you had a total of " + strconv.Itoa(totalTransactions)).EndSession(false)
	} else if transactionType != "" {

		totalTransactionType := 0
		for _, v := range data {
			if strings.HasPrefix(v.Type, transactionType) {
				totalTransactionType++
			}
		}
		echoResp.OutputSpeech("For the transaction type " + transactionType + " you had a total of " + strconv.Itoa(totalTransactionType)).EndSession(false)
	} else {
		echoResp.OutputSpeech("You had a total amount of transactions of " + strconv.Itoa(len(data))).EndSession(false)
	}
}

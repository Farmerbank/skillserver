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
			<p>With our newly created banking platform <emphasis level="moderate">farmer bank</emphasis>,
			it is possible to interact with your bank account in various ways using just speech.</p>
			<p><emphasis level="moderate">Farmer bank</emphasis> will be your new personal financial assistant who can take care of all your financial needs and questions.</p>
			<p>take care of your banking needs while you cook, watch television, make breakfast of even do the dishes</p>
			<p>Our features include: mortgage estimates, transaction overviews, balance inquiries and bill payment management, all without ever touching a computer, tablet or even a phone.
			Customer interaction has never been this easy</p>
			<p>All you have to do is ask questions</p>
			<p><emphasis level="moderate">farmer bank</emphasis> brings an unparralled level of convinience to every home and office when it comes to banking services</p>
			<p>lowering the barrier between customer an bank to the level of natural language </p>
			<break time="1s"/>
			<p>by using state of the art speech recognition and artificial intelligence coupled with financial services aimed at satisfying even the most demanding customers <emphasis level="moderate">farmer bank</emphasis> will dominate the banking industry for years to come</p>
			<p><emphasis level="moderate">farmer bank</emphasis> is open and accesible to everyone </p>
			<p>ranging from the youngest customers to the elderly and even the visually impaired</p>
			<break time="1s"/>
			<p> <emphasis level="moderate">farmer bank</emphasis> the future of banking is here, and <emphasis level="moderate">everyone</emphasis> is invited</p>
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
		based on 1.7%% with a fixed rate of 10 years
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

	if counterParty != "" {
		transactions := getAllTransactions()
		for _, v := range transactions {
			if strings.HasPrefix(v.CounterParty, counterParty) {
				echoResp.OutputSpeech("For the counterparty " + counterParty + " you had a total of " + v.Amount + " of the type " + v.Type).EndSession(false)
				break
			}
		}
	} else if transactionType != "" {

		transactions := getAllTransactions()
		totalTransactionType := 0
		for _, v := range transactions {
			if strings.HasPrefix(v.Type, transactionType) {
				totalTransactionType++
			}
		}
		echoResp.OutputSpeech("For the transaction type " + transactionType + " you had a total of " + strconv.Itoa(totalTransactionType)).EndSession(false)
	} else {

		creditTransactions := getAllCreditTransactions()
		totalCredit := 0
		for _, v := range creditTransactions {
			creditConv, _ := strconv.Atoi(v.Amount[1:len(v.Amount)])
			totalCredit += creditConv
		}

		debitTransactions := getAllDebitTransactions()
		totalDebit := 0
		for _, v := range debitTransactions {
			debitConv, _ := strconv.Atoi(v.Amount[1:len(v.Amount)])
			totalDebit += debitConv
		}

		balance := totalDebit - totalCredit

		echoResp.OutputSpeech("Your current account balance is " + strconv.Itoa(balance)).EndSession(false)
	}
}

func getAllTransactions() Transactions {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Get("https://farmerbank.nl/transactions/Transactions")

	defer resp.Body.Close()

	var data Transactions
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

func getAllDebitTransactions() Transactions {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Get("https://farmerbank.nl/transactions/Transactions/Debit")

	defer resp.Body.Close()

	var data Transactions
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

func getAllCreditTransactions() Transactions {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, _ := client.Get("https://farmerbank.nl/transactions/Transactions/Credit")

	defer resp.Body.Close()

	var data Transactions
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

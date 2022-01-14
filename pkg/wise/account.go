package wise

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
)

var accountsPath = "/v1/accounts"

type Account struct {
	ID                int         `json:"id"`
	Business          interface{} `json:"business"`
	Profile           int         `json:"profile"`
	AccountHolderName string      `json:"accountHolderName"`
	Currency          string      `json:"currency"`
	Country           string      `json:"country"`
	Type              string      `json:"type"`
	Details           struct {
		Address struct {
			Country     string      `json:"country"`
			CountryCode string      `json:"countryCode"`
			FirstLine   string      `json:"firstLine"`
			PostCode    string      `json:"postCode"`
			City        string      `json:"city"`
			State       interface{} `json:"state"`
		} `json:"address"`
		Email                   interface{} `json:"email"`
		LegalType               string      `json:"legalType"`
		AccountHolderName       interface{} `json:"accountHolderName"`
		AccountNumber           interface{} `json:"accountNumber"`
		SortCode                interface{} `json:"sortCode"`
		Abartn                  interface{} `json:"abartn"`
		AccountType             interface{} `json:"accountType"`
		BankgiroNumber          interface{} `json:"bankgiroNumber"`
		IfscCode                interface{} `json:"ifscCode"`
		BsbCode                 interface{} `json:"bsbCode"`
		InstitutionNumber       interface{} `json:"institutionNumber"`
		TransitNumber           interface{} `json:"transitNumber"`
		PhoneNumber             interface{} `json:"phoneNumber"`
		BankCode                interface{} `json:"bankCode"`
		RussiaRegion            interface{} `json:"russiaRegion"`
		RoutingNumber           interface{} `json:"routingNumber"`
		BranchCode              interface{} `json:"branchCode"`
		Cpf                     interface{} `json:"cpf"`
		CardToken               interface{} `json:"cardToken"`
		IDType                  interface{} `json:"idType"`
		IDNumber                interface{} `json:"idNumber"`
		IDCountryIso3           interface{} `json:"idCountryIso3"`
		IDValidFrom             interface{} `json:"idValidFrom"`
		IDValidTo               interface{} `json:"idValidTo"`
		Clabe                   interface{} `json:"clabe"`
		SwiftCode               interface{} `json:"swiftCode"`
		DateOfBirth             interface{} `json:"dateOfBirth"`
		ClearingNumber          interface{} `json:"clearingNumber"`
		BankName                interface{} `json:"bankName"`
		BranchName              interface{} `json:"branchName"`
		BusinessNumber          interface{} `json:"businessNumber"`
		Province                interface{} `json:"province"`
		City                    interface{} `json:"city"`
		Rut                     interface{} `json:"rut"`
		Token                   interface{} `json:"token"`
		Cnpj                    interface{} `json:"cnpj"`
		PayinReference          interface{} `json:"payinReference"`
		PspReference            interface{} `json:"pspReference"`
		OrderID                 interface{} `json:"orderId"`
		IDDocumentType          interface{} `json:"idDocumentType"`
		IDDocumentNumber        interface{} `json:"idDocumentNumber"`
		TargetProfile           string      `json:"targetProfile"`
		TargetUserID            interface{} `json:"targetUserId"`
		TaxID                   interface{} `json:"taxId"`
		Job                     interface{} `json:"job"`
		Nationality             interface{} `json:"nationality"`
		InteracAccount          interface{} `json:"interacAccount"`
		Bban                    interface{} `json:"bban"`
		Town                    interface{} `json:"town"`
		PostCode                interface{} `json:"postCode"`
		Language                interface{} `json:"language"`
		BillerCode              interface{} `json:"billerCode"`
		CustomerReferenceNumber interface{} `json:"customerReferenceNumber"`
		Prefix                  interface{} `json:"prefix"`
		Iban                    interface{} `json:"IBAN"`
		Bic                     interface{} `json:"BIC"`
	} `json:"details"`
	User            int  `json:"user"`
	Active          bool `json:"active"`
	OwnedByCustomer bool `json:"ownedByCustomer"`
}

type AccountListOptions struct {
	ProfileID int    `json:"profileId"`
	Currency  string `json:"currency"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}

func accountListOptionsToQueryParams(opts *AccountListOptions) (map[string]string, error) {
	qp := make(map[string]string)
	raw := map[string]string{
		"profileId": strconv.FormatInt(int64(opts.ProfileID), 10),
		"currency":  opts.Currency,
		"limit":     strconv.FormatInt(int64(opts.Limit), 10),
		"offset":    strconv.FormatInt(int64(opts.Offset), 10),
	}
	for k, v := range raw {
		if v != "" && v != "0" {
			qp[k] = v
		}
	}
	return qp, nil
}

func PrintAccountsTable(accounts []Account, out io.Writer) {
	t := table.NewWriter()
	t.SetStyle(defaultTableStyle)
	t.SetOutputMirror(out)
	t.AppendHeader(table.Row{"Id", "Type", "IBAN"})
	for _, a := range accounts {
		t.AppendRow(table.Row{a.ID, a.Type, a.Details.Iban})
	}
	t.Render()
}

func (w WiseClient) ListAccounts(opts *AccountListOptions) ([]Account, error) {
	var (
		accounts []Account
		err      error
	)

	qp, err := accountListOptionsToQueryParams(opts)
	if err != nil {
		return nil, err
	}

	resp, err := w.cli.R().
		SetQueryParams(qp).
		Get(accountsPath)
	if err != nil {
		return nil, err
	}

	data := resp.Body()
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

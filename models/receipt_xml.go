package models

type ReceiptXML struct {
	Ide       IdentificationXML `xml:"proc>nfeProc>NFe>infNFe>ide"`
	Iss       IssuerXML         `xml:"proc>nfeProc>NFe>infNFe>emit"`
	Items     []ItemXML         `xml:"proc>nfeProc>NFe>infNFe>det"`
	Summary   TotalXML          `xml:"proc>nfeProc>NFe>infNFe>total"`
	ConsultAt string            `xml:"dataHora"`
}

type IdentificationXML struct {
	Title     string `xml:"natOp"`
	CreatedAt string `xml:"dhEmi"`
}

type IssuerXML struct {
	CNPJ    string          `xml:"CNPJ"`
	Name    string          `xml:"xNome"`
	Title   string          `xml:"xFant"`
	Address IsserAddressXML `xml:"enderEmit"`
}

type IsserAddressXML struct {
	Street   string `xml:"xLgr"`
	Number   string `xml:"nro"`
	District string `xml:"xBairro"`
	City     string `xml:"xMun"`
	UF       string `xml:"UF"`
	Zipcode  string `xml:"CEP"`
}

type ItemXML struct {
	Title string  `xml:"prod>xProd"`
	Price float64 `xml:"prod>vUnCom"`
	Qty   float64 `xml:"prod>qCom"`
	Total float64 `xml:"prod>vProd"`
	Tax   float64 `xml:"imposto>vTotTrib"`
}

type TotalXML struct {
	Price float64 `xml:"ICMSTot>vProd"`
	Tax   float64 `xml:"ICMSTot>vTotTrib"`
}

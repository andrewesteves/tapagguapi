package transformations

import "github.com/andrewesteves/tapagguapi/models"

func ReceiptToJSON(receiptXML models.ReceiptXML) models.ReceiptJSON {
	var receiptJSON models.ReceiptJSON
	receiptJSON.Ide.Title = receiptXML.Ide.Title
	receiptJSON.Ide.CreatedAt = receiptXML.Ide.CreatedAt

	receiptJSON.Iss.CNPJ = receiptXML.Iss.CNPJ
	receiptJSON.Iss.Name = receiptXML.Iss.Name
	receiptJSON.Iss.Title = receiptXML.Iss.Title

	receiptJSON.Iss.Address.Street = receiptXML.Iss.Address.Street
	receiptJSON.Iss.Address.Number = receiptXML.Iss.Address.Number
	receiptJSON.Iss.Address.District = receiptXML.Iss.Address.District
	receiptJSON.Iss.Address.City = receiptXML.Iss.Address.City
	receiptJSON.Iss.Address.UF = receiptXML.Iss.Address.UF
	receiptJSON.Iss.Address.Zipcode = receiptXML.Iss.Address.Zipcode

	for _, r := range receiptXML.Items {
		var item models.ItemJSON
		item.Title = r.Title
		item.Price = r.Price
		item.Qty = r.Qty
		item.Total = r.Total
		item.Tax = r.Tax
		receiptJSON.Items = append(receiptJSON.Items, item)
	}

	receiptJSON.Summary.Price = receiptXML.Summary.Price
	receiptJSON.Summary.Tax = receiptXML.Summary.Tax

	receiptJSON.ConsultAt = receiptXML.ConsultAt

	return receiptJSON
}

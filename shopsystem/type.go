package shopsystem

import "time"

// header
// authorization: bearer {accessToken}
// tobitauth: {accessToken}
// x-request-guid: {uuid}

// send array
type ServerOrderBatchRequest struct {
	Lastname  string        `json:"lastname"`
	OrderUID  interface{}   `json:"orderUid"`
	Vouchers  []interface{} `json:"vouchers"`
	UserID    int           `json:"userId"`
	CartPrice int           `json:"cartPrice"`
	PersonID  string        `json:"personId"`
	Fullname  string        `json:"fullname"`
	Firstname string        `json:"firstname"`
	Articles  []struct {
		ID      int           `json:"id"`
		GroupID string        `json:"groupId"`
		Amount  int           `json:"amount"`
		Price   int           `json:"price"`
		System  bool          `json:"system"`
		Options []interface{} `json:"options"`
	} `json:"articles"`
}

// get array
type ServerOrderBatchResponse struct {
	PersonID         string      `json:"personId"`
	Firstname        string      `json:"firstname"`
	Lastname         string      `json:"lastname"`
	Expires          int64       `json:"expires"`
	ProvisioningTime interface{} `json:"provisioningTime"`
	Articles         []struct {
		OrderedArticleID     int           `json:"orderedArticleId"`
		OrderID              int           `json:"orderId"`
		ID                   int           `json:"id"`
		Name                 string        `json:"name"`
		Amount               int           `json:"amount"`
		Price                float64       `json:"price"`
		PriceSum             float64       `json:"priceSum"`
		PayPrice             float64       `json:"payPrice"`
		TipPrice             float64       `json:"tipPrice"`
		TaxRate              float64       `json:"taxRate"`
		System               bool          `json:"system"`
		Unit                 string        `json:"unit"`
		Info                 interface{}   `json:"info"`
		Meta                 interface{}   `json:"meta"`
		ArticleLocationID    int           `json:"articleLocationId"`
		Options              []interface{} `json:"options"`
		Discountable         bool          `json:"discountable"`
		Config               []interface{} `json:"config"`
		Subscription         []interface{} `json:"subscription"`
		PerformancePeriodEnd interface{}   `json:"PerformancePeriodEnd"`
	} `json:"articles"`
	Vouchers    []interface{} `json:"vouchers"`
	VoucherList []interface{} `json:"voucherList"`
	UserID      int           `json:"userId"`
	Fullname    string        `json:"fullname"`
	Nfc         interface{}   `json:"nfc"`
	QrCode      interface{}   `json:"qrCode"`
	SiteID      interface{}   `json:"siteId"`
	OrderUID    string        `json:"orderUid"`
}

type ServerOrderConfirmCheck struct {
	ConfirmCheckOrderResults   []interface{} `json:"confirmCheckOrderResults"`
	ConfirmCheckArticleResults []interface{} `json:"confirmCheckArticleResults"`
	Proceedable                bool          `json:"proceedable"`
}

type ServerOrderConfirmStatus struct {
	ConfirmStatus bool `json:"confirmStatus"`
}

type ServerConfirmOrderBatchRequest struct {
	OrderList []struct {
		Lastname         string        `json:"lastname"`
		InvoiceRequested bool          `json:"invoiceRequested"`
		OrderUID         string        `json:"orderUid"`
		Vouchers         []interface{} `json:"vouchers"`
		UserID           int           `json:"userId"`
		TipFlag          bool          `json:"tipFlag"`
		BranchTipFactor  int           `json:"branchTipFactor"`
		CartPrice        int           `json:"cartPrice"`
		VoucherList      []interface{} `json:"voucherList"`
		PersonID         string        `json:"personId"`
		TipSum           int           `json:"tipSum"`
		Expires          int64         `json:"expires"`
		Fullname         string        `json:"fullname"`
		ProvisioningTime interface{}   `json:"provisioningTime"`
		Firstname        string        `json:"firstname"`
		Articles         []struct {
			PayPrice             int           `json:"payPrice"`
			PriceSum             int           `json:"priceSum"`
			Config               []interface{} `json:"config"`
			Unit                 string        `json:"unit"`
			OrderedArticleID     int           `json:"orderedArticleId"`
			Meta                 interface{}   `json:"meta"`
			ArticleLocationID    int           `json:"articleLocationId"`
			OrderID              int           `json:"orderId"`
			Price                int           `json:"price"`
			PerformancePeriodEnd interface{}   `json:"PerformancePeriodEnd"`
			Name                 string        `json:"name"`
			TaxRate              int           `json:"taxRate"`
			Info                 interface{}   `json:"info"`
			TipPrice             int           `json:"tipPrice"`
			System               bool          `json:"system"`
			Amount               int           `json:"amount"`
			ID                   int           `json:"id"`
			Discountable         bool          `json:"discountable"`
			Subscription         []interface{} `json:"subscription"`
			Options              []interface{} `json:"options"`
		} `json:"articles"`
	} `json:"orderList"`
	ConfirmDto struct {
		Output struct {
			BranchProcessorID int         `json:"branchProcessorId"`
			ProcessorText     string      `json:"processorText"`
			ProcessorName     string      `json:"processorName"`
			Comment           string      `json:"comment"`
			Name              string      `json:"name"`
			Subtitle          string      `json:"subtitle"`
			DeliveryTime      interface{} `json:"deliveryTime"`
			Text              interface{} `json:"text"`
			OutputFlags       int         `json:"outputFlags"`
			Email             string      `json:"email"`
			Phone             string      `json:"phone"`
		} `json:"output"`
		FirstName     string        `json:"firstName"`
		UserID        int           `json:"userId"`
		OrderUID      string        `json:"orderUid"`
		DeliveryHours int           `json:"deliveryHours"`
		ForwardURL    string        `json:"forwardUrl"`
		PaymentMode   int           `json:"paymentMode"`
		ConfirmTime   time.Time     `json:"confirmTime"`
		Vouchers      []interface{} `json:"vouchers"`
		CouponPayment bool          `json:"couponPayment"`
		CashPayment   bool          `json:"cashPayment"`
	} `json:"confirmDto"`
	TobitAccessToken string `json:"tobitAccessToken"`
}

// get array
type ServerConfirmOrderBatchResponse struct {
	Message    interface{} `json:"message"`
	StatusCode int         `json:"statusCode"`
	User       struct {
		ID        int         `json:"id"`
		Name      string      `json:"name"`
		OrderUID  string      `json:"orderUid"`
		OrderHash string      `json:"orderHash"`
		Documents interface{} `json:"documents"`
	} `json:"user"`
	OrderUID  string `json:"orderUid"`
	OrderHash string `json:"orderHash"`
}

type ServerOrderTrue struct {
	ID                 int         `json:"id"`
	UID                string      `json:"uid"`
	UserID             int         `json:"userId"`
	UserPersonID       string      `json:"userPersonId"`
	UserFirstName      string      `json:"userFirstName"`
	UserLastName       string      `json:"userLastName"`
	UserName           string      `json:"userName"`
	PayerUserID        int         `json:"payerUserId"`
	PayerUserPersonID  string      `json:"payerUserPersonId"`
	PayerUserFirstName string      `json:"payerUserFirstName"`
	PayerUserLastName  string      `json:"payerUserLastName"`
	PayerUserName      string      `json:"payerUserName"`
	ArticleLocationID  int         `json:"articleLocationId"`
	Extras             interface{} `json:"extras"`
	TipFlag            bool        `json:"tipFlag"`
	TipPriceSum        float64     `json:"tipPriceSum"`
	BranchTipFactor    float64     `json:"branchTipFactor"`
	Expires            int64       `json:"expires"`
	Articles           []struct {
		OrderedArticleID     int           `json:"orderedArticleId"`
		OrderID              int           `json:"orderId"`
		ID                   int           `json:"id"`
		Name                 string        `json:"name"`
		Amount               int           `json:"amount"`
		Price                float64       `json:"price"`
		PriceSum             float64       `json:"priceSum"`
		PayPrice             float64       `json:"payPrice"`
		TipPrice             float64       `json:"tipPrice"`
		TaxRate              float64       `json:"taxRate"`
		System               bool          `json:"system"`
		Unit                 string        `json:"unit"`
		Info                 interface{}   `json:"info"`
		Meta                 interface{}   `json:"meta"`
		ArticleLocationID    int           `json:"articleLocationId"`
		Options              []interface{} `json:"options"`
		Discountable         bool          `json:"discountable"`
		Config               []interface{} `json:"config"`
		Subscription         []interface{} `json:"subscription"`
		PerformancePeriodEnd interface{}   `json:"PerformancePeriodEnd"`
	} `json:"articles"`
	VoucherList          []interface{} `json:"voucherList"`
	Vouchers             []interface{} `json:"vouchers"`
	SubscriptionArticles []interface{} `json:"subscriptionArticles"`
	InvoiceRequested     bool          `json:"invoiceRequested"`
	ProvisioningTime     interface{}   `json:"provisioningTime"`
	ReservationUID       interface{}   `json:"reservationUid"`
}

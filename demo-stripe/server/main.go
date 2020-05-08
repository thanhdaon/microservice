package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"github.com/stripe/stripe-go/paymentintent"
)

func init() {
	stripe.Key = "sk_test_fsxIrATcKXUOe0NrBZvhzmTw00ocgSAiya"
}

type CheckoutData struct {
	ClientSecret string `json:"client_secret"`
}

func main() {
	http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(1099),
			Currency: stripe.String(string(stripe.CurrencyUSD)),
		}
		params.AddMetadata("integration_check", "accept_a_payment")
		pi, err := paymentintent.New(params)
		check(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CheckoutData{
			ClientSecret: pi.ClientSecret,
		})
	})

	http.HandleFunc("/sessionid", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		params := &stripe.CheckoutSessionParams{
			PaymentMethodTypes: stripe.StringSlice([]string{
				"card",
			}),
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				{
					Name:        stripe.String("T-shirt"),
					Description: stripe.String("Comfortable cotton t-shirt"),
					Amount:      stripe.Int64(500),
					Currency:    stripe.String(string(stripe.CurrencyUSD)),
					Quantity:    stripe.Int64(1),
				},
			},
			SuccessURL: stripe.String("https://example.com/success?session_id={CHECKOUT_SESSION_ID}"),
			CancelURL:  stripe.String("https://example.com/cancel"),
		}

		session, err := session.New(params)
		check(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(session)
	})

	log.Println("server is running on port: 8000")

	http.ListenAndServe(":8000", nil)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func check(err error) {
	if err != nil {
		log.Println("okok")
	}
}

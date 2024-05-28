// microservice2/model/event.go
package model

type Event struct {
    ID   string `json:"id"`
    Type string `json:"type"`
    Data string `json:"data"`
}
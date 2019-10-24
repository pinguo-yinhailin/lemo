package name

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Name interface {
	Generate() string
}

type name struct {
	dataFile string
	data     *DataName
}

func New(dataFile string) (Name, error) {
	n := &name{}

	b, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}
	n.dataFile = dataFile

	d := new(DataName)
	if err := json.Unmarshal(b, d); err != nil {
		return nil, fmt.Errorf("failed to decode name data: %v", err)
	}
	n.data = d

	return n, nil
}

func (n *name) Generate() string {
	data := n.data

	lastName := n.RandOne(data.LastName)

	var firstName string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= rand.Intn(2); i++ {
		firstName += n.RandOne(data.FirstName)
	}

	return lastName + firstName
}

func (n *name) RandOne(data []string) string {
	rand.Seed(time.Now().UnixNano())
	return data[rand.Intn(len(data))]
}

type DataName struct {
	FirstName []string `json:"first_name"`
	LastName  []string `json:"last_name"`
}

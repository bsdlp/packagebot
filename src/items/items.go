package items

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/gobuffalo/packr"
)

var fieldsPerRecord = 8

func unmarshalItem(records []string) (Item, error) {
	fields := len(records)
	if fields != fieldsPerRecord {
		return Item{}, fmt.Errorf("expected %d fields, got %d", fieldsPerRecord, fields)
	}
	item := Item{
		Name: records[0],
		Tier: records[1],
		Type: records[2],
		/*
			Cooldown:  records[5],
			BuyPrice:  records[6],
			SellPrice: records[7],
		*/
	}

	var err error
	item.Minimum, err = strconv.ParseUint(records[3], 10, 64)
	if err != nil {
		return item, err
	}

	item.Maximum, err = strconv.ParseUint(records[4], 10, 64)
	if err != nil {
		return item, err
	}
	return item, nil
}

// Price is the price of a weapon in cents. If the price is negative, then it
// cannot be bought/sold
type Price int64

// TODO: implement
func (p *Price) unmarshalText(s string) error {
	return nil
}

// Item describes an item
type Item struct {
	Name      string
	Tier      string
	Type      string
	Minimum   uint64
	Maximum   uint64
	Cooldown  time.Duration
	BuyPrice  Price
	SellPrice Price
}

func loadItems(data []byte, dst map[string]Item) error {
	reader := csv.NewReader(bytes.NewReader(data))
	reader.ReuseRecord = true
	reader.FieldsPerRecord = fieldsPerRecord

	for {
		records, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		item, err := unmarshalItem(records)
		if err != nil {
			return err
		}

		if _, exists := dst[item.Name]; exists {
			return fmt.Errorf("duplicate record for %q", item.Name)
		}

		dst[item.Name] = item
	}
}

var items map[string]Item

func init() {
	items = map[string]Item{}
	err := loadItems(packr.NewBox("./data").Bytes("items.csv"), items)
	if err != nil {
		panic(err)
	}
}

// GetItem retrieves an item
func GetItem(item string) (Item, bool) {
	i, ok := items[item]
	return i, ok
}

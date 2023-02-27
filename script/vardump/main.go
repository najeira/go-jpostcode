package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/syumai/go-jpostcode/internal/address"
)

func main() {
	var firstPostCodes []string

	var buf bytes.Buffer
	filepath.Walk("./jpostcode-data/data/json", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}

		dataFile, err := os.Open("./jpostcode-data/data/json/" + info.Name())
		if err != nil {
			panic(err)
		}
		defer dataFile.Close()

		var addressMap map[string]interface{}
		if err := json.NewDecoder(dataFile).Decode(&addressMap); err != nil {
			panic(err)
		}

		firstPostCode := strings.TrimSuffix(info.Name(), ".json")
		firstPostCodes = append(firstPostCodes, firstPostCode)

		buf.Reset()
		fmt.Fprintln(&buf, `package data`)
		//fmt.Fprintln(&buf, `import "github.com/syumai/go-jpostcode"`)
		fmt.Fprintln(&buf, `import "github.com/syumai/go-jpostcode/internal/address"`)

		fmt.Fprintf(
			&buf,
			`var Addresses%s = map[string][]*address.Address{`,
			firstPostCode,
		)
		fmt.Fprintln(&buf)

		for secondPostCode, val := range addressMap {
			postCode := firstPostCode + secondPostCode
			fmt.Println(postCode)

			fmt.Fprintf(
				&buf,
				`"%s": {`,
				secondPostCode,
			)
			fmt.Fprintln(&buf)

			addresses := decodeAddresses(val)
			for _, addr := range addresses {
				s := fmt.Sprintf("%#v", addr)
				//s = strings.TrimPrefix(s, "&jpostcode.Address")
				s = strings.TrimPrefix(s, "&address.Address")
				fmt.Fprint(&buf, s)
				fmt.Fprintln(&buf, ",")
			}

			fmt.Fprintln(&buf, "},")
		}
		fmt.Fprintln(&buf, "}")

		fileName := fmt.Sprintf("./internal/data/a%s.go", firstPostCode)
		writeFile(fileName, buf.Bytes())
		return nil
	})

	buf.Reset()
	//fmt.Fprintln(&buf, `package data`)
	////fmt.Fprintln(&buf, `import "github.com/syumai/go-jpostcode"`)
	//fmt.Fprintln(&buf, `import "github.com/syumai/go-jpostcode/internal/address"`)
	//fmt.Fprintln(
	//	&buf,
	//	`var Addresses = map[string]map[string][]*address.Address{`,
	//)
	//for _, firstPostCode := range firstPostCodes {
	//	fmt.Fprintf(
	//		&buf,
	//		`"%s": addresses%s,`,
	//		firstPostCode,
	//		firstPostCode,
	//	)
	//	fmt.Fprintln(&buf)
	//}
	//fmt.Fprintln(&buf, "}")
	//writeFile("000", buf.Bytes())

	fmt.Fprintln(&buf, `package jpostcode`)
	fmt.Fprintln(&buf, `import "github.com/syumai/go-jpostcode/internal/address"`)
	fmt.Fprintln(&buf, `import "github.com/syumai/go-jpostcode/internal/data"`)
	fmt.Fprintln(
		&buf,
		`type varAdapter struct {}`,
	)
	fmt.Fprintln(
		&buf,
		`func (a *varAdapter) SearchAddressesFromPostCode(postCode string) ([]*Address, error) {
first := postCode[:3]
second := postCode[3:]
switch first {`,
	)
	for _, firstPostCode := range firstPostCodes {
		fmt.Fprintf(
			&buf,
			`case "%s":`,
			firstPostCode,
		)
		fmt.Fprintln(&buf)
		fmt.Fprintf(
			&buf,
			`return fromDataAddresses(data.Addresses%s, second)`,
			firstPostCode,
		)
		fmt.Fprintln(&buf)
	}
	fmt.Fprintln(
		&buf,
		`}
	return nil, ErrNotFound
}

func fromDataAddresses(m map[string][]*address.Address, p string) ([]*Address, error) {
	d, ok := m[p]
	if !ok {
		return nil, ErrNotFound
	}
	r := make([]*Address, len(d))
	for i, a := range d {
		r[i] = (*Address)(a)
	}
	return r, nil
}`,
	)
	writeFile("./var_adapter.go", buf.Bytes())
}

func decodeAddresses(addressData interface{}) []*address.Address {
	var addresses []*address.Address
	switch reflect.TypeOf(addressData).Kind() {
	case reflect.Slice:
		rawAddrs, ok := addressData.([]interface{})
		if !ok {
			panic("")
		}
		for _, rawAddr := range rawAddrs {
			addr, err := address.FromMap(rawAddr)
			if err != nil {
				panic("")
			}
			addresses = append(addresses, addr)
		}
	default:
		addr, err := address.FromMap(addressData)
		if err != nil {
			panic("")
		}
		addresses = append(addresses, addr)
	}
	return addresses
}

func writeFile(fileName string, data []byte) {
	formatted, err := format.Source(data)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile(fileName, formatted, 0666); err != nil {
		panic(err)
	}
}

## Usage:

1. Run go server
```
$ go run main.go
```
2. You may test your api with postman. Just look for your port 8000 and choose "GET" method for request.

![Alt text](./assets/image.png)

3. It's important to note that Go doesn't have any methods for implementing self-closing tags. To solve this problem, we can use standard methods for working with I/O interfaces.For example:

```Go
package main

import (
    "bytes"
    "encoding/xml"
    "fmt"
    "os"
)

// Person represents a <person> node in the XML
type Person struct {
    XMLName   xml.Name   `xml:"Players"`
    DataItems []dataItem `xml:"DataItem"`
}

// Skill represents a <skill> node in the XML
type dataItem struct {
    XMLName        xml.Name `xml:"DataItem"`
    Name           string   `xml:"skillName,attr"`
    YearsPracticed int64    `xml:"practice,attr"`
    Level          string   `xml:"level,attr"`
}

func main() {
    players := Person{
        DataItems: []dataItem{
            {Name: "Soccer", YearsPracticed: 3, Level: "Newbie"},
            {Name: "Basketball", YearsPracticed: 4, Level: "State"},
            {Name: "Baseball", YearsPracticed: 10, Level: "National"},
        },
    }
    players.DataItems = append(players.DataItems, players.DataItems...)
    players.DataItems = append(players.DataItems, players.DataItems...)
    players.DataItems = append(players.DataItems, players.DataItems...)
    players.DataItems = append(players.DataItems, players.DataItems...)
    players.DataItems = append(players.DataItems, players.DataItems...)

    out := new(bytes.Buffer)

    enc := xml.NewEncoder(out)
    enc.Indent("", "  ")
    if err := enc.Encode(players); err != nil {
        fmt.Printf("error: %v\n", err)
    }

    b := bytes.ReplaceAll(out.Bytes(), []byte("></DataItem>"), []byte("/>"))
    os.Stdout.Write(b)
}

```

*****

### Resources:

* [link](https://stackoverflow.com/questions/38118100/go-encoding-xml-how-can-i-marshal-self-closing-elements)
* [another link))](https://stackoverflow.com/questions/26371965/when-generating-an-xml-file-with-go-how-do-you-create-a-doctype-declaration)
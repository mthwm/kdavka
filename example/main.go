package main

import (
	"fmt"
	"os"
	"time"

	"kdavka/ascii/davka16"
)

func main() {
	// Sample header for dávka 16
	header := davka16.Header{
		Char:  "P",          // Original batch
		Dtyp:  "16",         // Dávka type
		Ico:   "12345678",   // Sample IČZ
		Pob:   "0000",       // Branch code
		Rok:   2025,         // Year
		Mes:   8,            // Month (August)
		Cid:   1,            // Batch number
		Poc:   0,            // Auto-set to number of docs
		Body:  0,            // Optional points
		Fin:   5000.00,      // Total amount
		Dpp:   "1",          // DPP flag
		Dvdr1: "16:6.2.XXX", // Version per spec
		Dvdr2: "",           // Empty
		Ddtyp: "",           // Empty
	}

	// Sample document with Czech diacritics to test CP852 encoding
	now := time.Date(2025, 8, 18, 0, 0, 0, 0, time.UTC)
	docs := []davka16.Document{
		{
			L: davka16.DocumentL{
				Icll:    123,
				Cdok:    4567890,
				Ind1:    "IND123456",
				Cop:     "COP1",
				TypLp:   "K", // Comprehensive care
				Jmeno:   "Novák Jan",
				Cp:      "1234567890",
				JmenoPr: "",
				CpPr:    "",
				Dnast:   now.AddDate(0, 0, -10), // 08.08.2025
				Dukon:   now,                    // 18.08.2025
				Dodj:    now.AddDate(0, 0, 1),   // 19.08.2025
				Jmevyst: "Lékař Petr",
				Dvyst:   now,
				Prod:    0,
				KodUko:  "1",
				CenaPob: 5000.00,
			},
			Naklads: []davka16.NakladU{
				{
					Datod:     now.AddDate(0, 0, -10),
					KodNak:    "1", // Accommodation
					KodNak1:   "",
					Doba:      10,
					Sazba:     500.00,
					Cena:      5000.00,
					Luzko:     1,
					Kateg:     "STD",
					KodPrerus: "0",
				},
			},
			Sdelenis: []davka16.SdeleniS{
				{
					CisR: 1,
					Text: "Pacient dokončil léčbu bez komplikací.",
				},
			},
		},
	}

	// Generate dávka 16 file content
	data, err := davka16.Generate(header, docs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating dávka: %v\n", err)
		os.Exit(1)
	}

	// Write to file
	if err := os.WriteFile("KDAVKA.001", data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated KDAVKA.001")
}

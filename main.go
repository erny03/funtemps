package main

import (
	"flag"
	"fmt"

	"github.com/erny03/funtemps/conv"
)

var tempVerdi float64
var tempEnhet string
var outputEnhet string

func init() {
	flag.Float64Var(&tempVerdi, "temp", 0.0, "temperaturverdi som skal konverteres")
	flag.StringVar(&tempEnhet, "enhet", "C", "enhetsbetegnelse for temperaturverdien som skal konverteres (C, F, K)")
	flag.StringVar(&outputEnhet, "out", "C", "enhetsbetegnelse for den konverterte temperaturverdien (C, F, K)")
}

func main() {
	flag.Parse()

	var konvertertTemp float64
	var outputEnhetNavn string

	if tempEnhet == "C" {
		if outputEnhet == "F" {
			konvertertTemp = conv.CelsiusToFarhenheit(tempVerdi)
			outputEnhetNavn = "°F"
		} else if outputEnhet == "K" {
			konvertertTemp = conv.CelsiusToKelvin(tempVerdi)
			outputEnhetNavn = "K"
		} else {
			konvertertTemp = tempVerdi
			outputEnhetNavn = "°C"
		}
	}

	if tempEnhet == "F" {
		if outputEnhet == "C" {
			konvertertTemp = conv.FarhenheitToCelsius(tempVerdi)
			outputEnhetNavn = "°C"
		} else if outputEnhet == "K" {
			konvertertTemp = conv.FarhenheitToKelvin(tempVerdi)
			outputEnhetNavn = "K"
		} else {
			konvertertTemp = tempVerdi
			outputEnhetNavn = "°F"
		}
	}

	if tempEnhet == "K" {
		if outputEnhet == "C" {
			konvertertTemp = conv.KelvinToCelsius(tempVerdi)
			outputEnhetNavn = "°C"
		} else if outputEnhet == "F" {
			konvertertTemp = conv.KelvinToFarhenheit(tempVerdi)
			outputEnhetNavn = "°F"
		} else {
			konvertertTemp = tempVerdi
			outputEnhetNavn = "K"
		}
	}

	fmt.Printf("%.2f%s er %.2f%s\n", tempVerdi, tempEnhet, konvertertTemp, outputEnhetNavn)
}

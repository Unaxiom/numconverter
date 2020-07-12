# numconverter

Package to convert numbers to text for Indian Number System

# API

```golang

numconverter.Convert(19) // nineteen
numconverter.Convert(101010) // one lakh one thousand ten
numconverter.Convert(9999999) // ninety nine lakh ninety nine thousand nine hundred ninety nine

numconverter.Convert(96273927) // nine crore sixty two lakh seventy three thousand nine hundred twenty seven

numconverter.ConvertAnd(121) // one hundred and twenty one
numconverter.Convert(121) // one hundred twenty one

```

# Limits

Package can handle numbers upto 9999999999 (nine hundred ninety nine crore ninety nine lakh ninety nine thousand nine hundred ninety nine)
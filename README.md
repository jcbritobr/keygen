# Keygen
Keygen is a cli golang application that automates the creation of user keys. It possible to choose between symbols, uppercase, lowercase and number characters. Also, is possible to set the key size.

* Compile - In root folder run the command:
```
$ go build -v -ldflags "-s -w"
```

* Test - Also in root folder run the command:
```
$ go test .\uniengine\
ok      github.com/jcbritobr/keygen/uniengine   0.594s
PS C:\Workspace\go\keygen> go test .\uniengine\ -v
=== RUN   TestRandomInRangeIsCharacterKindAndLength
--- PASS: TestRandomInRangeIsCharacterKindAndLength (0.00s)
=== RUN   TestSumTrueValue
--- PASS: TestSumTrueValue (0.00s)
=== RUN   TestGenerateKeyFunction
--- PASS: TestGenerateKeyFunction (0.00s)
PASS
ok      github.com/jcbritobr/keygen/uniengine   (cached)
``` 

* Usage
```
$ .\keygen.exe --help           
Keygen is a very fast random key generator.
        Its possible to choose between lowercase, uppercase, number and symbol characters,
        and also the key's size

Usage:
  keygen [flags]

Flags:
  -h, --help         help for keygen
      --length int   inserts the key length (default 8)
  -l, --lowercase    inserts lowercase characters in key
  -n, --number       inserts digit characters in key
  -s, --symbol       inserts symbol characters in key
  -u, --uppercase    inserts uppercase characters in key

$ .\keygen.exe --length 80 -luns
g07ouGw7&l0UWKEm(%&+I"GG-7g573"7(43!j++-(QBg1mWd6N gDpa4M'0MmI1tlPy!Q 1M,0P8sa)1
```
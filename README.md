# Keygen
Keygen is a cli golang application that automates the creation of user keys. It possible to choose between symbols, uppercase, lowercase and number characters. Also, is possible to set the key size.

* Compile - In root folder run the command:
```
$ go build -v -ldflags "-s -w"
```

* Test - Also in root folder run the command:
```
$ go test . -v

=== RUN   TestGenerateKeyMustGenerateLowercaseString\
--- PASS: TestGenerateKeyMustGenerateLowercaseString (0.00s)\
=== RUN   TestGenerateKeyMustGenerateUppercaseString\
--- PASS: TestGenerateKeyMustGenerateUppercaseString (0.00s)\
=== RUN   TestGenerateKeyMustGenerateNumberString\
--- PASS: TestGenerateKeyMustGenerateNumberString (0.00s)\
=== RUN   TestGenerateKeyMustGenerateSymbolString\
--- PASS: TestGenerateKeyMustGenerateSymbolString (0.00s)\
=== RUN   TestGenerateKeyMustHaveCorrectKeySize\
--- PASS: TestGenerateKeyMustHaveCorrectKeySize (0.00s)\
PASS\
ok      github.com/jcbritobr/keygen     0.745s
``` 

* Usage
```
$ .\keygen.exe -h         
Usage of C:\Workspace\go\keygen\keygen.exe:
  -k int
        Use -k for the key size (default 14)
  -l    Use -l to insert lowercase characters in key (default true)
  -n    Use -n to insert uppercase characters in key
  -s    Use -s to insert uppercase characters in key
  -u    Use -u to insert uppercase characters in ke

$ .\keygen.exe -k 20 -s -n
70$50#lwu>3%?51g2(m:
```
# Introduction

REST API testing in Go with Resty. The API is [Zippopotam.us API](http://api.zippopotam.us/).

This defines a struct LocationResponse with a single element Country. The json:"country" tag tells Go that it should populate this element with the value of the country element in the JSON response body. We don’t have to worry about the other elements in the response, those simply will not be mapped (unless you need them in a check, then you’ll need to add them to the struct, too).

# Run tests

````
Thanhs-MacBook-Pro:go-resty-fun tamara-thanh$ go test -v
=== RUN   TestZippopotamUsSuite
=== RUN   TestZippopotamUsSuite/Test_GetUs90210_ContentTypeShouldEqualApplicationJson
=== RUN   TestZippopotamUsSuite/Test_GetUs90210_CountryShouldEqualUnitedStates
=== RUN   TestZippopotamUsSuite/Test_GetUs90210_StatusCodeShouldEqual200
--- PASS: TestZippopotamUsSuite (0.58s)
    --- PASS: TestZippopotamUsSuite/Test_GetUs90210_ContentTypeShouldEqualApplicationJson (0.22s)
    --- PASS: TestZippopotamUsSuite/Test_GetUs90210_CountryShouldEqualUnitedStates (0.19s)
    --- PASS: TestZippopotamUsSuite/Test_GetUs90210_StatusCodeShouldEqual200 (0.17s)
PASS
ok      github.com/peterngtr/apiFunWithGo       0.784s
````
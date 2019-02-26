# gotry
golang try catch实现

## example
```
gotry.Try(func(assert *gotry.Assert) {
		assert.Assert(true, "sss%s", "dd")
	}).Catch(func(err error) {
		fmt.Println(err.Error())
	})

	gotry.Try(func(assert *gotry.Assert) {
		assert.AssertErr(fmt.Errorf("ddd%d", 222), "sss%s", "dd")
	}).Catch(func(err error) {
		fmt.Println(err.Error())
	})
```
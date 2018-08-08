# slow
Go package for Debounce and Throttle functions.  
See [stackoverflow answers](https://stackoverflow.com/a/25991510) for more details about what is throttle and debounce.

# install
```bash
go get github.com/chneau/slow
```

# usage
```go
    ...
    i := 0
    fn := func() {
        i++
    }
    options := &slow.Options{
        Trailing: true,
        Leading:  true,
        MaxWait:  time.Millisecond * 500,
    }
    // options = nil // can be null, see defaults
    throttled := slow.Throttle(fn, time.Millisecond*50, options)
    debounced := slow.Debounce(fn, time.Millisecond*50, options)
    // the returned functions will be debounced / throttled.
    // so some calls will be no-op depending on the options
    throttled()
    throttled()
    ...
```

# thanks
Thanks to [lodash](https://lodash.com/) from where the implementation comes from. 

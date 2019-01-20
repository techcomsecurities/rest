# rest
HTTP utilities for working with REST API 
 
## Usage 
- Request
```
resp, err := rest.NewRequest().
            AddHeader("Content-Type", "application-json").
            AddHeader("Authorization", "Bearer "+token).
            Get("https://api.tcbs.com.vn/)
```
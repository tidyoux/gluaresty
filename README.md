# gluaresty

```lua
local resp, err = resty.R()
    :SetQueryParams({key = "value"})
    :SetHeaders({key = "value"})
    :SetBody({key = "value"})
    :Execute("POST", "http://url")
if err ~= nil then
    error(err.Error())
end

resp = resty.P(resp)
```

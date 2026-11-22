# wrldview

.env layout:
``````
LOGPATH = path/to/logs
PORT = 6767
``````
## endpoints:

`POST /book` - add book to base
```
{
  "title": "Clocks",
  "author": "agata kristi",
  "pages": 333,
}
```
`PATCH /book?id=` - patch book in base. You can use only certain fields.
```
{
  "title": "Fa",
  "author": "fafa",
  "pages": 34343
}
```

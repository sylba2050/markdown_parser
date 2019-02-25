**I don't recommend using this**

# API server to convert markdown to HTML
## Demo
### Form data
request
```bash
curl -X POST http://153.126.139.150:8080/api/md -d 'md=hoge'
```
response
```html
<p>hoge</p>
```

### JSON data
request
```bash
curl -X POST http://153.126.139.150:8080/api/md -H 'Content-Type: application/json' -d '{"md":"hoge"}'
```
response
```html
<p>hoge</p>
```

### File data
in hoge.md
```text:hoge.md
hoge
```

request
```bash
curl -X POST http://153.126.139.150:8080/api/md/file -F "file=@hoge.md"
```

response
```html
<p>hoge</p>
```

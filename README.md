# wikibase-dump-filter

## Compile

        CGO_ENABLED=0 go build -o src/main/target/filter -v -work -x src/main/go/filter/main.go


## Execute

        
        cat [dir]/sample-dump-20150815.json.gz | gzip -d | ./filter  --claim P31:Q5

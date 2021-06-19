package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	claim      string
	keep       string
	verbose    bool
	validClaim = regexp.MustCompile(`^[A-Z][0-9]+$`)
)

func init() {
	flag.StringVar(&claim, "claim", "", "Entity filter. Example: P31:Q5")
	flag.StringVar(&keep, "keep", "", "Comma-separated attributes to keep")
	flag.BoolVar(&verbose, "verbose", false, "Should speak on stdout")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// {"type":"item","aliases":{},"labels":{},"descriptions":{},"sitelinks":{},"id":"Q216","claims":{"P31":[{"rank":"normal","mainsnak":{"snaktype":"value","property":"P31","datavalue":{"type":"wikibase-entityid","value":{"entity-type":"item","numeric-id":515}},"datatype":"wikibase-item"},"id":"q216$71CEE092-9B75-4783-B479-F651841ECCEA","type":"statement"},{"rank":"normal","mainsnak":{"snaktype":"value","property":"P31","datavalue":{"type":"wikibase-entityid","value":{"entity-type":"item","numeric-id":5119}},"datatype":"wikibase-item"},"id":"q216$91CCAEAD-8B4E-4E1B-AC52-9552A411031F","type":"statement"},{"rank":"normal","mainsnak":{"snaktype":"value","property":"P31","datavalue":{"type":"wikibase-entityid","value":{"entity-type":"item","numeric-id":1363145}},"datatype":"wikibase-item"},"id":"Q216$dff956ef-483d-feec-f6a9-baf0d915e3db","type":"statement"}]}},
	argsWithProg := os.Args
	claimFilter := parseClaimFilter()

	if verbose {
		log.Println("Executing filter")
		log.Println(argsWithProg)
		log.Println(strings.Repeat("▔", 65))
		log.Println(strings.Repeat("▔", 65))
		log.Println(strings.Repeat("▔", 65))
		log.Printf("Claim filter: [%s]\n", claimFilter)
		log.Println("reading from stdin:")
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if "]" == text || "[" == text {
			continue
		}
		data := parseText(clean(text))
		isValidClaim := filterClaims(data, claimFilter)
		if isValidClaim {
			fmt.Println(text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func clean(text string) (trimmed string) {
	return strings.Trim(text, ",$")
}

func parseText(text string) (data map[string]interface{}) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(text), &dat); err != nil {
		panic(err)
	}
	return dat
}

type ClaimFilter struct {
	P string //P
	Q string //Q
}

func parseClaimFilter() (claimFilter *ClaimFilter) {
	split := strings.Split(claim, ":")
	if len(split) != 2 {
		return nil
	}
	var p = split[0]
	var q = split[1]
	if validClaim.MatchString(p) && validClaim.MatchString(q) {
		return &ClaimFilter{
			P: split[0],
			Q: split[1],
		}
	} else {
		return nil
	}
}

func filterClaims(dat map[string]interface{}, claimFilter *ClaimFilter) bool {
	if claimFilter == nil {
		return true
	}
	if dat["claims"] != nil {
		var claims = dat["claims"].(map[string]interface{})
		p := claims[claimFilter.P]
		if p == nil {
			return false
		} else {
			for _, v := range p.([]interface{}) {
				props := v.(map[string]interface{})
				if props["id"] != nil {
					var id = strings.ToUpper(props["id"].(string))
					if strings.Split(id, "$")[0] == claimFilter.Q {
						return true
					}
				}
			}
			return false
		}
	}
	return false
}

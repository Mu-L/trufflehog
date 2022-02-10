package textmagic

import (
	"context"
	b64 "encoding/base64"
	"fmt"

	// "log"
	"regexp"
	"strings"

	"net/http"

	"github.com/trufflesecurity/trufflehog/v3/pkg/common"
	"github.com/trufflesecurity/trufflehog/v3/pkg/detectors"
	"github.com/trufflesecurity/trufflehog/v3/pkg/pb/detectorspb"
)

type Scanner struct{}

// Ensure the Scanner satisfies the interface at compile time
var _ detectors.Detector = (*Scanner)(nil)

var (
	client = common.SaneHttpClient()

	//Make sure that your group is surrounded in boundry characters such as below to reduce false positives
	keyPat  = regexp.MustCompile(detectors.PrefixRegex([]string{"textmagic"}) + `\b([0-9A-Za-z]{30})\b`)
	userPat = regexp.MustCompile(detectors.PrefixRegex([]string{"textmagic"}) + `\b([0-9A-Za-z]{1,25})\b`)
)

// Keywords are used for efficiently pre-filtering chunks.
// Use identifiers in the secret preferably, or the provider name.
func (s Scanner) Keywords() []string {
	return []string{"textmagic"}
}

// FromData will find and optionally verify Textmagic secrets in a given set of bytes.
func (s Scanner) FromData(ctx context.Context, verify bool, data []byte) (results []detectors.Result, err error) {
	dataStr := string(data)

	matches := keyPat.FindAllStringSubmatch(dataStr, -1)
	userMatches := userPat.FindAllStringSubmatch(dataStr, -1)

	for _, match := range matches {
		if len(match) != 2 {
			continue
		}
		resMatch := strings.TrimSpace(match[1])

		for _, userMatch := range userMatches {
			if len(match) != 2 {
				continue
			}
			resUser := strings.TrimSpace(userMatch[1])

			s1 := detectors.Result{
				DetectorType: detectorspb.DetectorType_Textmagic,
				Raw:          []byte(resMatch),
			}

			if verify {
				data := fmt.Sprintf("%s:%s", resUser, resMatch)
				sEnc := b64.StdEncoding.EncodeToString([]byte(data))
				req, _ := http.NewRequestWithContext(ctx, "GET", "https://rest.textmagic.com/api/v2/user", nil)
				req.Header.Add("Content-Type", "application/json")
				req.Header.Add("Authorization", fmt.Sprintf("Basic %s", sEnc))
				res, err := client.Do(req)
				if err == nil {
					defer res.Body.Close()
					if res.StatusCode >= 200 && res.StatusCode < 300 {
						s1.Verified = true
					} else {
						//This function will check false positives for common test words, but also it will make sure the key appears 'random' enough to be a real key
						if detectors.IsKnownFalsePositive(resMatch, detectors.DefaultFalsePositives, true) {
							continue
						}
					}
				}
			}

			results = append(results, s1)
		}
	}

	return detectors.CleanResults(results), nil
}

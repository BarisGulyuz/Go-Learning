package main

import (
	"fmt"
	"strings"
)

// ...
var specialRouteConfig map[string]string = map[string]string{
	"{any}":     ".*",
	"{numeric}": `\d+`,
	"{string}":  `.*`,
}

func matchRoute(requestPath, routePath, targetUrl string) (bool, string) {

	requestPathParams := clearEmptyStrings(strings.Split(requestPath, "/"))
	routePathParams := clearEmptyStrings(strings.Split(routePath, "/"))

	lenRequestParam := len(requestPathParams)
	lenRouteParam := len(routePathParams)

	fmt.Println(lenRequestParam, lenRouteParam)

	isMatch := false
	var targetURL string = targetUrl

	for i, param := range routePathParams {

		if param == "{any}" {
			if lenRequestParam == lenRouteParam {

				targetURL = strings.Replace(targetURL, "/{any}", ("/" + requestPathParams[i]), -1)

			} else if i+1 == lenRouteParam && lenRequestParam > lenRouteParam {

				targetURL = strings.Replace(targetURL, "/{any}", "", -1)
				j := lenRequestParam - lenRouteParam
				for i := j; i < lenRequestParam; i++ {
					targetURL += ("/" + requestPathParams[i])
				}
				isMatch = true
				break

			} else {
				targetURL = strings.Replace(targetURL, "/{any}", "", -1)
			}
		}

		if lenRequestParam > i+1 {
			requestParam := requestPathParams[i]

			isMatch = requestParam == param
			if !isMatch {
				break
			}
		}

	}

	return isMatch, targetURL
}

func clearEmptyStrings(input []string) []string {
	var result []string

	for _, str := range input {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}

// re := regexp.MustCompile(value)
// isMatch = re.MatchString(requestPathParams[i])

func main() {

	requestPath := "/categories"
	routePath := "/categories/{any}"

	targetURL := "https://localhost:44359/api/categories/{any}"

	fmt.Println(matchRoute(requestPath, routePath, targetURL))
}

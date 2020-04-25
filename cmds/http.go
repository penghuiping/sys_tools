package cmds

import (
	"errors"
	"os"
	"sort"
	"strings"
	"sys_tools/utils"
)

func httpCmd() {
	url, err := getHTTPUrl()
	if err != nil {
		utils.Println(err)
		return
	}

	method, err := getHTTPMethod()
	if err != nil {
		utils.Println(err)
		return
	}

	headers, err := getHTTPHeaders()
	if err != nil {
		utils.Println(err)
		return
	}

	switch method {
	case "get":
		{
			res, headers, err3 := new(utils.HTTP).Get(url, headers)
			if err3 != nil {
				utils.Println(err3)
				return
			}

			utils.Println()
			//打印http headers信息
			keys := make([]string, 0)
			for k := range headers {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for i := range keys {
				utils.Println(keys[i], headers[keys[i]])
			}
			utils.Println()
			//打印http body信息
			utils.Println(string(res))
			break
		}
	case "post":
		{
			body, err3 := getHTTPBody()
			if err3 != nil {
				utils.Println(err3)
				return
			}

			res, headers, err4 := new(utils.HTTP).Post(url, headers, body)
			if err4 != nil {
				utils.Println(err4)
				return
			}

			utils.Println()
			//打印http headers信息
			keys := make([]string, 0)
			for k := range headers {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for i := range keys {
				utils.Println(keys[i], headers[keys[i]])
			}
			utils.Println()

			//打印http body信息
			utils.Println(string(res))
			break
		}
	default:
		break
	}
}

func getHTTPMethod() (string, error) {
	for i, arg := range os.Args {
		if arg == "-X" && i+1 < len(os.Args) {
			if os.Args[i+1] == "post" || os.Args[i+1] == "get" {
				return os.Args[i+1], nil
			}
		}
	}
	return "", errors.New("没有提供正确http方法:get/post")
}

func getHTTPHeaders() (map[string]string, error) {
	headers := make(map[string]string)
	for i, arg := range os.Args {
		if arg == "-H" && i+1 < len(os.Args) {
			if !strings.Contains(os.Args[i+1], ":") {
				return nil, errors.New("http header格式不正确，正确的为: -H 'Content-Type: application/json;charset=UTF-8'")
			}
			header := strings.Split(os.Args[i+1], ":")
			headers[header[0]] = header[1]
		}
	}
	return headers, nil
}

func getHTTPBody() (string, error) {
	for i, arg := range os.Args {
		if arg == "-D" && i+1 < len(os.Args) {
			return os.Args[i+1], nil
		}
	}
	return "", nil
}

func getHTTPUrl() (string, error) {
	url := utils.GetSecondCmdLineArgs()
	if utils.IsBlankStr(url) {
		return "", errors.New("没有提供url")
	} else if strings.Contains(url, "http") || strings.Contains(url, "https") {

		return url, nil
	} else {
		return "", errors.New("请提供正确的url,必须以http或者https开头")
	}

}

package requests

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"runtime"
	"time"
)

func Get(url string)(*http.Response, error){
	UserAgent := GetRandomUa()
	var client *http.Client

	//----Modificando resolver y dialer en caso que sea android
	if(runtime.GOOS == "android"){
		dialer := LoadResolverAndDialer()
		client = &http.Client{
			Timeout: 50 * time.Second,
			Transport: &http.Transport{
				DialContext: dialer.DialContext,
			},
		}
	}else{
		client = &http.Client{
			Timeout: 50 * time.Second,
		}
	}



	req, err := http.NewRequest("GET", url, nil)
	if(err != nil){
		//fmt.Println("error Get:", err.Error())
		return nil, fmt.Errorf("Error http NewRquests: %s", err.Error())
	}
	
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Accept", "applicattion/json")
	

	resp, err := client.Do(req)
	if(err != nil){

		return nil, fmt.Errorf("Error get: %s", err.Error())
		
	}
	if(resp.StatusCode != 200){
		return nil , fmt.Errorf("Error status code: %d", resp.StatusCode)
	}


	
	return resp, nil


}





var UserAgents [2]string = [2]string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36",

}



func GetRandomUa()string{
	ua := rand.Intn(len(UserAgents)-1)
	return UserAgents[ua]
}


func LoadResolverAndDialer()*net.Dialer{
	resolver := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}

	dialer := net.Dialer{
		Resolver: &resolver,
	}

	return &dialer

}


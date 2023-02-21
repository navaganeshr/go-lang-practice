package main

import "encoding/json"

func main() {
	//parse the json file
	//input json file
	jsonData := []byte(`{
		"users": [
			{
				"username": "user1",
				"password": "pass1"
			}
		]
	}`)
	//parse the jsonData
	err := json.Unmarshal(jsonData, &users[username])
}


docker run --rm -it --net=host \
    --cap-add=net_admin --cap-add=net_raw --cap-add=sys_nice \
    jasonish/suricata:latest -i eth0


docker run --rm -it --net=host \
--cap-add=net_admin --cap-add=net_raw --cap-add=sys_nice \
-v /var/log/suricata/:/var/log/suricata \
jasonish/suricata:latest -i eth0
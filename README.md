# simple-go-redis
This is a simple key-value web based api which is inspired by [redis](https://redis.io/). It is open source. But be warned that it is so primitive. And it was actually created for [YemekSepeti](https://en.wikipedia.org/wiki/Yemeksepeti)/[DeliveryHero](https://en.wikipedia.org/wiki/Delivery_Hero) Case Study. You may find the details of this case study [here](/docs/requirements.pdf).

## Running locally
To run this application on your local machine, follow the instructions below;
1. Install necessary libraries; [Go](https://golang.org/doc/install) and [git](https://www.linode.com/docs/guides/how-to-install-git-on-linux-mac-and-windows/) to your local machine.
2. Download this repository by running this command; `git clone https://github.com/ivedi/simple-go-redis.git`
3. Run `cd simple-go-redis` to enter the downloaded folder.
4. After that, run `go run cmd/main/main.go`
5. That's all!

To set a key-value data, you must send a post request to *http://localhost:10000/set/{key}* address with a proper payload. The payload will be used as value of the key. And the value type is always string.   
To get a value of a key, you must send a get request to *http://localhost:10000/get/{key}* address. If there is no data it will return an empty string.
If you want to delete all stored data, you must send a delete request to *http://localhost:10000/flush*.
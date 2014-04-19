cryptogo
=======

 One wallet to rule them all
 Cryptogo go in an interface that let you inspect all you cryptocurrency wallets in one place.
 You don't need to open all of your wallet and look where are your coins. Just fire up cryptogo and have everything under control

##Installation
	go get github.com/Zaibon/cryptogo

##Configuration
Complete the config.json file in the conf directory
```
	{
		//web interface
		"host": "localhost",
		"port" : 8080
	}
```
Create some configuration file for your wallets in conf/walltes
```
{
	"enable":true,
	"name" : "bitcoin",
    "symbol":"BTC",

    "host": "localhost",
    "port": 19332,

    "user": "bitcoinrpc",
    "password": "YOUPASSWORD"
}
```
These configurations should reflect the **bitcoin.conf** file that is in you datadir directory of your wallet. If there is no bitcoin.conf file, create it

##Usage
Just run cryptogo and you're go to go
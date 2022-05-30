# tls-forwarder

A simple proxy/forwarder to add TLS to programs using plain TCP.

> Usage of tls-forwarder:  
> &ensp;&ensp;-c string  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Connect to the specific address  
> &ensp;&ensp;-cert string  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Specify the path to the certificate file  
> &ensp;&ensp;-d  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Run as daemon  
> &ensp;&ensp;-insecuretls  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Allow insecure TLS  
> &ensp;&ensp;-key string  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Specify the path to the private key file  
> &ensp;&ensp;-l string  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Listen at the specific address  
> &ensp;&ensp;-log string  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Specify the path to the log file  
> &ensp;&ensp;-noctls  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Do not connect TLS  
> &ensp;&ensp;-noltls  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Do not listen TLS  
> &ensp;&ensp;-pid string  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Specify the path to the pid file  
> &ensp;&ensp;-servername string  
> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;Specify the server name in the certificate presented by the server  

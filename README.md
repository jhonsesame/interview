1.Test calculation module:  
  cd myapp  
  go test test/module_test.go  

2.Test http server with a test client:  
  a.start server:  
    cd myapp  
    go run server.go  
  b.start client:  
    cd myapp/client  
    go run client.go -sk <private key with hex form and '0x' prefix> -addr <wallet address with hex form and '0x' prefix>  
    (eg: go run client.go -sk 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19 -addr 0x96216849c49358B10257cb55b28eA603c874b05E)  
      
3.Test result:  

    server print info with right pk, sk pair:  
    Server provide random_message: sAyjpKCuTlQPNzU3513f6OYqtkHbO6MC  
    Provided Wallet Address:  0x96216849c49358B10257cb55b28eA603c874b05E  
    Signature Derived Address:  0x96216849c49358B10257cb55b28eA603c874b05E  
    matches:  true  
    client print info with right pk, sk pair:  
    GET /get_message: {"message":"sAyjpKCuTlQPNzU3513f6OYqtkHbO6MC"}  
    POST /verify: {"verified":true}  

    server print info with false pk, sk pair:  
    Server provide random_message: VgDxM5uu9B5qV6HM7mew9eSk4U71i2fE  
    Provided Wallet Address:  0x96216849c49358B10257cb55b28eA603c874b05e  
    Signature Derived Address:  0x96216849c49358B10257cb55b28eA603c874b05E  
    matches:  false  
    client print info with false pk, sk pair:  
    GET /get_message: {"message":"VgDxM5uu9B5qV6HM7mew9eSk4U71i2fE"}  
    POST /verify: {"verified":false}

    

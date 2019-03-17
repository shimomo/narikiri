# Narikiri
Ngrok is recommended for testing webhooks in stripe. However, ngrok's free plan will issue a random URL every time. Therefore it is not suitable for continuous webhook test. Narikiri is a command to synchronize the random URL issued by ngrok and all webhook endpoint URLs of stripe.

## Installation
Via homebrew.
```
$ brew tap shimomo/narikiri
$ brew install narikiri
```

## Usage
Launch ngrok on any port.
```
$ ngrok http 8080
```

Launch narikiri by passing the secret key for testing. All webhook endpoint URLs will be replaced. Please be careful too.
```
$ narikiri --key sk_test_************************
```

## License
The narikiri is open source software licensed under the [MIT license](LICENSE).

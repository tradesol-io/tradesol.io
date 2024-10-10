[size=24][b]tradesol.io[/b][/size]
[i]Trade SOL to Any Token with a Single API Request[/i]

[hr]

[b]No API keys. No logs. No limits.[/b]

[hr]

[size=18][b]Features[/b][/size]

[list]
[*][b]Simple Integration[/b]: No need for additional libraries. Start trading immediately with a single API call.
[*][b]Privacy Focused[/b]: We keep no logs of your requests. Your transactions remain private and secure.
[*][b]Unlimited Access[/b]: No usage limits or restrictions. Trade as much as you need.
[*][b]Lightweight[/b]: Minimal overhead and super fast. Enhance your projects without the bloat.
[/list]

[size=18][b]Get Started[/b][/size]

Make a single POST request to [code]https://api.tradesol.io/[/code] with the following JSON payload:

[code=javascript]
{
  "private_key": "your_private_key",
  "token_mint": "token_address",
  "gas_fee": 0.000001,
  "amount_sol": 0.1
}
[/code]

Replace the parameters with your own values:

- [b]private_key[/b]: Your Solana wallet private key in base58 format.
- [b]token_mint[/b]: The mint address of the token you want to receive.
- [b]gas_fee[/b]: Optional; defaults to 0.000001 if not provided.
- [b]amount_sol[/b]: The amount of SOL you want to swap.

[size=18][b]Code Examples[/b][/size]

Switch between languages to see how to integrate:

[spoiler=JavaScript][code=javascript]
const fetch = require('node-fetch');

const payload = {
  private_key: 'your_private_key',
  token_mint: 'token_address',
  gas_fee: 0.000001,
  amount_sol: 0.1,
};

fetch('https://api.tradesol.io/', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify(payload),
})
  .then((res) => res.json())
  .then((data) => console.log(data))
  .catch((err) => console.error(err));
[/code][/spoiler]

[spoiler=Python][code=python]
import requests

payload = {
    'private_key': 'your_private_key',
    'token_mint': 'token_address',
    'gas_fee': 0.000001,
    'amount_sol': 0.1,
}

response = requests.post('https://api.tradesol.io/', json=payload)
print(response.json())
[/code][/spoiler]

[spoiler=Go][code=go]
package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
)

func main() {
  payload := map[string]interface{}{
    "private_key": "your_private_key",
    "token_mint":  "token_address",
    "gas_fee":     0.000001,
    "amount_sol":  0.1,
  }

  jsonData, _ := json.Marshal(payload)
  resp, err := http.Post("https://api.tradesol.io/", "application/json", bytes.NewBuffer(jsonData))
  if err != nil {
    fmt.Println(err)
    return
  }
  defer resp.Body.Close()

  var result map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&result)
  fmt.Println(result)
}
[/code][/spoiler]

[spoiler=PHP][code=php]
<?php
$payload = [
    'private_key' => 'your_private_key',
    'token_mint' => 'token_address',
    'gas_fee' => 0.000001,
    'amount_sol' => 0.1,
];

$ch = curl_init('https://api.tradesol.io/');
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, ['Content-Type: application/json']);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));

$response = curl_exec($ch);
curl_close($ch);

echo $response;
?>
[/code][/spoiler]

[spoiler=Shell][code]
curl -X POST https://api.tradesol.io/ \
-H 'Content-Type: application/json' \
-d '{
  "private_key": "your_private_key",
  "token_mint": "token_address",
  "gas_fee": 0.000001,
  "amount_sol": 0.1
}'
[/code][/spoiler]

[spoiler=Ruby][code=ruby]
require 'net/http'
require 'json'

uri = URI('https://api.tradesol.io/')
payload = {
  private_key: 'your_private_key',
  token_mint: 'token_address',
  gas_fee: 0.000001,
  amount_sol: 0.1,
}

http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = true
request = Net::HTTP::Post.new(uri.path, {'Content-Type' => 'application/json'})
request.body = payload.to_json

response = http.request(request)
puts response.body
[/code][/spoiler]

[size=18][b]Frequently Asked Questions[/b][/size]

[b]Do I need an API key?[/b]
No, you can start using the API immediately without any API keys.

[b]Are there any usage limits?[/b]
There are no usage limits. You can make as many requests as you need.

[b]Do you keep logs of my requests?[/b]
We respect your privacy and do not keep any logs.

[b]How do I get started?[/b]
Simply make a POST request to our endpoint with the required parameters.

[b]Are there any fees?[/b]
Yes, there are fees, but they are less than one percent per transaction.

[b]Which DEXs is it compatible with?[/b]
It is compatible with Raydium and Pumpfun.

[b]Is there any support available?[/b]
Yes, you can contact us at [email]contact@tradesol.io[/email] for any inquiries.

[hr]

[size=18][b]Contact Us[/b][/size]

For support or inquiries, please reach out to us:

- [email]contact@tradesol.io[/email]
- [url=https://github.com/tradesol-io/tradesol.io]GitHub Repository[/url]
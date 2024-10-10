# tradesol.io
*Trade SOL to Any Token with a Single API Request*

---

**No API keys. No logs. No limits.**

---

## Features

- **Simple Integration**: No need for additional libraries. Start trading immediately with a single API call.
- **Privacy Focused**: We keep no logs of your requests. Your transactions remain private and secure.
- **Unlimited Access**: No usage limits or restrictions. Trade as much as you need.
- **Lightweight**: Minimal overhead and super fast. Enhance your projects without the bloat.

## Get Started

### You can host your own server with source code or use https://api.tradesol.io now!

Make a single POST request to `https://api.tradesol.io/` with the following JSON payload:

```json
{
  "private_key": "your_private_key",
  "token_mint": "token_address",
  "gas_fee": 0.000001,
  "amount_sol": 0.1
}
```

Replace the parameters with your own values:

- **private_key**: Your Solana wallet private key in base58 format.
- **token_mint**: The mint address of the token you want to receive.
- **gas_fee**: Optional; defaults to 0.000001 if not provided.
- **amount_sol**: The amount of SOL you want to swap.

## Code Examples

Switch between languages to see how to integrate:

<details>
<summary>JavaScript</summary>

```javascript
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
```

</details>

<details>
<summary>Python</summary>

```python
import requests

payload = {
    'private_key': 'your_private_key',
    'token_mint': 'token_address',
    'gas_fee': 0.000001,
    'amount_sol': 0.1,
}

response = requests.post('https://api.tradesol.io/', json=payload)
print(response.json())
```

</details>

<details>
<summary>Go</summary>

```go
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
```

</details>

<details>
<summary>PHP</summary>

```php
&lt;?php
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
```

</details>

<details>
<summary>Shell</summary>

```shell
curl -X POST https://api.tradesol.io/ \
-H 'Content-Type: application/json' \
-d '{
  "private_key": "your_private_key",
  "token_mint": "token_address",
  "gas_fee": 0.000001,
  "amount_sol": 0.1
}'
```

</details>

<details>
<summary>Ruby</summary>

```ruby
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
```

</details>

## Frequently Asked Questions

**Do I need an API key?**

No, you can start using the API immediately without any API keys.

**Are there any usage limits?**

There are no usage limits. You can make as many requests as you need.

**Do you keep logs of my requests?**

We respect your privacy and do not keep any logs.

**How do I get started?**

Simply make a POST request to our endpoint with the required parameters.

**Are there any fees?**

Yes, there are fees, but they are less than one percent per transaction.

**Which DEXs is it compatible with?**

It is compatible with Raydium and Pumpfun.

**Is there any support available?**

Yes, you can contact us at [contact@tradesol.io](mailto:contact@tradesol.io) for any inquiries.

---

## Contact Us

For support or inquiries, please reach out to us:

- [contact@tradesol.io](mailto:contact@tradesol.io)
- [GitHub Repository](https://github.com/tradesol-io/tradesol.io)
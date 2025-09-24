package pc

import (
	"context"
	"easy-pay-go/payment/alipay"
	"easy-pay-go/pkg/xhttp"
	"testing"
)

var (
	ctx = context.Background()
)

func TestTradePagePay(t *testing.T) {
	config, err := alipay.DefaultConfig("9021000154656122", "MIIEpAIBAAKCAQEAsiIhN7QHi5dT7tAlNs+EjP5lGKCckXstpjxa1ZERRuEEuhKxGSPIgjSsNflKmAiPmt1zroROKMOmWQJUDEXEFE0AlIsknWdIP/BPnXHmsn4ZTwJxoCLyTzLlkLdnV52oY7MAAiinc7fgr/p3BpDVtGdyRUnJ3oVQet07w3wJtiWnRG9MOibXjfu/K1YpnYx4jGWnojswFV62tc1yJPrijNaOf9in5Z/TKYR4TbDNQJYaJJsg+xG4nVe0DyC9Di7JIrA6oylV+R/yoJY6avip834MH0oIdvduoRhdzilrOU6/vltc4sKieZGKiO0rlMBzPiMxNoy85LSgV+c26klogQIDAQABAoIBAECMZYUN034dw+hRvGp0IBAHTxLxgEqnXA9U54VkH+KTP2c9xrbZZJrqA007nHAjCgaQII8omjnjS7ANS+92iGrizEcHdKkk9+VTa0YeM/6R+xHKTOo+P5e7Vbxu+R+inZYjEum1WZhxFRYvqFnUudu5cIrzBheAL32WFUXs+Ifr2hxrovAp9YgKlqHBWhp5ViNCaaDcP8YiggB/RyLJk8LN+rwmKMiMCqkLG8STR/t1PeVyTBQndGuj5pLLnUS7gLcMm7bwK1b8+Yv9GQjbCL0KHK1iUpD0E3nAul/3/uyS6yRUMmyQY1b3m8FU9twaEjzexAOLko0R7dqhTOOOEMECgYEA3vRbTGqHvHoliqS3Rf7O/V6hszMlBroUStrHQtLBhrqA+oDeIPhF6XY9Ef3FJm3Oag89WUrNYSKu5VlAQ0tWT0EuUhUAlyzTagHhEqT5pHGMJkFKDdZ23ZM9TubXRrhPeIyKE5/2NoDNe9M8vuIB/NPaHy+/9FApnDcdJ3GPVI0CgYEAzIkbNQXFE2FzTbQytDvUiyBkepYW40MIcnar5ZCRKlwJ+SGYCEGZAVFr/9IbuUrEp6VtHTYHR0anZ14E3iOu+sfLRJO5Cpxu2di3xk1u/Ed8xF5rmwn8qOEjf1n+M3XPM/ixAM/l6AiECfTZ8lBV32wg9bcepQr8pbvQDykYuMUCgYBJJwXZsc5tgepVbiwQFEXYDjeh83L/nQhRcy0T+NC0ovb0ulmnma5epoPtJkWMkVS7qdpoNMnkSBv4dmtGaS6dfHZ1ShzSfUu0qX8uveCHGxZv6tYgajDADLJ98/HIa2rv07TuMXr4cRR6v1lcDA245c3Yk7M4tyCf99lh+rmbdQKBgQCttqC55r2bmT+htjL+KnZ49zj/eGvV9r+835ddMOxpyqPZyLNkB7qYiPut2VgchikBrZk/nvyNh690NkDyHeJmMbws8T3OilAHi22B6ZAiieCosvy3P8Mr95L7fYZNVHh3zncelnPQHzNtdVyJ1bspN5/CsRVA3VT4ucuapK/PdQKBgQCP/xEg/4cg6JTuZp9duQIYfgm4N2jmtUhYyCAT9YLHUGQyL8m+th9aiNQBcEbkp1vfcz9HuE3edyAsvOs4kgQSWmOqXB+GZ4mk/w1URXhUPXZFa6VZ41hz4hdiAEdnHaVuK1PTvkZnWzOYBcRoJBNYVekk71LxIKcFHCEZ83brqw==", true)
	if err != nil {
		t.Errorf("init config fail %v", err)
	}
	client := alipay.NewClient[*TradePagePayResponse](*config)
	model := TradePagePayModel{
		OutTradeNo:  "GZ201909081743431443",
		ProductCode: "FAST_INSTANT_TRADE_PAY",
		Subject:     "网站支付测试",
		TotalAmount: "88.88",
	}
	request := TradePagePayRequest[TradePagePayResponse]{
		BizModel:   &model,
		Encrypt:    false,
		NotifyUrl:  "https://www.fmm.ink",
		ReturnUrl:  "https://www.fmm.ink",
		HttpMethod: xhttp.HttpGet,
	}
	resp, exeErr := client.PageExecute(ctx, &request)
	if exeErr != nil {
		t.Errorf("execute fail %v", exeErr)
	}
	t.Logf("execute success, resp %s", resp.GetBody())
}

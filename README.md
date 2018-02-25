# human-readable-json
Util for making JSON human-readable

Go from:
```
$ curl -s https://api.robinhood.com
{"mfa":"https:\/\/api.robinhood.com\/mfa\/","margin_interest_charges":"https:\/\/api.robinhood.com\/cash_journal\/margin_interest_charges\/","margin_upgrades":"https:\/\/api.robinhood.com\/margin\/upgrades\/","instruments":"https:\/\/api.robinhood.com\/instruments\/","quotes":"https:\/\/api.robinhood.com\/quotes\/","accounts":"https:\/\/api.robinhood.com\/accounts\/","orders":"https:\/\/api.robinhood.com\/orders\/","subscription_fees":"https:\/\/api.robinhood.com\/subscription\/subscription_fees\/","id_documents":"https:\/\/api.robinhood.com\/upload\/photo_ids\/","portfolios":"https:\/\/api.robinhood.com\/portfolios\/","markets":"https:\/\/api.robinhood.com\/markets\/","wire_relationships":"https:\/\/api.robinhood.com\/wire\/relationships\/","ach_queued_deposit":"https:\/\/api.robinhood.com\/ach\/queued_deposit\/","subscriptions":"https:\/\/api.robinhood.com\/subscription\/subscriptions\/","wire_transfers":"https:\/\/api.robinhood.com\/wire\/transfers\/","dividends":"https:\/\/api.robinhood.com\/dividends\/","notification_settings":"https:\/\/api.robinhood.com\/settings\/notifications\/","applications":"https:\/\/api.robinhood.com\/applications\/","user":"https:\/\/api.robinhood.com\/user\/","ach_relationships":"https:\/\/api.robinhood.com\/ach\/relationships\/","ach_deposit_schedules":"https:\/\/api.robinhood.com\/ach\/deposit_schedules\/","ach_iav_auth":"https:\/\/api.robinhood.com\/ach\/iav\/auth\/","notifications":"https:\/\/api.robinhood.com\/notifications\/","ach_transfers":"https:\/\/api.robinhood.com\/ach\/transfers\/","positions":"https:\/\/api.robinhood.com\/positions\/","watchlists":"https:\/\/api.robinhood.com\/watchlists\/","document_requests":"https:\/\/api.robinhood.com\/upload\/document_requests\/","edocuments":"https:\/\/api.robinhood.com\/documents\/","password_reset":"https:\/\/api.robinhood.com\/password_reset\/request\/","password_change":"https:\/\/api.robinhood.com\/password_change\/"}
```

To:
```
$ curl -s https://api.robinhood.com | json
{
	accounts:                https://api.robinhood.com/accounts/
	ach_deposit_schedules:   https://api.robinhood.com/ach/deposit_schedules/
	ach_iav_auth:            https://api.robinhood.com/ach/iav/auth/
	ach_queued_deposit:      https://api.robinhood.com/ach/queued_deposit/
	ach_relationships:       https://api.robinhood.com/ach/relationships/
	ach_transfers:           https://api.robinhood.com/ach/transfers/
	applications:            https://api.robinhood.com/applications/
	dividends:               https://api.robinhood.com/dividends/
	document_requests:       https://api.robinhood.com/upload/document_requests/
	edocuments:              https://api.robinhood.com/documents/
	id_documents:            https://api.robinhood.com/upload/photo_ids/
	instruments:             https://api.robinhood.com/instruments/
	margin_interest_charges: https://api.robinhood.com/cash_journal/margin_interest_charges/
	margin_upgrades:         https://api.robinhood.com/margin/upgrades/
	markets:                 https://api.robinhood.com/markets/
	mfa:                     https://api.robinhood.com/mfa/
	notification_settings:   https://api.robinhood.com/settings/notifications/
	notifications:           https://api.robinhood.com/notifications/
	orders:                  https://api.robinhood.com/orders/
	password_change:         https://api.robinhood.com/password_change/
	password_reset:          https://api.robinhood.com/password_reset/request/
	portfolios:              https://api.robinhood.com/portfolios/
	positions:               https://api.robinhood.com/positions/
	quotes:                  https://api.robinhood.com/quotes/
	subscription_fees:       https://api.robinhood.com/subscription/subscription_fees/
	subscriptions:           https://api.robinhood.com/subscription/subscriptions/
	user:                    https://api.robinhood.com/user/
	watchlists:              https://api.robinhood.com/watchlists/
	wire_relationships:      https://api.robinhood.com/wire/relationships/
	wire_transfers:          https://api.robinhood.com/wire/transfers/
}
```

😀

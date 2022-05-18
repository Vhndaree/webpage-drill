package api

import "fmt"

func StartWeb() {
	fmt.Println(ReadWebPage("https://identity.gebuhrenfrei.com/auth/realms/advanzia-deu/protocol/openid-connect/auth?client_id=advanzia-web-client&redirect_uri=https%3A%2F%2Fmein.gebuhrenfrei.com%2Fretail-app-de%2Fredirect&state=6da6b2ed-082a-4294-8179-f9f3603b3835&response_mode=fragment&response_type=code%20id_token%20token&scope=openid&nonce=3e3ea2fd-ed28-414e-9ba6-feb3e0d0f13a&ui_locales=de&code_challenge=3S-SQeibDKuei3JaV8bQo3FbwP4yzHhp0nyLTCo3_NA&code_challenge_method=S256"))
}

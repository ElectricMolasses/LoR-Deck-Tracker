package main

import "fmt"

type fieldCard struct {
}

var playerDeck map[string]card

// Store the discard as an array to track order.
// The same will go for the field, but they'll contain
// maps that contain cards, to track board position, etc.
var playerDiscard []card
var playerField []int

var opponentField []int

func getField() {
	data, err := queryClient("positional-rectangles")
	fmt.Println(data, err)
}

/*
LOG ONE(EMPTY FIELD)::
{"PlayerName":"Red Onion King","OpponentName":"game_npe_tutorials_Tutorial9a_EnemyName","GameState":"InProgress","Screen":{"ScreenWidth":1348,"ScreenHeight":816},"Rectangles":[{"CardID":543469548,"CardCode":"face","TopLeftX":84,"TopLeftY":363,"Width":88,"Height":88,"LocalPlayer":true},{"CardID":1272090411,"CardCode":"face","TopLeftX":84,"TopLeftY":541,"Width":88,"Height":88,"LocalPlayer":false},{"CardID":545541096,"CardCode":"01FR008","TopLeftX":416,"TopLeftY":63,"Width":133,"Height":185,"LocalPlayer":true},{"CardID":895179405,"CardCode":"01FR027","TopLeftX":543,"TopLeftY":65,"Width":133,"Height":186,"LocalPlayer":true},{"CardID":736157868,"CardCode":"01FR025","TopLeftX":670,"TopLeftY":63,"Width":134,"Height":187,"LocalPlayer":true},{"CardID":1026058934,"CardCode":"01FR027","TopLeftX":799,"TopLeftY":59,"Width":134,"Height":188,"LocalPlayer":true}]}

LOG TWO(ONE PORO PLAYED)::
{"PlayerName":"Red Onion King","OpponentName":"game_npe_tutorials_Tutorial9a_EnemyName","GameState":"InProgress","Screen":{"ScreenWidth":1348,"ScreenHeight":816},"Rectangles":[{"CardID":543469548,"CardCode":"face","TopLeftX":84,"TopLeftY":363,"Width":88,"Height":88,"LocalPlayer":true},{"CardID":1272090411,"CardCode":"face","TopLeftX":84,"TopLeftY":541,"Width":88,"Height":88,"LocalPlayer":false},{"CardID":545541096,"CardCode":"01FR008","TopLeftX":625,"TopLeftY":196,"Width":96,"Height":121,"LocalPlayer":true},{"CardID":895179405,"CardCode":"01FR027","TopLeftX":416,"TopLeftY":63,"Width":133,"Height":185,"LocalPlayer":true},{"CardID":736157868,"CardCode":"01FR025","TopLeftX":543,"TopLeftY":65,"Width":133,"Height":186,"LocalPlayer":true},{"CardID":1026058934,"CardCode":"01FR027","TopLeftX":670,"TopLeftY":63,"Width":134,"Height":187,"LocalPlayer":true},{"CardID":263940831,"CardCode":"01DE049","TopLeftX":799,"TopLeftY":59,"Width":134,"Height":188,"LocalPlayer":true}]}

LOG THREE(Initial Draw/Mulligan Screen)::
{"PlayerName":"Red Onion King","OpponentName":"game_npe_tutorials_Tutorial9a_EnemyName","GameState":"InProgress","Screen":{"ScreenWidth":1348,"ScreenHeight":816},"Rectangles":[{"CardID":543469548,"CardCode":"face","TopLeftX":84,"TopLeftY":363,"Width":88,"Height":88,"LocalPlayer":true},{"CardID":1272090411,"CardCode":"face","TopLeftX":84,"TopLeftY":541,"Width":88,"Height":88,"LocalPlayer":false},{"CardID":545541096,"CardCode":"01FR008","TopLeftX":625,"TopLeftY":196,"Width":96,"Height":121,"LocalPlayer":true},{"CardID":895179405,"CardCode":"01FR027","TopLeftX":416,"TopLeftY":63,"Width":133,"Height":185,"LocalPlayer":true},{"CardID":736157868,"CardCode":"01FR025","TopLeftX":543,"TopLeftY":65,"Width":133,"Height":186,"LocalPlayer":true},{"CardID":1026058934,"CardCode":"01FR027","TopLeftX":670,"TopLeftY":63,"Width":134,"Height":187,"LocalPlayer":true},{"CardID":263940831,"CardCode":"01DE049","TopLeftX":799,"TopLeftY":59,"Width":134,"Height":188,"LocalPlayer":true}]}

LOG FOUR(Initial Draw/Mulligan with all cards set to replace)::
{"PlayerName":"Red Onion King","OpponentName":"game_npe_tutorials_Tutorial7_EnemyName","GameState":"InProgress","Screen":{"ScreenWidth":1348,"ScreenHeight":816},"Rectangles":[{"CardID":1845158949,"CardCode":"face","TopLeftX":84,"TopLeftY":363,"Width":88,"Height":88,"LocalPlayer":true},{"CardID":898120731,"CardCode":"face","TopLeftX":84,"TopLeftY":541,"Width":88,"Height":88,"LocalPlayer":false},{"CardID":1783516474,"CardCode":"00IO004","TopLeftX":256,"TopLeftY":549,"Width":191,"Height":282,"LocalPlayer":true},{"CardID":1466806486,"CardCode":"00IO015","TopLeftX":471,"TopLeftY":549,"Width":191,"Height":282,"LocalPlayer":true},{"CardID":1786059646,"CardCode":"00IO008","TopLeftX":685,"TopLeftY":549,"Width":191,"Height":282,"LocalPlayer":true},{"CardID":1056413465,"CardCode":"00IO008","TopLeftX":900,"TopLeftY":549,"Width":191,"Height":282,"LocalPlayer":true},{"CardID":664906281,"CardCode":"00IO005","TopLeftX":625,"TopLeftY":196,"Width":96,"Height":121,"LocalPlayer":true},{"CardID":1820382672,"CardCode":"00FR007","TopLeftX":507,"TopLeftY":740,"Width":96,"Height":121,"LocalPlayer":false},{"CardID":426651211,"CardCode":"00FR007","TopLeftX":625,"TopLeftY":740,"Width":96,"Height":121,"LocalPlayer":false},{"CardID":200743855,"CardCode":"00FR007","TopLeftX":744,"TopLeftY":740,"Width":96,"Height":121,"LocalPlayer":false}]}

FOLLOWING TWO::
Opponent has 3 sets of mana stones in play, and a wolf attacking.
Player has scout card defending.
Player has pending Steel Tempest spell.

LOG FIVE(Pending spell with no target)::
{"PlayerName":"Red Onion King","OpponentName":"game_npe_tutorials_Tutorial7_EnemyName","GameState":"InProgress","Screen":{"ScreenWidth":1348,"ScreenHeight":816},"Rectangles":[{"CardID":1845158949,"CardCode":"face","TopLeftX":84,"TopLeftY":363,"Width":88,"Height":88,"LocalPlayer":true},{"CardID":898120731,"CardCode":"face","TopLeftX":84,"TopLeftY":541,"Width":88,"Height":88,"LocalPlayer":false},{"CardID":664906281,"CardCode":"00IO005","TopLeftX":570,"TopLeftY":340,"Width":207,"Height":119,"LocalPlayer":true},{"CardID":1820382672,"CardCode":"00FR007","TopLeftX":744,"TopLeftY":740,"Width":96,"Height":121,"LocalPlayer":false},{"CardID":426651211,"CardCode":"00FR007","TopLeftX":507,"TopLeftY":740,"Width":96,"Height":121,"LocalPlayer":false},{"CardID":200743855,"CardCode":"00FR007","TopLeftX":625,"TopLeftY":740,"Width":96,"Height":121,"LocalPlayer":false},{"CardID":1720598905,"CardCode":"00IO006","TopLeftX":310,"TopLeftY":58,"Width":132,"Height":185,"LocalPlayer":true},{"CardID":1528662892,"CardCode":"00IO010","TopLeftX":428,"TopLeftY":63,"Width":133,"Height":185,"LocalPlayer":true},{"CardID":384260689,"CardCode":"00IO010","TopLeftX":630,"TopLeftY":452,"Width":87,"Height":87,"LocalPlayer":true},{"CardID":1656150882,"CardCode":"00IO014","TopLeftX":667,"TopLeftY":63,"Width":134,"Height":187,"LocalPlayer":true},{"CardID":643928399,"CardCode":"00IO012T2","TopLeftX":787,"TopLeftY":59,"Width":134,"Height":188,"LocalPlayer":true},{"CardID":1661864982,"CardCode":"00FR020","TopLeftX":570,"TopLeftY":596,"Width":207,"Height":119,"LocalPlayer":false},{"CardID":86250897,"CardCode":"00IO015","TopLeftX":909,"TopLeftY":51,"Width":135,"Height":188,"LocalPlayer":true},{"CardID":1086841230,"CardCode":"00FR020T1","TopLeftX":625,"TopLeftY":196,"Width":96,"Height":121,"LocalPlayer":true}]}

LOG SIX(PENDING SPELL TARGETING WOLF)::
{"PlayerName":"Red Onion King","OpponentName":"game_npe_tutorials_Tutorial7_EnemyName","GameState":"InProgress","Screen":{"ScreenWidth":1348,"ScreenHeight":816},"Rectangles":
[
	// Presumably the two face cards are the two nexi/health pools.
	{"CardID":1845158949,"CardCode":"face","TopLeftX":84,"TopLeftY":363,"Width":88,"Height":88,"LocalPlayer":true},
	{"CardID":898120731,"CardCode":"face","TopLeftX":84,"TopLeftY":541,"Width":88,"Height":88,"LocalPlayer":false},

	// Nimble Poro
	{"CardID":664906281,"CardCode":"00IO005","TopLeftX":570,"TopLeftY":340,"Width":207,"Height":119,"LocalPlayer":true},
	// Babbling Bjerg, 90% it was a card in hand.
	{"CardID":1820382672,"CardCode":"00FR007","TopLeftX":744,"TopLeftY":740,"Width":96,"Height":121,"LocalPlayer":false},
	// Babbling Bjerg.  Alright.
	{"CardID":426651211,"CardCode":"00FR007","TopLeftX":507,"TopLeftY":740,"Width":96,"Height":121,"LocalPlayer":false},
	// Babbling Bjerg...
	{"CardID":200743855,"CardCode":"00FR007","TopLeftX":625,"TopLeftY":740,"Width":96,"Height":121,"LocalPlayer":false},
	// Greenglade Duo
	{"CardID":1720598905,"CardCode":"00IO006","TopLeftX":310,"TopLeftY":58,"Width":132,"Height":185,"LocalPlayer":true},
	// Stand United.
	{"CardID":1528662892,"CardCode":"00IO010","TopLeftX":428,"TopLeftY":63,"Width":133,"Height":185,"LocalPlayer":true},
	// Stand United again.
	{"CardID":384260689,"CardCode":"00IO010","TopLeftX":630,"TopLeftY":452,"Width":87,"Height":87,"LocalPlayer":true},
	// Greenglade Elder
	{"CardID":1656150882,"CardCode":"00IO014","TopLeftX":667,"TopLeftY":63,"Width":134,"Height":187,"LocalPlayer":true},

	// Steel Tempest, what I'm looking for, but looks to be hand card.
	{"CardID":643928399,"CardCode":"00IO012T2","TopLeftX":787,"TopLeftY":59,"Width":134,"Height":188,"LocalPlayer":true},
	// Avalanche
	{"CardID":1661864982,"CardCode":"00FR020","TopLeftX":570,"TopLeftY":596,"Width":207,"Height":119,"LocalPlayer":false},
	// Yasuo
	{"CardID":86250897,"CardCode":"00IO015","TopLeftX":909,"TopLeftY":51,"Width":135,"Height":188,"LocalPlayer":true},
	// Avalanche with T1, this cardCode is not referenced in set1.
	{"CardID":1086841230,"CardCode":"00FR020T1","TopLeftX":625,"TopLeftY":196,"Width":96,"Height":121,"LocalPlayer":true}
	]
}
*/

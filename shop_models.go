package fnapiio

import (
	"encoding/json"
	"time"
)

type Shop struct {
	Result            bool            `json:"result"`
	FullShop          bool            `json:"fullShop"`
	LastUpdate        ShopLastUpdate  `json:"lastUpdate"`
	CurrentRotation   json.RawMessage `json:"currentRotation"`
	NextRotation      json.RawMessage `json:"nextRotation"`
	Carousel          *string         `json:"carousel,omitempty"`
	SpecialOfferVideo *string         `json:"specialOfferVideo,omitempty"`
	CustomBackground  *string         `json:"customBackground,omitempty"`
	Shop              []ShopOffer     `json:"shop"`
}

type ShopLastUpdate struct {
	UID  string    `json:"uid"`
	Date time.Time `json:"date"`
}

type ShopOffer struct {
	MainID              string             `json:"mainId"`
	DisplayName         string             `json:"displayName"`
	DisplayDescription  string             `json:"displayDescription"`
	DisplayType         string             `json:"displayType"`
	MainType            string             `json:"mainType"`
	OfferID             string             `json:"offerId"`
	DisplayAssets       []ShopDisplayAsset `json:"displayAssets"`
	FirstReleaseDate    string             `json:"firstReleaseDate"`
	PreviousReleaseDate string             `json:"previousReleaseDate"`
	GiftAllowed         bool               `json:"giftAllowed"`
	BuyAllowed          bool               `json:"buyAllowed"`
	Price               ShopOfferPrice     `json:"price"`
	Rarity              SimpleRarity       `json:"rarity"`
	Series              *ShopOfferSeries   `json:"series,omitempty"`
	Banner              ShopOfferBanner    `json:"banner"`
	OfferTag            *ShopOfferTag      `json:"offerTag,omitempty"`
	Granted             []FullCosmetic     `json:"granted"`
}

type ShopDisplayAsset struct {
	DisplayAsset      string `json:"displayAsset"`
	MaterialInstance  string `json:"materialInstance"`
	URL               string `json:"url"`
	Flipbook          string `json:"flipbook"`
	BackgroundTexture string `json:"background_texture"`
	Background        string `json:"background"`
	FullBackground    string `json:"full_background"`
}

type ShopOfferPrice struct {
	Floor   float64 `json:"floorPrice"`
	Final   float64 `json:"finalPrice"`
	Regular float64 `json:"regularPrice"`
}

type SimpleRarity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ShopOfferSeries struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ShopOfferBanner struct {
	Intensity *string `json:"intensity,omitempty"`
	Name      string  `json:"name"`
	ID        string  `json:"id"`
}

type ShopOfferTag struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type FullCosmetic struct {
	ID               string                `json:"id"`
	Type             CosmeticType          `json:"type"`
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	Rarity           SimpleRarity          `json:"rarity"`
	Series           *CosmeticSeries       `json:"series,omitempty"`
	Price            float64               `json:"price"`
	Added            CosmeticAdded         `json:"added"`
	BuiltInEmote     *CosmeticBuiltInEmote `json:"builtInEmote,omitempty"`
	CopyrightedAudio bool                  `json:"copyrightedAudio"`
	Upcoming         bool                  `json:"upcoming"`
	Reactive         bool                  `json:"reactive"`
	ReleaseDate      *string               `json:"releaseDate,omitempty"`
	LastAppearance   *string               `json:"lastAppearance,omitempty"`
	Interest         float64               `json:"interest"`
	Images           CosmeticImages        `json:"images"`
	Video            *string               `json:"video,omitempty"`
	Audio            *string               `json:"audio,omitempty"`
	GameplayTags     []string              `json:"gameplayTags"`
	ApiTags          []string              `json:"apiTags"`
	SearchTags       []string              `json:"searchTags"`
	BattlePass       *CosmeticBattlePass   `json:"battlepass,omitempty"`
	Set              *CosmeticSet          `json:"set,omitempty"`
	Introduction     *CosmeticIntroduction `json:"introduction,omitempty"`
	DisplayAssets    []ShopDisplayAsset    `json:"displayAssets"`
	ShopHistory      []time.Time           `json:"shopHistory"`
	Styles           []CosmeticStyle       `json:"styles"`
	Grants           []json.RawMessage     `json:"grants"`
	GrantedBy        []json.RawMessage     `json:"grantedBy"`
}

type CosmeticType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CosmeticSeries struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CosmeticAdded struct {
	Version string    `json:"version"`
	Date    time.Time `json:"date"`
}

type CosmeticBuiltInEmote struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CosmeticImages struct {
	FullBackground string `json:"full_background"`
	IconBackground string `json:"icon_background"`
	Background     string `json:"background"`
	Featured       string `json:"featured"`
	Icon           string `json:"icon"`
}

type CosmeticBattlePass struct {
	Season         float64               `json:"season"`
	Tier           float64               `json:"tier"`
	Type           string                `json:"type"`
	DisplayText    BattlePassDisplayText `json:"displayText"`
	BattlePassName string                `json:"battlePassName"`
}

type BattlePassDisplayText struct {
	ChapterSeason string `json:"chapterSeason"`
	Season        string `json:"season"`
	Chapter       string `json:"chapter"`
}

type CosmeticSet struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	PartOf string `json:"partOf"`
}

type CosmeticIntroduction struct {
	Chapter string `json:"chapter"`
	Season  string `json:"season"`
	Text    string `json:"text"`
}

type CosmeticStyle struct {
	Name           string  `json:"name"`
	Channel        string  `json:"channel"`
	ChannelName    string  `json:"channelName"`
	Tag            string  `json:"tag"`
	IsDefault      bool    `json:"isDefault"`
	StartUnlocked  bool    `json:"startUnlocked"`
	HideIfNotOwned bool    `json:"hideIfNotOwned"`
	Image          *string `json:"image"`
}

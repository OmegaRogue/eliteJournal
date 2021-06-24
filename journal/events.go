/*
 * Copyright (c) 2021 OmegaRogue.
 *  This file is part of eliteJournal.
 *
 *     eliteJournal is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or any later version.
 *
 *     eliteJournal is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with eliteJournal.  If not, see <https://www.gnu.org/licenses/>.
 */

package journal

type ApproachBody struct {
	Body          string
	BodyID        float64
	StarSystem    string
	SystemAddress float64
	Event
}
type ApproachSettlement struct {
	Latitude  float64
	Longitude float64
	MarketID  float64
	Name      string
	Event
}
type AsteroidCracked struct {
	Body string
	Event
}
type BackPack struct {
	Event
}
type Backpack struct {
	Components  []interface{}
	Consumables []interface{}
	Data        []interface{}
	Items       []interface{}
	Event
}
type BackpackChange struct {
	Removed []struct {
		LocalizedValue
		OwnerID float64
		Type    string
	}
	Event
}
type BookTaxi struct {
	Cost                float64
	DestinationLocation string
	DestinationSystem   string
	Event
}
type Bounty struct {
	Rewards []struct {
		Faction string
		Reward  float64
	}
	Target        string
	TotalReward   float64
	VictimFaction string
	Event
}
type BuyAmmo struct {
	Cost float64
	Event
}
type BuyDrones struct {
	BuyPrice  float64
	Count     float64
	TotalCost float64
	Type      string
	Event
}
type BuyMicroResources struct {
	Category string

	MarketID float64
	LocalizedValue
	Price float64
	Event
}
type BuySuit struct {
	LocalizedValue
	Price  float64
	SuitID float64
	Event
}
type BuyWeapon struct {
	LocalizedValue
	Price        float64
	SuitModuleID float64
	Event
}

type CargoDepot struct {
	CargoType           string
	CargoTypeLocalised  string `json:"CargoType_Localised,omitempty"`
	Count               float64
	EndMarketID         float64
	ItemsCollected      float64
	ItemsDelivered      float64
	MissionID           float64
	Progress            float64
	StartMarketID       float64
	TotalItemsToDeliver float64
	UpdateType          string
	Event
}
type CarrierDepositFuel struct {
	Amount    float64
	CarrierID float64
	Total     float64
	Event
}
type CarrierJump struct {
	Body     string
	BodyID   float64
	BodyType string
	Docked   bool
	Factions []struct {
		Allegiance         string
		FactionState       string
		Government         string
		Happiness          string
		HappinessLocalised string `json:"Happiness_Localised,omitempty"`
		Influence          float64
		MyReputation       float64
		Name               string
	}
	MarketID         float64
	Multicrew        bool
	Population       float64
	PowerplayState   string
	Powers           []string
	StarPos          []float64
	StarSystem       string
	StationEconomies []struct {
		LocalizedValue
		Proportion float64
	}
	StationEconomy          string
	StationEconomyLocalised string `json:"StationEconomy_Localised,omitempty"`
	StationFaction          struct {
		Name string
	}
	StationGovernment          string
	StationGovernmentLocalised string `json:"StationGovernment_Localised,omitempty"`
	StationName                string
	StationServices            []string
	StationType                string
	SystemAddress              float64
	SystemAllegiance           string
	SystemEconomy              string
	SystemEconomyLocalised     string `json:"SystemEconomy_Localised,omitempty"`
	SystemFaction              struct {
		Name string
	}
	SystemGovernment             string
	SystemGovernmentLocalised    string `json:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string
	SystemSecondEconomyLocalised string `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string
	SystemSecurityLocalised      string `json:"SystemSecurity_Localised,omitempty"`
	Taxi                         bool
	Event
}
type ChangeCrewRole struct {
	Role string
	Event
}
type CodexEntry struct {
	Category          string
	CategoryLocalised string `json:"Category_Localised,omitempty"`
	EntryID           float64
	IsNewEntry        bool
	LocalizedValue
	Region               string
	RegionLocalised      string `json:"Region_Localised,omitempty"`
	SubCategory          string
	SubCategoryLocalised string `json:"SubCategory_Localised,omitempty"`
	System               string
	SystemAddress        float64
	Event
}
type CollectCargo struct {
	Stolen        bool
	Type          string
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Event
}
type CollectItems struct {
	LocalizedValue
	OwnerID float64
	Stolen  bool
	Type    string
	Event
}

type CommitCrime struct {
	CrimeType string
	Faction   string
	Fine      float64
	Event
}
type CreateSuitLoadout struct {
	LoadoutID   float64
	LoadoutName string
	Modules     []struct {
		ModuleName          string
		ModuleNameLocalised string `json:"ModuleName_Localised,omitempty"`
		SlotName            string
		SuitModuleID        float64
	}
	SuitID            float64
	SuitName          string
	SuitNameLocalised string `json:"SuitName_Localised,omitempty"`
	Event
}
type CrewAssign struct {
	CrewID float64
	Name   string
	Role   string
	Event
}
type CrewFire struct {
	CrewID float64
	Name   string
	Event
}
type CrewHire struct {
	CombatRank float64
	Cost       float64
	CrewID     float64
	Faction    string
	Name       string
	Event
}
type CrewMemberJoins struct {
	Crew string
	Event
}
type CrewMemberQuits struct {
	Crew string
	Event
}
type CrewMemberRoleChange struct {
	Crew string
	Role string
	Event
}
type CrimeVictim struct {
	CrimeType string
	Fine      float64
	Offender  string
	Event
}
type DataScanned struct {
	Type          string
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Event
}
type DeleteSuitLoadout struct {
	LoadoutID         float64
	LoadoutName       string
	SuitID            float64
	SuitName          string
	SuitNameLocalised string `json:"SuitName_Localised,omitempty"`
	Event
}
type Died struct {
	Event
}
type Disembark struct {
	Body          string
	BodyID        float64
	ID            float64
	Multicrew     bool
	OnPlanet      bool
	OnStation     bool
	SRV           bool
	StarSystem    string
	SystemAddress float64
	Taxi          bool
	Event
}
type DockFighter struct {
	ID float64
	Event
}
type DockSRV struct {
	ID float64
	Event
}
type Docked struct {
	DistFromStarLS    float64
	MarketID          float64
	StarSystem        string
	StationAllegiance string
	StationEconomies  []struct {
		LocalizedValue
		Proportion float64
	}
	StationEconomy             string
	StationEconomyLocalised    string `json:"StationEconomy_Localised,omitempty"`
	StationFaction             string
	StationGovernment          string
	StationGovernmentLocalised string `json:"StationGovernment_Localised,omitempty"`
	StationName                string
	StationServices            []string
	StationType                string
	SystemAddress              float64
	Event
}
type DockingCancelled struct {
	MarketID    float64
	StationName string
	StationType string
	Event
}
type DockingDenied struct {
	MarketID    float64
	Reason      string
	StationName string
	StationType string
	Event
}
type DockingGranted struct {
	LandingPad  float64
	MarketID    float64
	StationName string
	StationType string
	Event
}
type DockingRequested struct {
	MarketID    float64
	StationName string
	StationType string
	Event
}
type EjectCargo struct {
	Abandoned     bool
	Count         float64
	Type          string
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Event
}
type Embark struct {
	Body          string
	BodyID        float64
	ID            float64
	Multicrew     bool
	OnPlanet      bool
	OnStation     bool
	SRV           bool
	StarSystem    string
	SystemAddress float64
	Taxi          bool
	Event
}
type EndCrewSession struct {
	OnCrime bool
	Event
}
type EngineerContribution struct {
	Engineer      string
	EngineerID    float64
	Quantity      float64
	TotalQuantity float64
	Type          string
	Event
}
type EngineerCraft struct {
	BlueprintID   float64
	BlueprintName string
	Engineer      string
	EngineerID    float64
	Ingredients   []LocalizedValue
	Level         float64
	Modifiers     []struct {
		Label         string
		LessIsGood    float64
		OriginalValue float64
		Value         float64
	}
	Module  string
	Quality float64
	Slot    string
	Event
}
type EngineerProgress struct {
	Engineers []interface{}
	Event
}
type EscapeInterdiction struct {
	Interdictor string
	IsPlayer    bool
	Event
}
type FSDJump struct {
	FactionState string
	Factions     []struct {
		ActiveStates []struct {
			State string
		}
		Allegiance         string
		FactionState       string
		Government         string
		Happiness          string
		HappinessLocalised string `json:"Happiness_Localised,omitempty"`
		Influence          float64
		MyReputation       float64
		Name               string
	}
	FuelLevel                    float64
	FuelUsed                     float64
	JumpDist                     float64
	Population                   float64
	StarPos                      []float64
	StarSystem                   string
	SystemAddress                float64
	SystemAllegiance             string
	SystemEconomy                string
	SystemEconomyLocalised       string `json:"SystemEconomy_Localised,omitempty"`
	SystemFaction                string
	SystemGovernment             string
	SystemGovernmentLocalised    string `json:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string
	SystemSecondEconomyLocalised string `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string
	SystemSecurityLocalised      string `json:"SystemSecurity_Localised,omitempty"`
	Event
}
type FSDTarget struct {
	Name          string
	SystemAddress float64
	Event
}
type FSSAllBodiesFound struct {
	Count         float64
	SystemAddress float64
	SystemName    string
	Event
}
type FSSDiscoveryScan struct {
	BodyCount    float64
	NonBodyCount float64
	Progress     float64
	Event
}
type FSSSignalDiscovered struct {
	SignalName    string
	SystemAddress float64
	Event
}
type FetchRemoteModule struct {
	ServerId            float64
	Ship                string
	ShipID              float64
	StorageSlot         float64
	StoredItem          string
	StoredItemLocalised string `json:"StoredItem_Localised,omitempty"`
	TransferCost        float64
	TransferTime        float64
	Event
}
type FighterDestroyed struct {
	ID float64
	Event
}
type FighterRebuilt struct {
	ID      float64
	Loadout string
	Event
}
type Fileheader struct {
	Build       string  `json:"build,omitempty"`
	Event       string  `json:"event,omitempty"`
	GameVersion string  `json:"gameversion,omitempty"`
	Language    string  `json:"language,omitempty"`
	Part        float64 `json:"part,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
}
type Friends struct {
	Name   string
	Status string
	Event
}
type FuelScoop struct {
	Scooped float64
	Total   float64
	Event
}
type HeatDamage struct {
	Event
}
type HeatWarning struct {
	Event
}
type HullDamage struct {
	Fighter     bool
	Health      float64
	PlayerPilot bool
	Event
}
type Interdicted struct {
	Faction     string
	Interdictor string
	IsPlayer    bool
	Submitted   bool
	Event
}
type Interdiction struct {
	Faction  string
	IsPlayer bool
	Success  bool
	Event
}
type JoinACrew struct {
	Captain string
	Event
}
type JoinedSquadron struct {
	SquadronName string
	Event
}
type KickCrewMember struct {
	Crew    string
	OnCrime bool
	Event
}
type LaunchDrone struct {
	Type string
	Event
}
type LaunchFighter struct {
	ID               float64
	Loadout          string
	PlayerControlled bool
	Event
}
type LaunchSRV struct {
	ID               float64
	Loadout          string
	PlayerControlled bool
	Event
}
type LeaveBody struct {
	Body          string
	BodyID        float64
	StarSystem    string
	SystemAddress float64
	Event
}
type Liftoff struct {
	Latitude         float64
	Longitude        float64
	PlayerControlled bool
	Event
}

type LoadoutEquipModule struct {
	LoadoutID           float64
	LoadoutName         string
	ModuleName          string
	ModuleNameLocalised string `json:"ModuleName_Localised,omitempty"`
	SlotName            string
	SuitID              float64
	SuitModuleID        float64
	SuitName            string
	SuitNameLocalised   string `json:"SuitName_Localised,omitempty"`
	Event
}
type Location struct {
	Body     string
	BodyID   float64
	BodyType string
	Docked   bool
	Factions []struct {
		Allegiance         string
		FactionState       string
		Government         string
		Happiness          string
		HappinessLocalised string `json:"Happiness_Localised,omitempty"`
		Influence          float64
		MyReputation       float64
		Name               string
	}
	MarketID                     float64
	Population                   float64
	StarPos                      []float64
	StarSystem                   string
	StationName                  string
	StationType                  string
	SystemAddress                float64
	SystemAllegiance             string
	SystemEconomy                string
	SystemEconomyLocalised       string `json:"SystemEconomy_Localised,omitempty"`
	SystemFaction                string
	SystemGovernment             string
	SystemGovernmentLocalised    string `json:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string
	SystemSecondEconomyLocalised string `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string
	SystemSecurityLocalised      string `json:"SystemSecurity_Localised,omitempty"`
	Event
}
type Market struct {
	MarketID    float64
	StarSystem  string
	StationName string
	Event
}
type MarketBuy struct {
	BuyPrice      float64
	Count         float64
	MarketID      float64
	TotalCost     float64
	Type          string
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Event
}
type MarketSell struct {
	AvgPricePaid float64
	Count        float64
	MarketID     float64
	SellPrice    float64
	TotalSale    float64
	Type         string
	Event
}
type MassModuleStore struct {
	Items []struct {
		Hot bool
		LocalizedValue
		Slot string
	}
	MarketID float64
	Ship     string
	ShipID   float64
	Event
}
type MaterialCollected struct {
	Category string
	LocalizedValue
	Event
}
type MaterialDiscovered struct {
	Category        string
	DiscoveryNumber float64
	LocalizedValue
	Event
}

type MiningRefined struct {
	Type          string
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Event
}
type MissionAbandoned struct {
	MissionID float64
	Name      string
	Event
}
type MissionAccepted struct {
	DestinationStation  string
	DestinationSystem   string
	Expiry              string
	Faction             string
	Influence           string
	LocalisedName       string
	MissionID           float64
	Name                string
	Reputation          string
	Reward              float64
	Target              string
	TargetFaction       string
	TargetType          string
	TargetTypeLocalised string `json:"TargetType_Localised,omitempty"`
	Wing                bool
	Event
}
type MissionCompleted struct {
	DestinationStation string
	DestinationSystem  string
	Faction            string
	FactionEffects     []struct {
		Effects []struct {
			Effect          string
			EffectLocalised string `json:"Effect_Localised,omitempty"`
			Trend           string
		}
		Faction   string
		Influence []struct {
			Influence     string
			SystemAddress float64
			Trend         string
		}
		Reputation      string
		ReputationTrend string
	}
	MissionID             float64
	Name                  string
	NewDestinationStation string
	NewDestinationSystem  string
	Reward                float64
	Target                string
	TargetFaction         string
	TargetType            string
	TargetTypeLocalised   string `json:"TargetType_Localised,omitempty"`
	Event
}
type MissionFailed struct {
	MissionID float64
	Name      string
	Event
}
type MissionRedirected struct {
	MissionID             float64
	Name                  string
	NewDestinationStation string
	NewDestinationSystem  string
	OldDestinationStation string
	OldDestinationSystem  string
	Event
}
type ModuleBuy struct {
	BuyItem          string
	BuyItemLocalised string `json:"BuyItem_Localised,omitempty"`
	BuyPrice         float64
	MarketID         float64
	Ship             string
	ShipID           float64
	Slot             string
	Event
}
type ModuleInfo struct {
	Event
}
type ModuleRetrieve struct {
	Hot                    bool
	MarketID               float64
	RetrievedItem          string
	RetrievedItemLocalised string `json:"RetrievedItem_Localised,omitempty"`
	Ship                   string
	ShipID                 float64
	Slot                   string
	Event
}
type ModuleSell struct {
	MarketID          float64
	SellItem          string
	SellItemLocalised string `json:"SellItem_Localised,omitempty"`
	SellPrice         float64
	Ship              string
	ShipID            float64
	Slot              string
	Event
}
type ModuleSellRemote struct {
	SellItem          string
	SellItemLocalised string `json:"SellItem_Localised,omitempty"`
	SellPrice         float64
	ServerId          float64
	Ship              string
	ShipID            float64
	StorageSlot       float64
	Event
}
type ModuleStore struct {
	Hot                 bool
	MarketID            float64
	Ship                string
	ShipID              float64
	Slot                string
	StoredItem          string
	StoredItemLocalised string `json:"StoredItem_Localised,omitempty"`
	Event
}
type ModuleSwap struct {
	FromItem          string
	FromItemLocalised string `json:"FromItem_Localised,omitempty"`
	FromSlot          string
	MarketID          float64
	Ship              string
	ShipID            float64
	ToItem            string
	ToSlot            string
	Event
}
type MultiSellExplorationData struct {
	BaseValue  float64
	Bonus      float64
	Discovered []struct {
		NumBodies  float64
		SystemName string
	}
	TotalEarnings float64
	Event
}
type Music struct {
	MusicTrack string
	Event
}
type NavBeaconScan struct {
	NumBodies     float64
	SystemAddress float64
	Event
}
type NavRoute struct {
	Event
}
type NpcCrewPaidWage struct {
	Amount      float64
	NpcCrewId   float64
	NpcCrewName string
	Event
}
type NpcCrewRank struct {
	NpcCrewId   float64
	NpcCrewName string
	RankCombat  float64
	Event
}
type Outfitting struct {
	MarketID    float64
	StarSystem  string
	StationName string
	Event
}
type Passengers struct {
	Manifest []struct {
		Count     float64
		MissionID float64
		Type      string
		VIP       bool
		Wanted    bool
	}
	Event
}
type PayBounties struct {
	Amount           float64
	BrokerPercentage float64
	Faction          string
	FactionLocalised string `json:"Faction_Localised,omitempty"`
	ShipID           float64
	Event
}
type PayFines struct {
	AllFines bool
	Amount   float64
	ShipID   float64
	Event
}
type Powerplay struct {
	Merits      float64
	Power       string
	Rank        float64
	TimePledged float64
	Votes       float64
	Event
}
type PowerplayCollect struct {
	Count         float64
	Power         string
	Type          string
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Event
}
type PowerplayDeliver struct {
	Count         float64
	Power         string
	Type          string
	TypeLocalised string `json:"Type_Localised,omitempty"`
	Event
}
type PowerplayJoin struct {
	Power string
	Event
}
type PowerplaySalary struct {
	Amount float64
	Power  string
	Event
}
type Progress struct {
	CQC        float64
	Combat     float64
	Empire     float64
	Explore    float64
	Federation float64
	Trade      float64
	Event
}
type Promotion struct {
	Trade float64
	Event
}
type ProspectedAsteroid struct {
	Content          string
	ContentLocalised string `json:"Content_Localised,omitempty"`
	Materials        []struct {
		LocalizedValue
		Proportion float64
	}
	Remaining float64
	Event
}
type QuitACrew struct {
	Captain string
	Event
}
type Rank struct {
	CQC        float64
	Combat     float64
	Empire     float64
	Explore    float64
	Federation float64
	Trade      float64
	Event
}
type ReceiveText struct {
	Channel          string
	From             string
	Message          string
	MessageLocalised string `json:"Message_Localised,omitempty"`
	Event
}
type RedeemVoucher struct {
	Amount   float64
	Factions []struct {
		Amount  float64
		Faction string
	}
	Type string
	Event
}
type RefuelAll struct {
	Amount float64
	Cost   float64
	Event
}
type RenameSuitLoadout struct {
	LoadoutID         float64
	LoadoutName       string
	SuitID            float64
	SuitName          string
	SuitNameLocalised string `json:"SuitName_Localised,omitempty"`
	Event
}
type Repair struct {
	Cost float64
	Item string
	Event
}
type RepairAll struct {
	Cost float64
	Event
}
type RepairDrone struct {
	CockpitRepaired float64
	HullRepaired    float64
	Event
}
type Reputation struct {
	Event
}
type ReservoirReplenished struct {
	FuelMain      float64
	FuelReservoir float64
	Event
}
type RestockVehicle struct {
	Cost    float64
	Count   float64
	Loadout string
	Type    string
	Event
}
type Resurrect struct {
	Bankrupt bool
	Cost     float64
	Option   string
	Event
}
type SAAScanComplete struct {
	BodyID           float64
	BodyName         string
	EfficiencyTarget float64
	ProbesUsed       float64
	Event
}
type SAASignalsFound struct {
	BodyID   float64
	BodyName string
	Signals  []struct {
		Count float64
		Type  string
	}
	SystemAddress float64
	Event
}
type Scan struct {
	BodyID                float64
	BodyName              string
	DistanceFromArrivalLS float64
	Parents               []struct {
		Ring float64
	}
	ScanType string
	Event
}
type Scanned struct {
	ScanType string
	Event
}
type SelfDestruct struct {
	Event
}
type SellDrones struct {
	Count     float64
	SellPrice float64
	TotalSale float64
	Type      string
	Event
}
type SellExplorationData struct {
	BaseValue     float64
	Bonus         float64
	Discovered    []interface{}
	Systems       []string
	TotalEarnings float64
	Event
}
type SellSuit struct {
	LocalizedValue
	Price    float64
	SuitID   float64
	SuitMods []interface{}
	Event
}
type SellWeapon struct {
	Class float64
	LocalizedValue
	Price        float64
	SuitModuleID float64
	WeaponMods   []interface{}
	Event
}
type SendText struct {
	Message string
	To      string
	Event
}
type SetUserShipName struct {
	Ship         string
	ShipID       float64
	UserShipId   string
	UserShipName string
	Event
}
type SharedBookmarkToSquadron struct {
	SquadronName string
	Event
}
type ShieldState struct {
	ShieldsUp bool
	Event
}
type OwnedItem struct {
	LocalizedValue
	OwnerID int
}
type ShipLocker struct {
	Components  []OwnedItem
	Consumables []OwnedItem
	Data        []interface{}
	Items       []OwnedItem
	Event
}
type ShipLockerMaterials struct {
	Components  []interface{}
	Consumables []interface{}
	Data        []interface{}
	Items       []interface{}
	Event
}
type ShipTargeted struct {
	Bounty             float64
	Faction            string
	HullHealth         float64
	LegalStatus        string
	PilotName          string
	PilotNameLocalised string `json:"PilotName_Localised,omitempty"`
	PilotRank          string
	ScanStage          float64
	ShieldHealth       float64
	Ship               string
	ShipLocalised      string `json:"Ship_Localised,omitempty"`
	TargetLocked       bool
	Event
}
type Shipyard struct {
	MarketID    float64
	StarSystem  string
	StationName string
	Event
}
type ShipyardBuy struct {
	MarketID     float64
	ShipPrice    float64
	ShipType     string
	StoreOldShip string
	StoreShipID  float64
	Event
}
type ShipyardNew struct {
	NewShipID float64
	ShipType  string
	Event
}
type ShipyardSell struct {
	MarketID     float64
	SellShipID   float64
	ShipMarketID float64
	ShipPrice    float64
	ShipType     string
	System       string
	Event
}
type ShipyardSwap struct {
	MarketID     float64
	ShipID       float64
	ShipType     string
	StoreOldShip string
	StoreShipID  float64
	Event
}
type ShipyardTransfer struct {
	Distance      float64
	MarketID      float64
	ShipID        float64
	ShipMarketID  float64
	ShipType      string
	System        string
	TransferPrice float64
	TransferTime  float64
	Event
}
type Shutdown struct {
	Event
}
type SquadronCreated struct {
	SquadronName string
	Event
}
type SquadronStartup struct {
	CurrentRank  float64
	SquadronName string
	Event
}
type StartJump struct {
	JumpType      string
	StarClass     string
	StarSystem    string
	SystemAddress float64
	Event
}
type StoredModules struct {
	Items       []interface{}
	MarketID    float64
	StarSystem  string
	StationName string
	Event
}
type StoredShips struct {
	MarketID    float64
	ShipsHere   []interface{}
	ShipsRemote []interface{}
	StarSystem  string
	StationName string
	Event
}
type SuitLoadout struct {
	LoadoutID   float64
	LoadoutName string
	Modules     []struct {
		ModuleName          string
		ModuleNameLocalised string `json:"ModuleName_Localised,omitempty"`
		SlotName            string
		SuitModuleID        float64
	}
	SuitID            float64
	SuitName          string
	SuitNameLocalised string `json:"SuitName_Localised,omitempty"`
	Event
}
type SupercruiseEntry struct {
	StarSystem    string
	SystemAddress float64
	Event
}
type SupercruiseExit struct {
	Body          string
	BodyID        float64
	BodyType      string
	StarSystem    string
	SystemAddress float64
	Event
}
type SwitchSuitLoadout struct {
	LoadoutID   float64
	LoadoutName string
	Modules     []struct {
		ModuleName          string
		ModuleNameLocalised string `json:"ModuleName_Localised,omitempty"`
		SlotName            string
		SuitModuleID        float64
	}
	SuitID            float64
	SuitName          string
	SuitNameLocalised string `json:"SuitName_Localised,omitempty"`
	Event
}
type Synthesis struct {
	Materials []LocalizedValue
	Name      string
	Event
}
type Touchdown struct {
	Latitude         float64
	Longitude        float64
	PlayerControlled bool
	Event
}
type TransferMicroResources struct {
	Transfers []struct {
		Category string
		LocalizedValue
		Direction string
	}
	Event
}
type USSDrop struct {
	USSThreat        float64
	USSType          string
	USSTypeLocalised string `json:"USSType_Localised,omitempty"`
	Event
}
type UnderAttack struct {
	Target string
	Event
}
type Undocked struct {
	MarketID    float64
	StationName string
	StationType string
	Event
}
type UseConsumable struct {
	LocalizedValue
	Type string
	Event
}
type VehicleSwitch struct {
	To string
	Event
}
type WingAdd struct {
	Name string
	Event
}
type WingInvite struct {
	Name string
	Event
}
type WingJoin struct {
	Others []interface{}
	Event
}
type WingLeave struct {
	Event
}

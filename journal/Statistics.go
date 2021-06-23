package journal

type Statistics struct {
	Event
	BankAccount struct {
		CurrentWealth          Cost `json:"Current_Wealth"`
		SpentOnShips           Cost `json:"Spent_On_Ships"`
		SpentOnOutfitting      Cost `json:"Spent_On_Outfitting"`
		SpentOnRepairs         Cost `json:"Spent_On_Repairs"`
		SpentOnFuel            Cost `json:"Spent_On_Fuel"`
		SpentOnAmmoConsumables Cost `json:"Spent_On_Ammo_Consumables"`
		InsuranceClaims        int  `json:"Insurance_Claims"`
		SpentOnInsurance       Cost `json:"Spent_On_Insurance"`
		OwnedShipCount         int  `json:"Owned_Ship_Count"`
	} `json:"Bank_Account"`
	Combat struct {
		BountiesClaimed      int `json:"Bounties_Claimed"`
		BountyHuntingProfit  int `json:"Bounty_Hunting_Profit"`
		CombatBonds          int `json:"Combat_Bonds"`
		CombatBondProfits    int `json:"Combat_Bond_Profits"`
		Assassinations       int `json:"Assassinations"`
		AssassinationProfits int `json:"Assassination_Profits"`
		HighestSingleReward  int `json:"Highest_Single_Reward"`
		SkimmersKilled       int `json:"Skimmers_Killed"`
	} `json:"Combat"`
	Crime struct {
		Notoriety        int  `json:"Notoriety"`
		Fines            int  `json:"Fines"`
		TotalFines       Cost `json:"Total_Fines"`
		BountiesReceived int  `json:"Bounties_Received"`
		TotalBounties    Cost `json:"Total_Bounties"`
		HighestBounty    Cost `json:"Highest_Bounty"`
	} `json:"Crime"`
	Smuggling struct {
		BlackMarketsTradedWith   int  `json:"Black_Markets_Traded_With"`
		BlackMarketsProfits      Cost `json:"Black_Markets_Profits"`
		ResourcesSmuggled        int  `json:"Resources_Smuggled"`
		AverageProfit            Cost `json:"Average_Profit"`
		HighestSingleTransaction Cost `json:"Highest_Single_Transaction"`
	} `json:"Smuggling"`
	Trading struct {
		MarketsTradedWith        int  `json:"Markets_Traded_With"`
		MarketProfits            Cost `json:"Market_Profits"`
		ResourcesTraded          int  `json:"Resources_Traded"`
		AverageProfit            Cost `json:"Average_Profit"`
		HighestSingleTransaction Cost `json:"Highest_Single_Transaction"`
	} `json:"Trading"`
	Mining struct {
		MiningProfits      Cost `json:"Mining_Profits"`
		QuantityMined      int  `json:"Quantity_Mined"`
		MaterialsCollected int  `json:"Materials_Collected"`
	} `json:"Mining"`
	Exploration struct {
		SystemsVisited            int     `json:"Systems_Visited"`
		ExplorationProfits        Cost    `json:"Exploration_Profits"`
		PlanetsScannedToLevel2    int     `json:"Planets_Scanned_To_Level_2"`
		PlanetsScannedToLevel3    int     `json:"Planets_Scanned_To_Level_3"`
		EfficientScans            int     `json:"Efficient_Scans"`
		HighestPayout             Cost    `json:"Highest_Payout"`
		TotalHyperspaceDistance   int     `json:"Total_Hyperspace_Distance"`
		TotalHyperspaceJumps      int     `json:"Total_Hyperspace_Jumps"`
		GreatestDistanceFromStart float64 `json:"Greatest_Distance_From_Start"`
		TimePlayed                int     `json:"Time_Played"`
	} `json:"Exploration"`
	Passengers struct {
		PassengersMissionsAccepted    int `json:"Passengers_Missions_Accepted"`
		PassengersMissionsDisgruntled int `json:"Passengers_Missions_Disgruntled"`
		PassengersMissionsBulk        int `json:"Passengers_Missions_Bulk"`
		PassengersMissionsVIP         int `json:"Passengers_Missions_VIP"`
		PassengersMissionsDelivered   int `json:"Passengers_Missions_Delivered"`
		PassengersMissionsEjected     int `json:"Passengers_Missions_Ejected"`
	} `json:"Passengers"`
	SearchAndRescue struct {
		SearchRescueTraded int  `json:"SearchRescue_Traded"`
		SearchRescueProfit Cost `json:"SearchRescue_Profit"`
		SearchRescueCount  int  `json:"SearchRescue_Count"`
	} `json:"Search_And_Rescue"`
	Crafting struct {
		CountOfUsedEngineers                      int  `json:"Count_Of_Used_Engineers"`
		RecipesGenerated                          int  `json:"Recipes_Generated"`
		RecipesGeneratedRank1                     int  `json:"Recipes_Generated_Rank_1"`
		RecipesGeneratedRank2                     int  `json:"Recipes_Generated_Rank_2"`
		RecipesGeneratedRank3                     int  `json:"Recipes_Generated_Rank_3"`
		RecipesGeneratedRank4                     int  `json:"Recipes_Generated_Rank_4"`
		RecipesGeneratedRank5                     int  `json:"Recipes_Generated_Rank_5"`
		RecipesApplied                            int  `json:"Recipes_Applied"`
		RecipesAppliedRank1                       int  `json:"Recipes_Applied_Rank_1"`
		RecipesAppliedRank2                       int  `json:"Recipes_Applied_Rank_2"`
		RecipesAppliedRank3                       int  `json:"Recipes_Applied_Rank_3"`
		RecipesAppliedRank4                       int  `json:"Recipes_Applied_Rank_4"`
		RecipesAppliedRank5                       int  `json:"Recipes_Applied_Rank_5"`
		SpentOnCrafting                           Cost `json:"Spent_On_Crafting"`
		RecipesAppliedOnPreviouslyModifiedModules int  `json:"Recipes_Applied_On_Previously_Modified_Modules"`
	} `json:"Crafting"`
	Crew struct {
		NpcCrewTotalWages Cost `json:"NpcCrew_TotalWages"`
		NpcCrewHired      int  `json:"NpcCrew_Hired"`
		NpcCrewFired      int  `json:"NpcCrew_Fired"`
		NpcCrewDied       int  `json:"NpcCrew_Died"`
	} `json:"Crew"`
	Multicrew struct {
		MulticrewTimeTotal        int  `json:"Multicrew_Time_Total"`
		MulticrewGunnerTimeTotal  int  `json:"Multicrew_Gunner_Time_Total"`
		MulticrewFighterTimeTotal int  `json:"Multicrew_Fighter_Time_Total"`
		MulticrewCreditsTotal     Cost `json:"Multicrew_Credits_Total"`
		MulticrewFinesTotal       Cost `json:"Multicrew_Fines_Total"`
	} `json:"Multicrew"`
	MaterialTraderStats struct {
		TradesCompleted int `json:"Trades_Completed"`
		MaterialsTraded int `json:"Materials_Traded"`
	} `json:"Material_Trader_Stats"`
	CQC struct {
		CQCCreditsEarned Cost    `json:"CQC_Credits_Earned"`
		CQCTimePlayed    int     `json:"CQC_Time_Played"`
		CQCKD            float64 `json:"CQC_KD"`
		CQCKills         int     `json:"CQC_Kills"`
		CQCWL            float64 `json:"CQC_WL"`
	} `json:"CQC"`
}

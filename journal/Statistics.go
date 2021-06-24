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

// Statistics Statistics
type Statistics struct {
	Event
	BankAccount struct {
		CurrentWealth          Cost `json:"Current_Wealth,omitempty"`
		SpentOnShips           Cost `json:"Spent_On_Ships,omitempty"`
		SpentOnOutfitting      Cost `json:"Spent_On_Outfitting,omitempty"`
		SpentOnRepairs         Cost `json:"Spent_On_Repairs,omitempty"`
		SpentOnFuel            Cost `json:"Spent_On_Fuel,omitempty"`
		SpentOnAmmoConsumables Cost `json:"Spent_On_Ammo_Consumables,omitempty"`
		InsuranceClaims        int  `json:"Insurance_Claims,omitempty"`
		SpentOnInsurance       Cost `json:"Spent_On_Insurance,omitempty"`
		OwnedShipCount         int  `json:"Owned_Ship_Count,omitempty"`
	} `json:"Bank_Account,omitempty"`
	Combat struct {
		BountiesClaimed      int `json:"Bounties_Claimed,omitempty"`
		BountyHuntingProfit  int `json:"Bounty_Hunting_Profit,omitempty"`
		CombatBonds          int `json:"Combat_Bonds,omitempty"`
		CombatBondProfits    int `json:"Combat_Bond_Profits,omitempty"`
		Assassinations       int `json:"Assassinations,omitempty"`
		AssassinationProfits int `json:"Assassination_Profits,omitempty"`
		HighestSingleReward  int `json:"Highest_Single_Reward,omitempty"`
		SkimmersKilled       int `json:"Skimmers_Killed,omitempty"`
	} `json:"Combat,omitempty"`
	Crime struct {
		Notoriety        int  `json:"Notoriety,omitempty"`
		Fines            int  `json:"Fines,omitempty"`
		TotalFines       Cost `json:"Total_Fines,omitempty"`
		BountiesReceived int  `json:"Bounties_Received,omitempty"`
		TotalBounties    Cost `json:"Total_Bounties,omitempty"`
		HighestBounty    Cost `json:"Highest_Bounty,omitempty"`
	} `json:"Crime,omitempty"`
	Smuggling struct {
		BlackMarketsTradedWith   int  `json:"Black_Markets_Traded_With,omitempty"`
		BlackMarketsProfits      Cost `json:"Black_Markets_Profits,omitempty"`
		ResourcesSmuggled        int  `json:"Resources_Smuggled,omitempty"`
		AverageProfit            Cost `json:"Average_Profit,omitempty"`
		HighestSingleTransaction Cost `json:"Highest_Single_Transaction,omitempty"`
	} `json:"Smuggling,omitempty"`
	Trading struct {
		MarketsTradedWith        int  `json:"Markets_Traded_With,omitempty"`
		MarketProfits            Cost `json:"Market_Profits,omitempty"`
		ResourcesTraded          int  `json:"Resources_Traded,omitempty"`
		AverageProfit            Cost `json:"Average_Profit,omitempty"`
		HighestSingleTransaction Cost `json:"Highest_Single_Transaction,omitempty"`
	} `json:"Trading,omitempty"`
	Mining struct {
		MiningProfits      Cost `json:"Mining_Profits,omitempty"`
		QuantityMined      int  `json:"Quantity_Mined,omitempty"`
		MaterialsCollected int  `json:"Materials_Collected,omitempty"`
	} `json:"Mining,omitempty"`
	Exploration struct {
		SystemsVisited            int     `json:"Systems_Visited,omitempty"`
		ExplorationProfits        Cost    `json:"Exploration_Profits,omitempty"`
		PlanetsScannedToLevel2    int     `json:"Planets_Scanned_To_Level_2,omitempty"`
		PlanetsScannedToLevel3    int     `json:"Planets_Scanned_To_Level_3,omitempty"`
		EfficientScans            int     `json:"Efficient_Scans,omitempty"`
		HighestPayout             Cost    `json:"Highest_Payout,omitempty"`
		TotalHyperspaceDistance   int     `json:"Total_Hyperspace_Distance,omitempty"`
		TotalHyperspaceJumps      int     `json:"Total_Hyperspace_Jumps,omitempty"`
		GreatestDistanceFromStart float64 `json:"Greatest_Distance_From_Start,omitempty"`
		TimePlayed                int     `json:"Time_Played,omitempty"`
	} `json:"Exploration,omitempty"`
	Passengers struct {
		PassengersMissionsAccepted    int `json:"Passengers_Missions_Accepted,omitempty"`
		PassengersMissionsDisgruntled int `json:"Passengers_Missions_Disgruntled,omitempty"`
		PassengersMissionsBulk        int `json:"Passengers_Missions_Bulk,omitempty"`
		PassengersMissionsVIP         int `json:"Passengers_Missions_VIP,omitempty"`
		PassengersMissionsDelivered   int `json:"Passengers_Missions_Delivered,omitempty"`
		PassengersMissionsEjected     int `json:"Passengers_Missions_Ejected,omitempty"`
	} `json:"Passengers,omitempty"`
	SearchAndRescue struct {
		SearchRescueTraded int  `json:"SearchRescue_Traded,omitempty"`
		SearchRescueProfit Cost `json:"SearchRescue_Profit,omitempty"`
		SearchRescueCount  int  `json:"SearchRescue_Count,omitempty"`
	} `json:"Search_And_Rescue,omitempty"`
	Crafting struct {
		CountOfUsedEngineers                      int  `json:"Count_Of_Used_Engineers,omitempty"`
		RecipesGenerated                          int  `json:"Recipes_Generated,omitempty"`
		RecipesGeneratedRank1                     int  `json:"Recipes_Generated_Rank_1,omitempty"`
		RecipesGeneratedRank2                     int  `json:"Recipes_Generated_Rank_2,omitempty"`
		RecipesGeneratedRank3                     int  `json:"Recipes_Generated_Rank_3,omitempty"`
		RecipesGeneratedRank4                     int  `json:"Recipes_Generated_Rank_4,omitempty"`
		RecipesGeneratedRank5                     int  `json:"Recipes_Generated_Rank_5,omitempty"`
		RecipesApplied                            int  `json:"Recipes_Applied,omitempty"`
		RecipesAppliedRank1                       int  `json:"Recipes_Applied_Rank_1,omitempty"`
		RecipesAppliedRank2                       int  `json:"Recipes_Applied_Rank_2,omitempty"`
		RecipesAppliedRank3                       int  `json:"Recipes_Applied_Rank_3,omitempty"`
		RecipesAppliedRank4                       int  `json:"Recipes_Applied_Rank_4,omitempty"`
		RecipesAppliedRank5                       int  `json:"Recipes_Applied_Rank_5,omitempty"`
		SpentOnCrafting                           Cost `json:"Spent_On_Crafting,omitempty"`
		RecipesAppliedOnPreviouslyModifiedModules int  `json:"Recipes_Applied_On_Previously_Modified_Modules,omitempty"`
	} `json:"Crafting,omitempty"`
	Crew struct {
		NpcCrewTotalWages Cost `json:"NpcCrew_TotalWages,omitempty"`
		NpcCrewHired      int  `json:"NpcCrew_Hired,omitempty"`
		NpcCrewFired      int  `json:"NpcCrew_Fired,omitempty"`
		NpcCrewDied       int  `json:"NpcCrew_Died,omitempty"`
	} `json:"Crew,omitempty"`
	Multicrew struct {
		MulticrewTimeTotal        int  `json:"Multicrew_Time_Total,omitempty"`
		MulticrewGunnerTimeTotal  int  `json:"Multicrew_Gunner_Time_Total,omitempty"`
		MulticrewFighterTimeTotal int  `json:"Multicrew_Fighter_Time_Total,omitempty"`
		MulticrewCreditsTotal     Cost `json:"Multicrew_Credits_Total,omitempty"`
		MulticrewFinesTotal       Cost `json:"Multicrew_Fines_Total,omitempty"`
	} `json:"Multicrew,omitempty"`
	MaterialTraderStats struct {
		TradesCompleted int `json:"Trades_Completed,omitempty"`
		MaterialsTraded int `json:"Materials_Traded,omitempty"`
	} `json:"Material_Trader_Stats,omitempty"`
	CQC struct {
		CQCCreditsEarned Cost    `json:"CQC_Credits_Earned,omitempty"`
		CQCTimePlayed    int     `json:"CQC_Time_Played,omitempty"`
		CQCKD            float64 `json:"CQC_KD,omitempty"`
		CQCKills         int     `json:"CQC_Kills,omitempty"`
		CQCWL            float64 `json:"CQC_WL,omitempty"`
	} `json:"CQC,omitempty"`
}

#pragma once

struct round {
	char course[100];
	int slope;
	int rating;
	char* par;
	char* score;
	char* putts;
	float puttsPerHole;
	char roundScoringAvg;
	float par3Avg;
	float par4Avg;
	float par5Avg;
	
	struct round* prev;
	struct round* next;
};

struct round_list {
	char numRounds;
	float currentHandicap;
	char validHandicap;
	float overallPuttsPerHole;
	float scoringAvg;
	float overallPar3_Avg;
	float overallPar4_Avg;
	float overallPar6_Avg;

	struct round* head;
	struct round* tail;
};

/*Loads the main menu and directs user to selected submenu.*/
void mainMenu();

/*Allows the user to add a new round in and removes the oldest round
if there are more than 20*/
void newRound();

/*Simply returns the current handicap*/
void checkHandicap();

/*Returns a table of round statistics with an internal menu to
explore the stats deeper*/
void roundStats();
#pragma once

class round {
	public:
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

		class round* prev;
		class round* next;
};

class roundList {
	public:
		char numRounds;
		float currentHandicap;
		char validHandicap;
		float overallPuttsPerHole;
		float scoringAvg;
		float overallPar3Avg;
		float overallPar4Avg;
		float overallPar5Avg;

		class round* head;
		class round* tail;
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
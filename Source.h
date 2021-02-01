#pragma once

/*
* All classes below will be saved in a CSV (maybe text file) and loaded up as needed.
*/

/* 
* Nodes of the doubly linked list with pointers to integer arrays for the par, score, and putts.
*/
class round {
public:
	string course;
	int slope;
	int rating;
	int* par;
	int* score;
	int* putts;
	float puttsPerHole;
	float roundScoringAvg;
	float par3Avg;
	float par4Avg;
	float par5Avg;
	int day;
	int month;
	int year;

	class round* prev;
	class round* next;
};

/*
* Contains overall stats and pointers to the head and tail of the list for easy access.
*/
class roundList {
public:
	int numRounds;
	float currentHandicap;
	float overallPuttsPerHole;
	float scoringAvg;
	float overallPar3Avg;
	float overallPar4Avg;
	float overallPar5Avg;

	class round* head;
	class round* tail;
};

/*
* Nodes for a stack (in the form of a DLL) of handicaps to track game stats over time.
*/
class statNode {
public:
	char* handicapDate;
	float handicap;
	float puttingAvg;
	float scoringAvg;
	float par3Avg;
	float par4Avg;
	float par5Avg;

	class statNode* prev;
	class statNode* next;
};

/*
* Keeps track of how many handicaps are being tracked to prevent high memory consumption.
* Contains pointers to the head and the tail for easy access to the list.
*/
class statHistory {
public:
	int numHandicaps;

	class statNode* head;
	class statNode* tail;
};

/*
* Loads the main menu and directs user to selected submenu.
*/
void mainMenu();

/*
* Allows the user to add a new round in and removes the oldest round if there are more than 20.
*/
void newRound();

/*
* Simply returns the current handicap.
*/
void checkHandicap();

/* 
* Returns a table of round statistics with an internal menu to
* explore the stats deeper
*/
void roundStats();